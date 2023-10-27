package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ReleaseGroup holds the schema definition for the ReleaseGroup entity.
type ReleaseGroup struct {
	ent.Schema
}

// Fields of the ReleaseGroup.
func (ReleaseGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty().Unique(),
	}
}

// Edges of the ReleaseGroup.
func (ReleaseGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("episodes", Episode.Type).StorageKey(edge.Column("release_group_id")),
	}
}

// Annotations of the ReleaseGroup.
func (ReleaseGroup) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
