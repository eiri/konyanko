package main

import (
	"context"

	"github.com/eiri/konyanko/ent/migrate"
	"github.com/spf13/cobra"
)

// FIXME use https://github.com/muyo/sno as global id instead of WithGlobalUniqueID
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate ent schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Schema.Create(
			context.Background(),
			migrate.WithGlobalUniqueID(true),
		)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
