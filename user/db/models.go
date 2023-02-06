package db

import (
	"database/sql"
)

type User struct {
	ID        sql.NullInt32
	Username  sql.NullString
	Email     sql.NullString
	Password  sql.NullString
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
}
