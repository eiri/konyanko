package main

import (
	"context"

	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate ent schema",
	RunE: func(cmd *cobra.Command, args []string) error {
		return client.Schema.Create(context.Background())
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
