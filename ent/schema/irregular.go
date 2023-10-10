package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Irregular holds the schema definition for the Irregular entity.
type Irregular struct {
	ent.Schema
}

// Fields of the Irregular.
func (Irregular) Fields() []ent.Field {
	return []ent.Field{
		field.String("view_url").Unique(),
		field.String("download_url").Unique(),
		field.String("file_name").NotEmpty(),
		field.Int("file_size").Positive(),
	}
}

// Edges of the Irregular.
func (Irregular) Edges() []ent.Edge {
	return nil
}
