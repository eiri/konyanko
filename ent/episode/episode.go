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
	// FieldNumber holds the string denoting the number field in the database.
	FieldNumber = "number"
	// FieldViewURL holds the string denoting the view_url field in the database.
	FieldViewURL = "view_url"
	// FieldDownloadURL holds the string denoting the download_url field in the database.
	FieldDownloadURL = "download_url"
	// FieldFileName holds the string denoting the file_name field in the database.
	FieldFileName = "file_name"
	// FieldFileSize holds the string denoting the file_size field in the database.
	FieldFileSize = "file_size"
	// FieldResolution holds the string denoting the resolution field in the database.
	FieldResolution = "resolution"
	// FieldVideoCodec holds the string denoting the video_codec field in the database.
	FieldVideoCodec = "video_codec"
	// FieldAudioCodec holds the string denoting the audio_codec field in the database.
	FieldAudioCodec = "audio_codec"
	// EdgeTitle holds the string denoting the title edge name in mutations.
	EdgeTitle = "title"
	// EdgeReleaseGroup holds the string denoting the release_group edge name in mutations.
	EdgeReleaseGroup = "release_group"
	// Table holds the table name of the episode in the database.
	Table = "episodes"
	// TitleTable is the table that holds the title relation/edge.
	TitleTable = "episodes"
	// TitleInverseTable is the table name for the Anime entity.
	// It exists in this package in order to avoid circular dependency with the "anime" package.
	TitleInverseTable = "animes"
	// TitleColumn is the table column denoting the title relation/edge.
	TitleColumn = "anime_id"
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
	FieldNumber,
	FieldViewURL,
	FieldDownloadURL,
	FieldFileName,
	FieldFileSize,
	FieldResolution,
	FieldVideoCodec,
	FieldAudioCodec,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "episodes"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"anime_id",
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
	// NumberValidator is a validator for the "number" field. It is called by the builders before save.
	NumberValidator func(int) error
	// FileSizeValidator is a validator for the "file_size" field. It is called by the builders before save.
	FileSizeValidator func(int) error
)

// OrderOption defines the ordering options for the Episode queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByNumber orders the results by the number field.
func ByNumber(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldNumber, opts...).ToFunc()
}

// ByViewURL orders the results by the view_url field.
func ByViewURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldViewURL, opts...).ToFunc()
}

// ByDownloadURL orders the results by the download_url field.
func ByDownloadURL(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDownloadURL, opts...).ToFunc()
}

// ByFileName orders the results by the file_name field.
func ByFileName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileName, opts...).ToFunc()
}

// ByFileSize orders the results by the file_size field.
func ByFileSize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFileSize, opts...).ToFunc()
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

// ByTitleField orders the results by title field.
func ByTitleField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTitleStep(), sql.OrderByField(field, opts...))
	}
}

// ByReleaseGroupField orders the results by release_group field.
func ByReleaseGroupField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newReleaseGroupStep(), sql.OrderByField(field, opts...))
	}
}
func newTitleStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TitleInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, TitleTable, TitleColumn),
	)
}
func newReleaseGroupStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ReleaseGroupInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, ReleaseGroupTable, ReleaseGroupColumn),
	)
}
