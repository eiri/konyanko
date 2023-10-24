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
		field.Int("episode_number").NonNegative().Default(0),
		field.Int("anime_season").NonNegative().Default(1),
		field.String("resolution").Optional().Nillable(),
		field.String("video_codec").Optional().Nillable(),
		field.String("audio_codec").Optional().Nillable(),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Item.Type).Ref("episode").Unique().Required(),
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
