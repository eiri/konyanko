package main

import (
	"context"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/eiri/konyanko/ent"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/releasegroup"

	"github.com/dustin/go-humanize"
	"github.com/nssteinbrenner/anitogo"
	"github.com/spf13/cobra"
	"github.com/zhengchun/syndfeed"
)

var (
	rssFile string
	r       io.ReadCloser

	importCmd = &cobra.Command{
		Use:   "import",
		Short: "Import all items from RSS feed in the database",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			var err error
			r, err = os.Open(rssFile)
			return err
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			feed, err := syndfeed.ParseRSS(r)
			if err != nil {
				return err
			}
			for _, item := range feed.Items {
				if err := CreateEpisode(item); err != nil {
					return err
				}
			}
			return nil
		},
		PostRun: func(cmd *cobra.Command, args []string) {
			r.Close()
		},
	}
)

func init() {
	importCmd.Flags().StringVarP(&rssFile, "file", "f", "nyaa.xml", "RSS file")
	importCmd.MarkFlagFilename("file", ".xml")

	rootCmd.AddCommand(importCmd)
}

func CreateEpisode(i *syndfeed.Item) error {
	ctx := context.Background()

	if c, err := client.Item.Query().Where(item.ViewURL(i.Id)).Aggregate(ent.Count()).Int(ctx); c == 1 && err == nil {
		log.Printf("Found %q, skip\n", i.Title)
		// should be const error EpisodeExists or something...
		return nil
	}

	log.Printf("Adding %q\n", i.Title)

	viewURL := i.Id
	downloadURL := i.Links[0].URL
	fileName := i.Title
	fileSize := 0
	for _, ext := range i.ElementExtensions {
		if ext.Name == "size" {
			fs, _ := humanize.ParseBytes(ext.Value)
			fileSize = int(fs)
			break
		}
	}
	publishDate := i.PublishDate

	item, err := client.Item.
		Create().
		SetViewURL(viewURL).
		SetDownloadURL(downloadURL).
		SetFileName(fileName).
		SetFileSize(fileSize).
		SetPublishDate(publishDate).
		Save(ctx)
	if err != nil {
		return err
	}

	e := anitogo.Parse(item.FileName, anitogo.DefaultOptions)
	if e.AnimeTitle == "" {
		// FIXME! same as above, should be a predefined error
		return nil
	}

	//FIXME!! obvsly needs transaction and anime/release_group cache

	anime, err := client.Anime.
		Query().
		Where(anime.Title(e.AnimeTitle)).
		Only(ctx)
	switch {
	case ent.IsNotFound(err):
		anime, err = client.Anime.
			Create().
			SetTitle(e.AnimeTitle).
			Save(ctx)
		if err != nil {
			return err
		}
	case err != nil:
		return err
	}

	episode := client.Episode.
		Create().
		SetItem(item).
		SetTitle(anime)

	if e.ReleaseGroup != "" {
		releaseGroup, err := client.ReleaseGroup.
			Query().
			Where(releasegroup.Name(e.ReleaseGroup)).
			Only(ctx)
		switch {
		case ent.IsNotFound(err):
			releaseGroup, err = client.ReleaseGroup.
				Create().
				SetName(e.ReleaseGroup).
				Save(ctx)
			if err != nil {
				return err
			}
		case err != nil:
			return err
		}
		episode = episode.SetReleaseGroup(releaseGroup)
	}

	//FIXME! if we have AnimeSeason+AnimeSeasonPrefix but don't have EpisodeNumber assume this is a batch
	if len(e.EpisodeNumber) > 0 {
		if n, err := strconv.Atoi(e.EpisodeNumber[0]); err == nil {
			episode = episode.SetEpisodeNumber(n)
		}
	}

	if len(e.AnimeSeason) > 0 {
		if n, err := strconv.Atoi(e.AnimeSeason[0]); err == nil {
			episode = episode.SetAnimeSeason(n)
		}
	}

	if e.VideoResolution != "" {
		episode = episode.SetResolution(e.VideoResolution)
	}

	if len(e.VideoTerm) > 0 {
		episode = episode.SetVideoCodec(e.VideoTerm[0])
	}

	if len(e.AudioTerm) > 0 {
		episode = episode.SetAudioCodec(e.AudioTerm[0])
	}

	_, err = episode.Save(ctx)

	return err
}
