package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// ItemMixin implements the ent.Mixin for sharing common feed item's attributes
type ItemMixin struct {
	mixin.Schema
}

func (ItemMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("view_url").Unique(),
		field.String("download_url").Unique(),
		field.String("file_name").NotEmpty(),
		field.Int("file_size").Positive(),
		field.Time("publish_date").Default(time.Now),
	}
}
