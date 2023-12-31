// Code generated by ent, DO NOT EDIT.

package episode

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eiri/konyanko/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldID, id))
}

// EpisodeNumber applies equality check predicate on the "episode_number" field. It's identical to EpisodeNumberEQ.
func EpisodeNumber(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldEpisodeNumber, v))
}

// AnimeSeason applies equality check predicate on the "anime_season" field. It's identical to AnimeSeasonEQ.
func AnimeSeason(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAnimeSeason, v))
}

// Resolution applies equality check predicate on the "resolution" field. It's identical to ResolutionEQ.
func Resolution(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldResolution, v))
}

// VideoCodec applies equality check predicate on the "video_codec" field. It's identical to VideoCodecEQ.
func VideoCodec(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldVideoCodec, v))
}

// AudioCodec applies equality check predicate on the "audio_codec" field. It's identical to AudioCodecEQ.
func AudioCodec(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAudioCodec, v))
}

// EpisodeNumberEQ applies the EQ predicate on the "episode_number" field.
func EpisodeNumberEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldEpisodeNumber, v))
}

// EpisodeNumberNEQ applies the NEQ predicate on the "episode_number" field.
func EpisodeNumberNEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldEpisodeNumber, v))
}

// EpisodeNumberIn applies the In predicate on the "episode_number" field.
func EpisodeNumberIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldEpisodeNumber, vs...))
}

// EpisodeNumberNotIn applies the NotIn predicate on the "episode_number" field.
func EpisodeNumberNotIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldEpisodeNumber, vs...))
}

// EpisodeNumberGT applies the GT predicate on the "episode_number" field.
func EpisodeNumberGT(v int) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldEpisodeNumber, v))
}

// EpisodeNumberGTE applies the GTE predicate on the "episode_number" field.
func EpisodeNumberGTE(v int) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldEpisodeNumber, v))
}

// EpisodeNumberLT applies the LT predicate on the "episode_number" field.
func EpisodeNumberLT(v int) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldEpisodeNumber, v))
}

// EpisodeNumberLTE applies the LTE predicate on the "episode_number" field.
func EpisodeNumberLTE(v int) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldEpisodeNumber, v))
}

// AnimeSeasonEQ applies the EQ predicate on the "anime_season" field.
func AnimeSeasonEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAnimeSeason, v))
}

// AnimeSeasonNEQ applies the NEQ predicate on the "anime_season" field.
func AnimeSeasonNEQ(v int) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldAnimeSeason, v))
}

// AnimeSeasonIn applies the In predicate on the "anime_season" field.
func AnimeSeasonIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldAnimeSeason, vs...))
}

// AnimeSeasonNotIn applies the NotIn predicate on the "anime_season" field.
func AnimeSeasonNotIn(vs ...int) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldAnimeSeason, vs...))
}

// AnimeSeasonGT applies the GT predicate on the "anime_season" field.
func AnimeSeasonGT(v int) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldAnimeSeason, v))
}

// AnimeSeasonGTE applies the GTE predicate on the "anime_season" field.
func AnimeSeasonGTE(v int) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldAnimeSeason, v))
}

// AnimeSeasonLT applies the LT predicate on the "anime_season" field.
func AnimeSeasonLT(v int) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldAnimeSeason, v))
}

// AnimeSeasonLTE applies the LTE predicate on the "anime_season" field.
func AnimeSeasonLTE(v int) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldAnimeSeason, v))
}

// ResolutionEQ applies the EQ predicate on the "resolution" field.
func ResolutionEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldResolution, v))
}

// ResolutionNEQ applies the NEQ predicate on the "resolution" field.
func ResolutionNEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldResolution, v))
}

// ResolutionIn applies the In predicate on the "resolution" field.
func ResolutionIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldResolution, vs...))
}

// ResolutionNotIn applies the NotIn predicate on the "resolution" field.
func ResolutionNotIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldResolution, vs...))
}

// ResolutionGT applies the GT predicate on the "resolution" field.
func ResolutionGT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldResolution, v))
}

// ResolutionGTE applies the GTE predicate on the "resolution" field.
func ResolutionGTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldResolution, v))
}

// ResolutionLT applies the LT predicate on the "resolution" field.
func ResolutionLT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldResolution, v))
}

// ResolutionLTE applies the LTE predicate on the "resolution" field.
func ResolutionLTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldResolution, v))
}

// ResolutionContains applies the Contains predicate on the "resolution" field.
func ResolutionContains(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContains(FieldResolution, v))
}

// ResolutionHasPrefix applies the HasPrefix predicate on the "resolution" field.
func ResolutionHasPrefix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasPrefix(FieldResolution, v))
}

// ResolutionHasSuffix applies the HasSuffix predicate on the "resolution" field.
func ResolutionHasSuffix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasSuffix(FieldResolution, v))
}

// ResolutionIsNil applies the IsNil predicate on the "resolution" field.
func ResolutionIsNil() predicate.Episode {
	return predicate.Episode(sql.FieldIsNull(FieldResolution))
}

// ResolutionNotNil applies the NotNil predicate on the "resolution" field.
func ResolutionNotNil() predicate.Episode {
	return predicate.Episode(sql.FieldNotNull(FieldResolution))
}

// ResolutionEqualFold applies the EqualFold predicate on the "resolution" field.
func ResolutionEqualFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEqualFold(FieldResolution, v))
}

// ResolutionContainsFold applies the ContainsFold predicate on the "resolution" field.
func ResolutionContainsFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContainsFold(FieldResolution, v))
}

// VideoCodecEQ applies the EQ predicate on the "video_codec" field.
func VideoCodecEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldVideoCodec, v))
}

// VideoCodecNEQ applies the NEQ predicate on the "video_codec" field.
func VideoCodecNEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldVideoCodec, v))
}

// VideoCodecIn applies the In predicate on the "video_codec" field.
func VideoCodecIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldVideoCodec, vs...))
}

// VideoCodecNotIn applies the NotIn predicate on the "video_codec" field.
func VideoCodecNotIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldVideoCodec, vs...))
}

// VideoCodecGT applies the GT predicate on the "video_codec" field.
func VideoCodecGT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldVideoCodec, v))
}

// VideoCodecGTE applies the GTE predicate on the "video_codec" field.
func VideoCodecGTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldVideoCodec, v))
}

// VideoCodecLT applies the LT predicate on the "video_codec" field.
func VideoCodecLT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldVideoCodec, v))
}

// VideoCodecLTE applies the LTE predicate on the "video_codec" field.
func VideoCodecLTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldVideoCodec, v))
}

// VideoCodecContains applies the Contains predicate on the "video_codec" field.
func VideoCodecContains(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContains(FieldVideoCodec, v))
}

// VideoCodecHasPrefix applies the HasPrefix predicate on the "video_codec" field.
func VideoCodecHasPrefix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasPrefix(FieldVideoCodec, v))
}

// VideoCodecHasSuffix applies the HasSuffix predicate on the "video_codec" field.
func VideoCodecHasSuffix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasSuffix(FieldVideoCodec, v))
}

// VideoCodecIsNil applies the IsNil predicate on the "video_codec" field.
func VideoCodecIsNil() predicate.Episode {
	return predicate.Episode(sql.FieldIsNull(FieldVideoCodec))
}

// VideoCodecNotNil applies the NotNil predicate on the "video_codec" field.
func VideoCodecNotNil() predicate.Episode {
	return predicate.Episode(sql.FieldNotNull(FieldVideoCodec))
}

// VideoCodecEqualFold applies the EqualFold predicate on the "video_codec" field.
func VideoCodecEqualFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEqualFold(FieldVideoCodec, v))
}

// VideoCodecContainsFold applies the ContainsFold predicate on the "video_codec" field.
func VideoCodecContainsFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContainsFold(FieldVideoCodec, v))
}

// AudioCodecEQ applies the EQ predicate on the "audio_codec" field.
func AudioCodecEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEQ(FieldAudioCodec, v))
}

// AudioCodecNEQ applies the NEQ predicate on the "audio_codec" field.
func AudioCodecNEQ(v string) predicate.Episode {
	return predicate.Episode(sql.FieldNEQ(FieldAudioCodec, v))
}

// AudioCodecIn applies the In predicate on the "audio_codec" field.
func AudioCodecIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldIn(FieldAudioCodec, vs...))
}

// AudioCodecNotIn applies the NotIn predicate on the "audio_codec" field.
func AudioCodecNotIn(vs ...string) predicate.Episode {
	return predicate.Episode(sql.FieldNotIn(FieldAudioCodec, vs...))
}

// AudioCodecGT applies the GT predicate on the "audio_codec" field.
func AudioCodecGT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGT(FieldAudioCodec, v))
}

// AudioCodecGTE applies the GTE predicate on the "audio_codec" field.
func AudioCodecGTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldGTE(FieldAudioCodec, v))
}

// AudioCodecLT applies the LT predicate on the "audio_codec" field.
func AudioCodecLT(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLT(FieldAudioCodec, v))
}

// AudioCodecLTE applies the LTE predicate on the "audio_codec" field.
func AudioCodecLTE(v string) predicate.Episode {
	return predicate.Episode(sql.FieldLTE(FieldAudioCodec, v))
}

// AudioCodecContains applies the Contains predicate on the "audio_codec" field.
func AudioCodecContains(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContains(FieldAudioCodec, v))
}

// AudioCodecHasPrefix applies the HasPrefix predicate on the "audio_codec" field.
func AudioCodecHasPrefix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasPrefix(FieldAudioCodec, v))
}

// AudioCodecHasSuffix applies the HasSuffix predicate on the "audio_codec" field.
func AudioCodecHasSuffix(v string) predicate.Episode {
	return predicate.Episode(sql.FieldHasSuffix(FieldAudioCodec, v))
}

// AudioCodecIsNil applies the IsNil predicate on the "audio_codec" field.
func AudioCodecIsNil() predicate.Episode {
	return predicate.Episode(sql.FieldIsNull(FieldAudioCodec))
}

// AudioCodecNotNil applies the NotNil predicate on the "audio_codec" field.
func AudioCodecNotNil() predicate.Episode {
	return predicate.Episode(sql.FieldNotNull(FieldAudioCodec))
}

// AudioCodecEqualFold applies the EqualFold predicate on the "audio_codec" field.
func AudioCodecEqualFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldEqualFold(FieldAudioCodec, v))
}

// AudioCodecContainsFold applies the ContainsFold predicate on the "audio_codec" field.
func AudioCodecContainsFold(v string) predicate.Episode {
	return predicate.Episode(sql.FieldContainsFold(FieldAudioCodec, v))
}

// HasItem applies the HasEdge predicate on the "item" edge.
func HasItem() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, ItemTable, ItemColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemWith applies the HasEdge predicate on the "item" edge with a given conditions (other predicates).
func HasItemWith(preds ...predicate.Item) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := newItemStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnime applies the HasEdge predicate on the "anime" edge.
func HasAnime() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AnimeTable, AnimeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnimeWith applies the HasEdge predicate on the "anime" edge with a given conditions (other predicates).
func HasAnimeWith(preds ...predicate.Anime) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := newAnimeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasReleaseGroup applies the HasEdge predicate on the "release_group" edge.
func HasReleaseGroup() predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ReleaseGroupTable, ReleaseGroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasReleaseGroupWith applies the HasEdge predicate on the "release_group" edge with a given conditions (other predicates).
func HasReleaseGroupWith(preds ...predicate.ReleaseGroup) predicate.Episode {
	return predicate.Episode(func(s *sql.Selector) {
		step := newReleaseGroupStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Episode) predicate.Episode {
	return predicate.Episode(sql.NotPredicates(p))
}
