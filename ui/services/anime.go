package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strconv"

	"github.com/dustin/go-humanize"

	"github.com/eiri/konyanko/ent"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ui/components"
)

type Anime struct {
	Client *ent.Client
}

func NewAnime(c *ent.Client) Anime {
	return Anime{Client: c}
}

func (a Anime) Episodes(ctx context.Context) ([]components.Episode, error) {
	episodes := make([]components.Episode, 0)
	// 1080p, 720p, 480p
	// also nil, 2160p
	entEpisodes, err := a.Client.Episode.
		Query().
		Where(episode.Resolution("1080p")).
		WithAnime().
		WithItem().
		WithReleaseGroup().
		All(ctx)

	grouped := make(map[string]components.Episode)
	for _, entEpisode := range entEpisodes {
		title := entEpisode.Edges.Anime.Title
		season := strconv.Itoa(entEpisode.AnimeSeason)
		number := strconv.Itoa(entEpisode.EpisodeNumber)
		hash := md5.Sum([]byte(title + season + number))
		key := hex.EncodeToString(hash[:])
		episode, ok := grouped[key]
		if !ok {
			episode = components.Episode{
				Title:      title,
				Season:     season,
				Number:     number,
				Resolution: *entEpisode.Resolution,
				Items:      make([]components.Item, 0),
			}
		}

		item := components.Item{
			URL:      entEpisode.Edges.Item.DownloadURL,
			ViewURL:  entEpisode.Edges.Item.ViewURL,
			FileSize: humanize.Bytes(uint64(entEpisode.Edges.Item.FileSize)),
		}
		//FIXME - we should just assign published and if not Anonymous to it
		if entEpisode.Edges.ReleaseGroup != nil {
			item.ReleaseGroup = entEpisode.Edges.ReleaseGroup.Name
		}
		episode.Items = append(episode.Items, item)

		grouped[key] = episode
	}

	for _, episode := range grouped {
		episodes = append(episodes, episode)
	}

	sort.Slice(episodes, func(i, j int) bool {
		return episodes[i].Title < episodes[j].Title
	})

	return episodes, err
}
