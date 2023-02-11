package migrator

import (
	"errors"
	"fmt"
	"os"
)

type Migrator interface {
	Make(root, name string) error
	Up(root string) error
	DownAll(root string) error
	Steps(root string, steps int) error
	Force(root string) error
}

func CreateMigrationFolderIfNotExists(dir string) error {
	if _, err := os.Stat(dir); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	return nil
}

func MakeMigration(root, filename, direction, template string) error {
	dir := fmt.Sprintf("%s/migrations", root)
	CreateMigrationFolderIfNotExists(dir)

	path := fmt.Sprintf("%s/%s.%s.sql", dir, filename, direction)
	err := os.WriteFile(path, []byte(template), 0644)
	if err != nil {
		return err
	}

	return nil
}
