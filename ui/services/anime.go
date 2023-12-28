package services

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"sort"
	"strconv"

	"github.com/eiri/konyanko/ent"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ui/components"
)

type Anime struct {
	Client *ent.Client
}

func NewAnime(c *ent.Client) Anime {
	return Anime{Client: c}
}

func (a Anime) Shows(ctx context.Context) ([]components.Anime, error) {
	resolution := "1080p"
	shows := make([]components.Anime, 0)
	// 1080p, 720p, 480p
	// also nil, 2160p
	animes, err := a.Client.Anime.
		Query().
		Where(anime.HasEpisodesWith(episode.Resolution(resolution))).
		WithEpisodes(func(q *ent.EpisodeQuery) {
			q.Where(episode.Resolution(resolution))
			q.WithItem()
			q.WithReleaseGroup()
		}).
		All(ctx)

	for _, a := range animes {
		show := components.Anime{
			Title:    a.Title,
			Episodes: make([]components.Episode, 0),
		}
		grouped := make(map[string]components.Episode)
		for _, e := range a.Edges.Episodes {
			season := strconv.Itoa(e.AnimeSeason)
			number := strconv.Itoa(e.EpisodeNumber)
			hash := md5.Sum([]byte(season + number))
			key := hex.EncodeToString(hash[:])
			episode, ok := grouped[key]
			if !ok {
				episode = components.Episode{
					Season:   season,
					Number:   number,
					Torrents: make([]components.Torrent, 0),
				}
			}
			torrent := components.Torrent{
				URL:         e.Edges.Item.DownloadURL,
				ViewURL:     e.Edges.Item.ViewURL,
				FileSize:    e.Edges.Item.FileSize,
				PublishDate: *e.Edges.Item.PublishDate,
			}
			//FIXME - we should just assign published and if not Anonymous to it
			if e.Edges.ReleaseGroup != nil {
				torrent.ReleaseGroup = e.Edges.ReleaseGroup.Name
			}
			episode.Torrents = append(episode.Torrents, torrent)

			grouped[key] = episode
		}

		for _, episode := range grouped {
			show.Episodes = append(show.Episodes, episode)
		}

		sort.Slice(show.Episodes, func(i, j int) bool {
			if show.Episodes[i].Season == show.Episodes[j].Season {
				return show.Episodes[i].Number > show.Episodes[j].Number
			}
			return show.Episodes[i].Season > show.Episodes[j].Season
		})

		shows = append(shows, show)
	}

	sort.Slice(shows, func(i, j int) bool {
		return shows[i].Title < shows[j].Title
	})

	return shows, err
}
