// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/eiri/konyanko/ent/anime"
)

// Anime is the model entity for the Anime schema.
type Anime struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AnimeQuery when eager-loading is set.
	Edges        AnimeEdges `json:"edges"`
	selectValues sql.SelectValues
}

// AnimeEdges holds the relations/edges for other nodes in the graph.
type AnimeEdges struct {
	// Episodes holds the value of the episodes edge.
	Episodes []*Episode `json:"episodes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// EpisodesOrErr returns the Episodes value or an error if the edge
// was not loaded in eager-loading.
func (e AnimeEdges) EpisodesOrErr() ([]*Episode, error) {
	if e.loadedTypes[0] {
		return e.Episodes, nil
	}
	return nil, &NotLoadedError{edge: "episodes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Anime) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case anime.FieldID:
			values[i] = new(sql.NullInt64)
		case anime.FieldTitle:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Anime fields.
func (a *Anime) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case anime.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case anime.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				a.Title = value.String
			}
		default:
			a.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Anime.
// This includes values selected through modifiers, order, etc.
func (a *Anime) Value(name string) (ent.Value, error) {
	return a.selectValues.Get(name)
}

// QueryEpisodes queries the "episodes" edge of the Anime entity.
func (a *Anime) QueryEpisodes() *EpisodeQuery {
	return NewAnimeClient(a.config).QueryEpisodes(a)
}

// Update returns a builder for updating this Anime.
// Note that you need to call Anime.Unwrap() before calling this method if this Anime
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Anime) Update() *AnimeUpdateOne {
	return NewAnimeClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Anime entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Anime) Unwrap() *Anime {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Anime is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Anime) String() string {
	var builder strings.Builder
	builder.WriteString("Anime(")
	builder.WriteString(fmt.Sprintf("id=%v, ", a.ID))
	builder.WriteString("title=")
	builder.WriteString(a.Title)
	builder.WriteByte(')')
	return builder.String()
}

// Animes is a parsable slice of Anime.
type Animes []*Anime
