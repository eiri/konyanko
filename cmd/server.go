package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/eiri/konyanko/ent"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/ogent"
	"github.com/go-faster/jx"
	"github.com/spf13/cobra"
)

var (
	port int

	serverCmd = &cobra.Command{
		Use:   "server",
		Short: "Run UI and API server",
		RunE:  serverRunner,
	}
)

func init() {
	importCmd.Flags().IntVarP(&port, "port", "p", 2315, "A port to run the server on")

	rootCmd.AddCommand(serverCmd)
}

type handler struct {
	*ogent.OgentHandler
	client *ent.Client
}

func (h handler) ListItemByDate(ctx context.Context, params ogent.ListItemByDateParams) (ogent.ListItemByDateRes, error) {
	q := h.client.Item.Query()
	// get date from path
	q.Where(
		item.And(
			item.PublishDateGTE(params.Day),
			item.PublishDateLT(params.Day.AddDate(0, 0, 1)),
		),
	)
	// set edges and then query
	es, err := q.WithEpisode(func(q *ent.EpisodeQuery) {
		q.WithTitle()
		q.WithReleaseGroup()
	}).All(ctx)
	if err != nil {
		var e jx.Encoder
		e.Str(err.Error())
		switch {
		case ent.IsNotFound(err):
			return &ogent.R404{
				Code:   http.StatusNotFound,
				Status: http.StatusText(http.StatusNotFound),
				Errors: e.Bytes(),
			}, nil
		case ent.IsNotSingular(err):
			return &ogent.R409{
				Code:   http.StatusConflict,
				Status: http.StatusText(http.StatusConflict),
				Errors: e.Bytes(),
			}, nil
		default:
			// Let the server handle the error.
			return nil, err
		}
	}

	r := make([]ogent.ItemByDateItem, 0)
	for _, e := range es {
		itemByDate := ogent.ItemByDateItem{
			ID:          e.ID,
			ViewURL:     e.ViewURL,
			DownloadURL: e.DownloadURL,
			FileName:    e.FileName,
			FileSize:    e.FileSize,
			PublishDate: ogent.NewOptDateTime(*e.PublishDate),
		}
		if e.Edges.Episode != nil {
			itemByDateEpisode := ogent.NewEpisodeList(e.Edges.Episode)
			itemByDate.Episode = ogent.NewOptEpisodeList(*itemByDateEpisode)
		}

		if e.Edges.Episode != nil && e.Edges.Episode.Edges.Title != nil {
			itemByDateAnime := ogent.NewAnimeList(e.Edges.Episode.Edges.Title)
			itemByDate.Anime = ogent.NewOptAnimeList(*itemByDateAnime)
		}

		if e.Edges.Episode != nil && e.Edges.Episode.Edges.ReleaseGroup != nil {
			itemByDateReleaseGroup := ogent.NewReleaseGroupList(e.Edges.Episode.Edges.ReleaseGroup)
			itemByDate.ReleaseGroup = ogent.NewOptReleaseGroupList(*itemByDateReleaseGroup)
		}
		r = append(r, itemByDate)
	}

	return (*ogent.ListItemByDateOKApplicationJSON)(&r), nil
}

func serverRunner(cmd *cobra.Command, args []string) error {
	h := handler{
		OgentHandler: ogent.NewOgentHandler(client),
		client:       client,
	}
	srv, err := ogent.NewServer(h)
	if err != nil {
		return err
	}
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), srv); err != nil {
		return err
	}
	return nil
}
