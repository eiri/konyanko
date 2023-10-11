package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Episode holds the schema definition for the Episode entity.
type Episode struct {
	ent.Schema
}

// Fields of the Episode.
func (Episode) Fields() []ent.Field {
	return []ent.Field{
		field.Int("number").NonNegative().Default(0),
		field.String("view_url").Unique(),
		field.String("download_url").Unique(),
		field.String("file_name").NotEmpty(),
		field.Int("file_size").Positive(),
		field.String("resolution").Optional(),
		field.String("video_codec").Optional(),
		field.String("audio_codec").Optional(),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("title", Anime.Type).Ref("episodes").Unique().Required(),
		edge.From("release_group", ReleaseGroup.Type).Ref("episodes").Unique(),
	}
}

// Indices of the Episode.
func (Episode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("resolution"),
	}
}
