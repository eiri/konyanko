package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/index"
)

// Irregular holds the schema definition for the Irregular entity.
type Irregular struct {
	ent.Schema
}

// Mixin of the Irregular.
func (Irregular) Mixin() []ent.Mixin {
	return []ent.Mixin{
		ItemMixin{},
	}
}

// Fields of the Irregular.
func (Irregular) Fields() []ent.Field {
	return []ent.Field{}
}

// Edges of the Irregular.
func (Irregular) Edges() []ent.Edge {
	return nil
}

// Indices of the Irregular.
func (Irregular) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("view_url"),
	}
}
