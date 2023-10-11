package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/eiri/konyanko/ent"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/irregular"
	"github.com/eiri/konyanko/ent/releasegroup"

	"github.com/dustin/go-humanize"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nssteinbrenner/anitogo"
	"github.com/zhengchun/syndfeed"
)

var (
	client  *ent.Client
	command string
)

func init() {
	flag.StringVar(&command, "command", "help", "command to execute")
	flag.StringVar(&command, "c", "help", "command to execute")
}

func main() {
	log.SetFlags(0)
	flag.Parse()

	if command == "help" {
		log.Println("available commands: help, migrate, import, list")
		os.Exit(0)
	}

	uri := "file:nyaa.db?mode=rwc&cache=shared&_fk=1"
	var err error
	client, err = ent.Open("sqlite3", uri)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	switch command {
	case "migrate":
		err := Migrate()
		if err != nil {
			log.Fatalf("failed creating schema resources: %v", err)
		}
	case "import":
		err := Import()
		if err != nil {
			log.Fatalf("failed to import episodes: %v", err)
		}
	case "list":
		err := List()
		if err != nil {
			log.Fatalf("failed to list episodes: %v", err)
		}
	default:
		log.Fatalf("unknown command: %s", command)
	}
}

func Migrate() error {
	// Run the auto migration tool.
	return client.Schema.Create(context.Background())
}

func Import() error {
	rssFile := "nyaa.xml"
	r, err := os.Open(rssFile)
	if err != nil {
		return fmt.Errorf("failed opening RSS file: %v", err)
	}
	defer r.Close()

	feed, _ := syndfeed.ParseRSS(r)
	for _, item := range feed.Items {
		ctx := context.Background()
		_, err := CreateEpisode(ctx, client, item)
		if err != nil {
			return fmt.Errorf("failed to add an episode: %v", err)
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

	e := anitogo.Parse(item.Title, anitogo.DefaultOptions)
	if e.AnimeTitle == "" {
		if c, err := client.Irregular.Query().Where(irregular.ViewURL(item.Id)).Aggregate(ent.Count()).Int(ctx); c == 0 && err == nil {
			_, err := client.Irregular.
				Create().
				SetViewURL(viewURL).
				SetDownloadURL(downloadURL).
				SetFileName(fileName).
				SetFileSize(fileSize).
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
		SetFileSize(fileSize)

	//FIXME! if we have AnimeSeason+AnimeSeasonPrefix but don't have EpisodeNumber assume this is a batch
	if len(e.EpisodeNumber) > 0 {
		if n, err := strconv.Atoi(e.EpisodeNumber[0]); err == nil {
			episode = episode.SetNumber(n)
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

func List() error {
	ctx := context.Background()
	episodes, err := client.Episode.
		Query().
		WithTitle().
		WithReleaseGroup().
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying episodes: %w", err)
	}

	for _, e := range episodes {
		data, err := json.MarshalIndent(e, "", "  ")
		if err != nil {
			return err
		}
		log.Printf("%s\n", data)
	}
	return nil
}
