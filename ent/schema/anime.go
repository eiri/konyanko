package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Anime holds the schema definition for the Anime entity.
type Anime struct {
	ent.Schema
}

// Fields of the Anime.
func (Anime) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty().Unique(),
	}
}

// Edges of the Anime.
func (Anime) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("episodes", Episode.Type).StorageKey(edge.Column("anime_id")),
	}
}

// Annotations of the Anime.
func (Anime) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
