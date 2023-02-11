package main

import (
	"time"

	"github.com/lucasacoutinho/gopi/internal/config"
	"github.com/lucasacoutinho/gopi/internal/database/migrator"
	"github.com/spf13/cobra"
)

const toolName = "migrate-cli"

func main() {
	c := config.Load()
	if c.DatabaseDriver != "pgx" {
		c.Logger.Errorln("Database driver not supported")
		return
	}

	migrator := migrator.NewPGX(c.DatabaseURL)
	Init(c, migrator)
}

func Init(c *config.Config, migrator migrator.Migrator) {
	make := &cobra.Command{
		Use:   "make",
		Short: "Create a new migration",
		Long:  "Create a new migration",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Creating migration")
			err := migrator.Make(c.Root, args[0])
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Migration created")
			if err != nil {
				c.Logger.Errorw("make", "ERROR", err)
			}
		},
	}

	migrate := &cobra.Command{
		Use:   "migrate",
		Short: "Run all migrations",
		Long:  "Run all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Running migrations")
			err := migrator.Up(c.Root)
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Migrations ran")
			if err != nil {
				c.Logger.Errorw("migrate", "ERROR", err)
			}
		},
	}

	rollback := &cobra.Command{
		Use:   "rollback",
		Short: "Rollback the last migration",
		Long:  "Rollback the last migration",
		Run: func(cmd *cobra.Command, args []string) {
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Rolling back")
			err := migrator.Steps(c.Root, -1)
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Rolled back")
			if err != nil {
				c.Logger.Errorw("rollback", "ERROR", err)
			}
		},
	}

	drop := &cobra.Command{
		Use:   "drop",
		Short: "Drop all tables",
		Long:  "Drop all tables",
		Run: func(cmd *cobra.Command, args []string) {
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Dropping tables")
			err := migrator.DownAll(c.Root)
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Tables dropped")
			if err != nil {
				c.Logger.Errorw("drop", "ERROR", err)
			}
		},
	}

	fresh := &cobra.Command{
		Use:   "fresh",
		Short: "Drop all tables and run all migrations",
		Long:  "Drop all tables and run all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Dropping tables")
			err := migrator.DownAll(c.Root)
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Tables dropped")
			if err != nil {
				c.Logger.Errorw("fresh", "ERROR", err)
			}
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Running migrations")
			err = migrator.Up(c.Root)
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Migrations ran")
			if err != nil {
				c.Logger.Errorw("fresh", "ERROR", err)
			}
		},
	}

	force := &cobra.Command{
		Use:   "force",
		Short: "Reset dirty migrations",
		Long:  "Reset dirty migrations and run all migrations",
		Run: func(cmd *cobra.Command, args []string) {
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Resetting dirty migrations")
			err := migrator.Force(c.Root)
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Dirty migrations reset")
			if err != nil {
				c.Logger.Errorw("force", "ERROR", err)
			}
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Running migrations")
			err = migrator.Up(c.Root)
			c.Logger.Infow(c.AppName, toolName, time.Now().String(), "Migrations ran")
			if err != nil {
				c.Logger.Errorw("force", "ERROR", err)
			}
		},
	}

	var root = &cobra.Command{
		Use:   toolName,
		Short: "Root command for our application",
		Long:  "Root command for our application, the main purpose is to help setup subcommands",
	}
	root.AddCommand(make)
	root.AddCommand(migrate)
	root.AddCommand(rollback)
	root.AddCommand(drop)
	root.AddCommand(fresh)
	root.AddCommand(force)
	root.Execute()
}
