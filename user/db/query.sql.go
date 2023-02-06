package db

import (
	"context"
	"database/sql"
)

const create = `
insert into users (username, email, password, created_at, updated_at) values (?, ?, ?, now(), now())
`

type CreateParams struct {
	Username sql.NullString
	Email    sql.NullString
	Password sql.NullString
}

func (q *Queries) Create(ctx context.Context, arg CreateParams) (sql.Result, error) {
	return q.db.ExecContext(
		ctx,
		create,
		arg.Username,
		arg.Email,
		arg.Password,
	)
}

const delete = `
delete from users where id = ?
`

func (q *Queries) Delete(ctx context.Context, id sql.NullInt32) error {
	_, err := q.db.ExecContext(ctx, delete, id)
	return err
}

const get = `
select id, username, email, password, created_at, updated_at from users where id = ?
`

func (q *Queries) Get(ctx context.Context, id sql.NullInt32) (User, error) {
	row := q.db.QueryRowContext(ctx, get, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const list = `
select id, username, email, password from users order by id
`

type ListRow struct {
	ID       sql.NullInt32
	Username sql.NullString
	Email    sql.NullString
	Password sql.NullString
}

func (q *Queries) List(ctx context.Context) ([]ListRow, error) {
	rows, err := q.db.QueryContext(ctx, list)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListRow
	for rows.Next() {
		var i ListRow
		err = rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
		)
		if err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const update = `
update users set username = ?, email = ?, password = ?, updated_at = now() where id = ?
`

type UpdateParams struct {
	ID       sql.NullInt32
	Username sql.NullString
	Email    sql.NullString
	Password sql.NullString
}

func (q *Queries) Update(ctx context.Context, arg UpdateParams) error {
	_, err := q.db.ExecContext(
		ctx,
		update,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.ID,
	)
	return err
}
