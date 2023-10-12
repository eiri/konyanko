package main

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/eiri/konyanko/ent"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/irregular"
	"github.com/eiri/konyanko/ent/releasegroup"

	"github.com/dustin/go-humanize"
	"github.com/nssteinbrenner/anitogo"
	"github.com/spf13/cobra"
	"github.com/zhengchun/syndfeed"
)

var (
	rssFile string

	importCmd = &cobra.Command{
		Use:   "import",
		Short: "Import all items from RSS feed in the database",
		RunE:  importRSS,
	}
)

func init() {
	importCmd.Flags().StringVarP(&rssFile, "file", "f", "nyaa.xml", "RSS file")
	importCmd.MarkFlagFilename("file", ".xml")

	rootCmd.AddCommand(importCmd)
}

func importRSS(cmd *cobra.Command, args []string) error {
	// pre and past conditions
	r, err := os.Open(rssFile)
	if err != nil {
		return err
	}
	defer r.Close()

	feed, err := syndfeed.ParseRSS(r)
	if err != nil {
		return err
	}
	for _, item := range feed.Items {
		ctx := context.Background()
		_, err := CreateEpisode(ctx, client, item)
		if err != nil {
			return err
		}
	}

	return nil
}

func CreateEpisode(ctx context.Context, client *ent.Client, item *syndfeed.Item) (*ent.Episode, error) {
	if c, err := client.Episode.Query().Where(episode.ViewURL(item.Id)).Aggregate(ent.Count()).Int(ctx); c == 1 && err == nil {
		log.Printf("Found %q, skip\n", item.Title)
		return nil, nil
	}

	log.Printf("Adding %q\n", item.Title)

	fileName := item.Title
	fileSize := 0
	for _, ext := range item.ElementExtensions {
		if ext.Name == "size" {
			fs, _ := humanize.ParseBytes(ext.Value)
			fileSize = int(fs)
			break
		}
	}
	viewURL := item.Id
	downloadURL := item.Links[0].URL
	publishDate := item.PublishDate

	e := anitogo.Parse(item.Title, anitogo.DefaultOptions)
	if e.AnimeTitle == "" {
		if c, err := client.Irregular.Query().Where(irregular.ViewURL(item.Id)).Aggregate(ent.Count()).Int(ctx); c == 0 && err == nil {
			_, err := client.Irregular.
				Create().
				SetViewURL(viewURL).
				SetDownloadURL(downloadURL).
				SetFileName(fileName).
				SetFileSize(fileSize).
				SetPublishDate(publishDate).
				Save(ctx)
			if err != nil {
				return nil, err
			}
		}
		return nil, nil
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
			return nil, err
		}
	case err != nil:
		return nil, err
	}

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
			return nil, err
		}
	case err != nil:
		return nil, err
	}

	episode := client.Episode.
		Create().
		SetTitle(anime).
		SetReleaseGroup(releaseGroup).
		SetViewURL(viewURL).
		SetDownloadURL(downloadURL).
		SetFileName(fileName).
		SetFileSize(fileSize).
		SetPublishDate(publishDate)

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

	return episode.Save(ctx)
}
