package migration

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/lucasacoutinho/gopi/internal/config"
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

func Make(c *config.Config, arg string) error {
	filename := fmt.Sprintf("%d_%s", time.Now().UnixMicro(), arg)

	err := MakeMigration(c.Root, filename, "up", UP)
	if err != nil {
		return err
	}

	err = MakeMigration(c.Root, filename, "down", DOWN)
	if err != nil {
		return err
	}

	return nil
}

func MakeMigration(root, filename, direction, template string) error {
	dir := fmt.Sprintf("%s/internal/database/migrations", root)
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	path := fmt.Sprintf("%s/%s.%s.sql", dir, filename, direction)
	err := os.WriteFile(path, []byte(template), 0644)
	if err != nil {
		return err
	}

	return nil
}

func Up(c *config.Config) error {
	m, err := migrate.New("file://"+c.Root+"/internal/database/migrations", c.DatabaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Up(); err != nil {
		return err
	}

	return nil
}

func DownAll(c *config.Config) error {
	m, err := migrate.New("file://"+c.Root+"/internal/database/migrations", c.DatabaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Down(); err != nil {
		return err
	}

	return nil
}

func Steps(c *config.Config, steps int) error {
	m, err := migrate.New("file://"+c.Root+"/internal/database/migrations", c.DatabaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Steps(steps); err != nil {
		return err
	}

	return nil
}

func Force(c *config.Config) error {
	m, err := migrate.New("file://"+c.Root+"/internal/database/migrations", c.DatabaseURL)
	if err != nil {
		return err
	}
	defer m.Close()

	if err = m.Force(-1); err != nil {
		return err
	}

	return nil
}
