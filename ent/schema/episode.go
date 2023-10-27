package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
		field.Int("episode_number").
			NonNegative().
			Default(0).
			Annotations(
				entgql.OrderField("EPISODE_NUMBER"),
			),
		field.Int("anime_season").
			NonNegative().
			Default(1).
			Annotations(
				entgql.OrderField("ANIME_SEASON"),
			),
		field.String("resolution").
			Optional().
			Nillable().
			Annotations(
				entgql.OrderField("RESOLUTION"),
			),
		field.String("video_codec").Optional().Nillable(),
		field.String("audio_codec").Optional().Nillable(),
	}
}

// Edges of the Episode.
func (Episode) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("item", Item.Type).
			Ref("episode").
			Unique().
			Required().
			Annotations(
				entgql.OrderField("ITEM_PUBLISH_DATE"),
			),
		edge.From("anime", Anime.Type).
			Ref("episodes").
			Unique().
			Required().
			Annotations(
				entgql.OrderField("ANIME_TITLE"),
			),
		edge.From("release_group", ReleaseGroup.Type).
			Ref("episodes").
			Unique(),
	}
}

// Indices of the Episode.
func (Episode) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("resolution"),
	}
}

// Annotations of the Episode.
func (Episode) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.MultiOrder(),
	}
}
