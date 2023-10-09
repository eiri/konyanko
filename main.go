package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/eiri/konyanko/ent"

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

	feed, _ := syndfeed.ParseRSS(r)
	for _, item := range feed.Items {
		log.Println(item.Title)
		log.Println(item.Links[0].URL)
		log.Println(item.Id)
		for _, ext := range item.ElementExtensions {
			if ext.Name == "size" {
				log.Println(ext.Value)
			}
		}
		ctx := context.Background()
		e, err := CreateEpisode(ctx, client, item)
		if err != nil {
			return fmt.Errorf("failed to add an episode: %v", err)
		}
		log.Println("episode was created: ", e)
	}
	return nil
}

func CreateEpisode(ctx context.Context, client *ent.Client, item *syndfeed.Item) (*ent.Episode, error) {
	return client.Episode.
		Create().
		SetFileName(item.Title).
		Save(ctx)
}

func List() error {
	ctx := context.Background()
	episodes, err := client.Episode.
		Query().
		All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying episodes: %w", err)
	}

	for _, e := range episodes {
		pe := anitogo.Parse(e.FileName, anitogo.DefaultOptions)
		log.Println("FileName", pe.FileName)
		log.Println("AnimeTitle", pe.AnimeTitle)
		log.Println("EpisodeNumber", pe.EpisodeNumber)
		log.Println("ReleaseGroup", pe.ReleaseGroup)
		log.Println("VideoResolution", pe.VideoResolution)
		log.Println("VideoTerm", pe.VideoTerm)
		log.Println("AudioTerm", pe.AudioTerm)
		log.Println("---")
	}
	return nil
}
