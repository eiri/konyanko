package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/eiri/konyanko/ent"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Print all the available episods in JSON format",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		titles, err := client.Anime.
			Query().
			WithEpisodes(func(q *ent.EpisodeQuery) {
				q.WithItem()
				q.WithReleaseGroup()
			}).
			All(ctx)
		if err != nil {
			return err
		}

		data, err := json.MarshalIndent(titles, "", "  ")
		if err != nil {
			return err
		}
		log.Printf("%s\n", data)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
