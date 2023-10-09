// Code generated by ent, DO NOT EDIT.

package anime

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the anime type in the database.
	Label = "anime"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// EdgeEpisodes holds the string denoting the episodes edge name in mutations.
	EdgeEpisodes = "episodes"
	// Table holds the table name of the anime in the database.
	Table = "animes"
	// EpisodesTable is the table that holds the episodes relation/edge.
	EpisodesTable = "episodes"
	// EpisodesInverseTable is the table name for the Episode entity.
	// It exists in this package in order to avoid circular dependency with the "episode" package.
	EpisodesInverseTable = "episodes"
	// EpisodesColumn is the table column denoting the episodes relation/edge.
	EpisodesColumn = "anime_id"
)

// Columns holds all SQL columns for anime fields.
var Columns = []string{
	FieldID,
	FieldTitle,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the Anime queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByEpisodesCount orders the results by episodes count.
func ByEpisodesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newEpisodesStep(), opts...)
	}
}

// ByEpisodes orders the results by episodes terms.
func ByEpisodes(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newEpisodesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newEpisodesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(EpisodesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, EpisodesTable, EpisodesColumn),
	)
}
