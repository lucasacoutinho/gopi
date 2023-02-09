package main

import (
	"github.com/lucasacoutinho/gopi/internal/config"
	"github.com/lucasacoutinho/gopi/internal/database/migration"
	"github.com/spf13/cobra"
)

func main() {
	c := config.Load()

	var make = &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Make(c, args[0])
			if err != nil {
				c.Logger.Errorw("make", "ERROR", err)
			}
		},
	}

	var migrate = &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Up(c)
			if err != nil {
				c.Logger.Errorw("migrate", "ERROR", err)
			}
		},
	}

	var rollback = &cobra.Command{
		Use: "rollback",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Steps(c, -1)
			if err != nil {
				c.Logger.Errorw("rollback", "ERROR", err)
			}
		},
	}

	var drop = &cobra.Command{
		Use: "drop",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.DownAll(c)
			if err != nil {
				c.Logger.Errorw("drop", "ERROR", err)
			}
		},
	}

	var fresh = &cobra.Command{
		Use: "fresh",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.DownAll(c)
			if err != nil {
				c.Logger.Errorw("fresh", "ERROR", err)
			}
			err = migration.Up(c)
			if err != nil {
				c.Logger.Errorw("fresh", "ERROR", err)
			}
		},
	}

	var force = &cobra.Command{
		Use: "force",
		Run: func(cmd *cobra.Command, args []string) {
			err := migration.Force(c)
			if err != nil {
				c.Logger.Errorw("force", "ERROR", err)
			}
			err = migration.Up(c)
			if err != nil {
				c.Logger.Errorw("force", "ERROR", err)
			}
		},
	}

	var root = &cobra.Command{Use: "database"}
	root.AddCommand(make)
	root.AddCommand(migrate)
	root.AddCommand(rollback)
	root.AddCommand(drop)
	root.AddCommand(fresh)
	root.AddCommand(force)
	root.Execute()
}
