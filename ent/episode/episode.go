// Code generated by ent, DO NOT EDIT.

package episode

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the episode type in the database.
	Label = "episode"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEpisodeNumber holds the string denoting the episode_number field in the database.
	FieldEpisodeNumber = "episode_number"
	// FieldAnimeSeason holds the string denoting the anime_season field in the database.
	FieldAnimeSeason = "anime_season"
	// FieldResolution holds the string denoting the resolution field in the database.
	FieldResolution = "resolution"
	// FieldVideoCodec holds the string denoting the video_codec field in the database.
	FieldVideoCodec = "video_codec"
	// FieldAudioCodec holds the string denoting the audio_codec field in the database.
	FieldAudioCodec = "audio_codec"
	// EdgeItem holds the string denoting the item edge name in mutations.
	EdgeItem = "item"
	// EdgeAnime holds the string denoting the anime edge name in mutations.
	EdgeAnime = "anime"
	// EdgeReleaseGroup holds the string denoting the release_group edge name in mutations.
	EdgeReleaseGroup = "release_group"
	// Table holds the table name of the episode in the database.
	Table = "episodes"
	// ItemTable is the table that holds the item relation/edge.
	ItemTable = "episodes"
	// ItemInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemInverseTable = "items"
	// ItemColumn is the table column denoting the item relation/edge.
	ItemColumn = "item_id"
	// AnimeTable is the table that holds the anime relation/edge.
	AnimeTable = "episodes"
	// AnimeInverseTable is the table name for the Anime entity.
	// It exists in this package in order to avoid circular dependency with the "anime" package.
	AnimeInverseTable = "animes"
	// AnimeColumn is the table column denoting the anime relation/edge.
	AnimeColumn = "anime_id"
	// ReleaseGroupTable is the table that holds the release_group relation/edge.
	ReleaseGroupTable = "episodes"
	// ReleaseGroupInverseTable is the table name for the ReleaseGroup entity.
	// It exists in this package in order to avoid circular dependency with the "releasegroup" package.
	ReleaseGroupInverseTable = "release_groups"
	// ReleaseGroupColumn is the table column denoting the release_group relation/edge.
	ReleaseGroupColumn = "release_group_id"
)

// Columns holds all SQL columns for episode fields.
var Columns = []string{
	FieldID,
	FieldEpisodeNumber,
	FieldAnimeSeason,
	FieldResolution,
	FieldVideoCodec,
	FieldAudioCodec,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "episodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"anime_id",
	"item_id",
	"release_group_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultEpisodeNumber holds the default value on creation for the "episode_number" field.
	DefaultEpisodeNumber int
	// EpisodeNumberValidator is a validator for the "episode_number" field. It is called by the builders before save.
	EpisodeNumberValidator func(int) error
	// DefaultAnimeSeason holds the default value on creation for the "anime_season" field.
	DefaultAnimeSeason int
	// AnimeSeasonValidator is a validator for the "anime_season" field. It is called by the builders before save.
	AnimeSeasonValidator func(int) error
)

// OrderOption defines the ordering options for the Episode queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByEpisodeNumber orders the results by the episode_number field.
func ByEpisodeNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldEpisodeNumber, opts...).ToFunc()
}

// ByAnimeSeason orders the results by the anime_season field.
func ByAnimeSeason(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAnimeSeason, opts...).ToFunc()
}

// ByResolution orders the results by the resolution field.
func ByResolution(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldResolution, opts...).ToFunc()
}

// ByVideoCodec orders the results by the video_codec field.
func ByVideoCodec(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldVideoCodec, opts...).ToFunc()
}

// ByAudioCodec orders the results by the audio_codec field.
func ByAudioCodec(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAudioCodec, opts...).ToFunc()
}

// ByItemField orders the results by item field.
func ByItemField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newItemStep(), sql.OrderByField(field, opts...))
	}
}

// ByAnimeField orders the results by anime field.
func ByAnimeField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAnimeStep(), sql.OrderByField(field, opts...))
	}
}

// ByReleaseGroupField orders the results by release_group field.
func ByReleaseGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReleaseGroupStep(), sql.OrderByField(field, opts...))
	}
}
func newItemStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ItemInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, true, ItemTable, ItemColumn),
	)
}
func newAnimeStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AnimeInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AnimeTable, AnimeColumn),
	)
}
func newReleaseGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReleaseGroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ReleaseGroupTable, ReleaseGroupColumn),
	)
}
