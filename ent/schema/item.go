package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Item holds the schema definition for the Item entity.
type Item struct {
	ent.Schema
}

// Fields of the Item.
func (Item) Fields() []ent.Field {
	return []ent.Field{
		field.String("view_url").Unique(),
		field.String("download_url").Unique(),
		field.String("file_name").NotEmpty(),
		field.Int("file_size").Positive(),
		field.Time("publish_date").Default(time.Now).Nillable(),
	}
}

// Edges of the Item.
func (Item) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("episode", Episode.Type).StorageKey(edge.Column("item_id")).Unique(),
	}
}

// Indices of the Item.
func (Item) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("view_url"),
	}
}

// Annotations of the Item.
func (Item) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
	}
}
