// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/releasegroup"
)

// Episode is the model entity for the Episode schema.
type Episode struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// EpisodeNumber holds the value of the "episode_number" field.
	EpisodeNumber int `json:"episode_number,omitempty"`
	// AnimeSeason holds the value of the "anime_season" field.
	AnimeSeason int `json:"anime_season,omitempty"`
	// Resolution holds the value of the "resolution" field.
	Resolution *string `json:"resolution,omitempty"`
	// VideoCodec holds the value of the "video_codec" field.
	VideoCodec *string `json:"video_codec,omitempty"`
	// AudioCodec holds the value of the "audio_codec" field.
	AudioCodec *string `json:"audio_codec,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EpisodeQuery when eager-loading is set.
	Edges            EpisodeEdges `json:"edges"`
	anime_id         *int
	item_id          *int
	release_group_id *int
	selectValues     sql.SelectValues
}

// EpisodeEdges holds the relations/edges for other nodes in the graph.
type EpisodeEdges struct {
	// Item holds the value of the item edge.
	Item *Item `json:"item,omitempty"`
	// Anime holds the value of the anime edge.
	Anime *Anime `json:"anime,omitempty"`
	// ReleaseGroup holds the value of the release_group edge.
	ReleaseGroup *ReleaseGroup `json:"release_group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
	// totalCount holds the count of the edges above.
	totalCount [3]map[string]int
}

// ItemOrErr returns the Item value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EpisodeEdges) ItemOrErr() (*Item, error) {
	if e.loadedTypes[0] {
		if e.Item == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: item.Label}
		}
		return e.Item, nil
	}
	return nil, &NotLoadedError{edge: "item"}
}

// AnimeOrErr returns the Anime value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EpisodeEdges) AnimeOrErr() (*Anime, error) {
	if e.loadedTypes[1] {
		if e.Anime == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: anime.Label}
		}
		return e.Anime, nil
	}
	return nil, &NotLoadedError{edge: "anime"}
}

// ReleaseGroupOrErr returns the ReleaseGroup value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EpisodeEdges) ReleaseGroupOrErr() (*ReleaseGroup, error) {
	if e.loadedTypes[2] {
		if e.ReleaseGroup == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: releasegroup.Label}
		}
		return e.ReleaseGroup, nil
	}
	return nil, &NotLoadedError{edge: "release_group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Episode) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case episode.FieldID, episode.FieldEpisodeNumber, episode.FieldAnimeSeason:
			values[i] = new(sql.NullInt64)
		case episode.FieldResolution, episode.FieldVideoCodec, episode.FieldAudioCodec:
			values[i] = new(sql.NullString)
		case episode.ForeignKeys[0]: // anime_id
			values[i] = new(sql.NullInt64)
		case episode.ForeignKeys[1]: // item_id
			values[i] = new(sql.NullInt64)
		case episode.ForeignKeys[2]: // release_group_id
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Episode fields.
func (e *Episode) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case episode.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case episode.FieldEpisodeNumber:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field episode_number", values[i])
			} else if value.Valid {
				e.EpisodeNumber = int(value.Int64)
			}
		case episode.FieldAnimeSeason:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field anime_season", values[i])
			} else if value.Valid {
				e.AnimeSeason = int(value.Int64)
			}
		case episode.FieldResolution:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field resolution", values[i])
			} else if value.Valid {
				e.Resolution = new(string)
				*e.Resolution = value.String
			}
		case episode.FieldVideoCodec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field video_codec", values[i])
			} else if value.Valid {
				e.VideoCodec = new(string)
				*e.VideoCodec = value.String
			}
		case episode.FieldAudioCodec:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field audio_codec", values[i])
			} else if value.Valid {
				e.AudioCodec = new(string)
				*e.AudioCodec = value.String
			}
		case episode.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field anime_id", value)
			} else if value.Valid {
				e.anime_id = new(int)
				*e.anime_id = int(value.Int64)
			}
		case episode.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field item_id", value)
			} else if value.Valid {
				e.item_id = new(int)
				*e.item_id = int(value.Int64)
			}
		case episode.ForeignKeys[2]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field release_group_id", value)
			} else if value.Valid {
				e.release_group_id = new(int)
				*e.release_group_id = int(value.Int64)
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Episode.
// This includes values selected through modifiers, order, etc.
func (e *Episode) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryItem queries the "item" edge of the Episode entity.
func (e *Episode) QueryItem() *ItemQuery {
	return NewEpisodeClient(e.config).QueryItem(e)
}

// QueryAnime queries the "anime" edge of the Episode entity.
func (e *Episode) QueryAnime() *AnimeQuery {
	return NewEpisodeClient(e.config).QueryAnime(e)
}

// QueryReleaseGroup queries the "release_group" edge of the Episode entity.
func (e *Episode) QueryReleaseGroup() *ReleaseGroupQuery {
	return NewEpisodeClient(e.config).QueryReleaseGroup(e)
}

// Update returns a builder for updating this Episode.
// Note that you need to call Episode.Unwrap() before calling this method if this Episode
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Episode) Update() *EpisodeUpdateOne {
	return NewEpisodeClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Episode entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Episode) Unwrap() *Episode {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Episode is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Episode) String() string {
	var builder strings.Builder
	builder.WriteString("Episode(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("episode_number=")
	builder.WriteString(fmt.Sprintf("%v", e.EpisodeNumber))
	builder.WriteString(", ")
	builder.WriteString("anime_season=")
	builder.WriteString(fmt.Sprintf("%v", e.AnimeSeason))
	builder.WriteString(", ")
	if v := e.Resolution; v != nil {
		builder.WriteString("resolution=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := e.VideoCodec; v != nil {
		builder.WriteString("video_codec=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	if v := e.AudioCodec; v != nil {
		builder.WriteString("audio_codec=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Episodes is a parsable slice of Episode.
type Episodes []*Episode
