package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all the available episods in JSON format",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		episodes, err := client.Episode.
			Query().
			WithTitle().
			WithReleaseGroup().
			All(ctx)
		if err != nil {
			return err
		}

		for _, e := range episodes {
			data, err := json.MarshalIndent(e, "", "  ")
			if err != nil {
				return err
			}
			log.Printf("%s\n", data)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
