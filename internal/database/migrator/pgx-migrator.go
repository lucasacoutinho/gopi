package migrator

import (
	"fmt"
	"strings"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/iancoleman/strcase"
)

const UP = `
-- CREATE TABLE example (
--	id serial PRIMARY KEY,
--	created_at TIMESTAMP NOT NULL,
--	updated_at TIMESTAMP NOT NULL,
-- )
`
const DOWN = `
-- DROP TABLE example
`

type PGX struct {
	conn string
}

func NewPGX(conn string) *PGX {
	return &PGX{conn: conn}
}

func (p *PGX) Make(root, name string) error {
	filename := fmt.Sprintf(
		"%d_%s",
		time.Now().UnixMicro(),
		strcase.ToSnake(strings.TrimSpace(name)),
	)

	err := MakeMigration(root, filename, "up", UP)
	if err != nil {
		return err
	}

	err = MakeMigration(root, filename, "down", DOWN)
	if err != nil {
		return err
	}

	return nil
}

func (p *PGX) Up(root string) error {
	m, err := migrate.New("file://"+root+"/migrations", p.conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		return err
	}

	return nil
}

func (p *PGX) DownAll(root string) error {
	m, err := migrate.New("file://"+root+"/migrations", p.conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Down(); err != nil {
		return err
	}

	return nil
}

func (p *PGX) Steps(root string, steps int) error {
	m, err := migrate.New("file://"+root+"/migrations", p.conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Steps(steps); err != nil {
		return err
	}

	return nil
}

func (p *PGX) Force(root string) error {
	m, err := migrate.New("file://"+root+"/migrations", p.conn)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Force(-1); err != nil {
		return err
	}

	return nil
}
