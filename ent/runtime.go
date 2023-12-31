// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/releasegroup"
	"github.com/eiri/konyanko/ent/schema"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	animeFields := schema.Anime{}.Fields()
	_ = animeFields
	// animeDescTitle is the schema descriptor for title field.
	animeDescTitle := animeFields[0].Descriptor()
	// anime.TitleValidator is a validator for the "title" field. It is called by the builders before save.
	anime.TitleValidator = animeDescTitle.Validators[0].(func(string) error)
	episodeFields := schema.Episode{}.Fields()
	_ = episodeFields
	// episodeDescEpisodeNumber is the schema descriptor for episode_number field.
	episodeDescEpisodeNumber := episodeFields[0].Descriptor()
	// episode.DefaultEpisodeNumber holds the default value on creation for the episode_number field.
	episode.DefaultEpisodeNumber = episodeDescEpisodeNumber.Default.(int)
	// episode.EpisodeNumberValidator is a validator for the "episode_number" field. It is called by the builders before save.
	episode.EpisodeNumberValidator = episodeDescEpisodeNumber.Validators[0].(func(int) error)
	// episodeDescAnimeSeason is the schema descriptor for anime_season field.
	episodeDescAnimeSeason := episodeFields[1].Descriptor()
	// episode.DefaultAnimeSeason holds the default value on creation for the anime_season field.
	episode.DefaultAnimeSeason = episodeDescAnimeSeason.Default.(int)
	// episode.AnimeSeasonValidator is a validator for the "anime_season" field. It is called by the builders before save.
	episode.AnimeSeasonValidator = episodeDescAnimeSeason.Validators[0].(func(int) error)
	itemFields := schema.Item{}.Fields()
	_ = itemFields
	// itemDescFileName is the schema descriptor for file_name field.
	itemDescFileName := itemFields[2].Descriptor()
	// item.FileNameValidator is a validator for the "file_name" field. It is called by the builders before save.
	item.FileNameValidator = itemDescFileName.Validators[0].(func(string) error)
	// itemDescFileSize is the schema descriptor for file_size field.
	itemDescFileSize := itemFields[3].Descriptor()
	// item.FileSizeValidator is a validator for the "file_size" field. It is called by the builders before save.
	item.FileSizeValidator = itemDescFileSize.Validators[0].(func(int) error)
	// itemDescPublishDate is the schema descriptor for publish_date field.
	itemDescPublishDate := itemFields[4].Descriptor()
	// item.DefaultPublishDate holds the default value on creation for the publish_date field.
	item.DefaultPublishDate = itemDescPublishDate.Default.(func() time.Time)
	releasegroupFields := schema.ReleaseGroup{}.Fields()
	_ = releasegroupFields
	// releasegroupDescName is the schema descriptor for name field.
	releasegroupDescName := releasegroupFields[0].Descriptor()
	// releasegroup.NameValidator is a validator for the "name" field. It is called by the builders before save.
	releasegroup.NameValidator = releasegroupDescName.Validators[0].(func(string) error)
}
