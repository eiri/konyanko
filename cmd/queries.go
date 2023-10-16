package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/eiri/konyanko/ent"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/item"
	"github.com/spf13/cobra"
)

var (
	titlesCmd = &cobra.Command{
		Use:   "titles",
		Short: "Print all the available anime titles",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			titles, err := client.Anime.
				Query().
				Select(anime.FieldTitle).
				Strings(ctx)
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

	animeCmd = &cobra.Command{
		Use:   "anime",
		Short: "Print all the available episods in JSON format",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			animes, err := client.Anime.
				Query().
				WithEpisodes(func(q *ent.EpisodeQuery) {
					q.WithItem()
					q.WithReleaseGroup()
				}).
				All(ctx)
			if err != nil {
				return err
			}

			data, err := json.MarshalIndent(animes, "", "  ")
			if err != nil {
				return err
			}
			log.Printf("%s\n", data)
			return nil
		},
	}

	irregularCmd = &cobra.Command{
		Use:   "irregular",
		Short: "Print all not parsed feed items",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			items, err := client.Item.
				Query().
				Select(item.FieldFileName, item.FieldViewURL).
				Where(item.Not(item.HasEpisodes())).
				All(ctx)
			if err != nil {
				return err
			}

			data, err := json.MarshalIndent(items, "", "  ")
			if err != nil {
				return err
			}
			log.Printf("%s\n", data)
			return nil
		},
	}
)

func init() {
	listCmd.AddCommand(titlesCmd)
	listCmd.AddCommand(animeCmd)
	listCmd.AddCommand(irregularCmd)
}

/*
	titles, err := client.Anime.
		Query().
		WithEpisodes(func(q *ent.EpisodeQuery) {
			q.WithItem()
			q.WithReleaseGroup()
		}).
		All(ctx)
*/
