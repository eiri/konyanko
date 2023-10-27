// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/predicate"
	"github.com/eiri/konyanko/ent/releasegroup"
)

// EpisodeUpdate is the builder for updating Episode entities.
type EpisodeUpdate struct {
	config
	hooks    []Hook
	mutation *EpisodeMutation
}

// Where appends a list predicates to the EpisodeUpdate builder.
func (eu *EpisodeUpdate) Where(ps ...predicate.Episode) *EpisodeUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetEpisodeNumber sets the "episode_number" field.
func (eu *EpisodeUpdate) SetEpisodeNumber(i int) *EpisodeUpdate {
	eu.mutation.ResetEpisodeNumber()
	eu.mutation.SetEpisodeNumber(i)
	return eu
}

// SetNillableEpisodeNumber sets the "episode_number" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableEpisodeNumber(i *int) *EpisodeUpdate {
	if i != nil {
		eu.SetEpisodeNumber(*i)
	}
	return eu
}

// AddEpisodeNumber adds i to the "episode_number" field.
func (eu *EpisodeUpdate) AddEpisodeNumber(i int) *EpisodeUpdate {
	eu.mutation.AddEpisodeNumber(i)
	return eu
}

// SetAnimeSeason sets the "anime_season" field.
func (eu *EpisodeUpdate) SetAnimeSeason(i int) *EpisodeUpdate {
	eu.mutation.ResetAnimeSeason()
	eu.mutation.SetAnimeSeason(i)
	return eu
}

// SetNillableAnimeSeason sets the "anime_season" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableAnimeSeason(i *int) *EpisodeUpdate {
	if i != nil {
		eu.SetAnimeSeason(*i)
	}
	return eu
}

// AddAnimeSeason adds i to the "anime_season" field.
func (eu *EpisodeUpdate) AddAnimeSeason(i int) *EpisodeUpdate {
	eu.mutation.AddAnimeSeason(i)
	return eu
}

// SetResolution sets the "resolution" field.
func (eu *EpisodeUpdate) SetResolution(s string) *EpisodeUpdate {
	eu.mutation.SetResolution(s)
	return eu
}

// SetNillableResolution sets the "resolution" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableResolution(s *string) *EpisodeUpdate {
	if s != nil {
		eu.SetResolution(*s)
	}
	return eu
}

// ClearResolution clears the value of the "resolution" field.
func (eu *EpisodeUpdate) ClearResolution() *EpisodeUpdate {
	eu.mutation.ClearResolution()
	return eu
}

// SetVideoCodec sets the "video_codec" field.
func (eu *EpisodeUpdate) SetVideoCodec(s string) *EpisodeUpdate {
	eu.mutation.SetVideoCodec(s)
	return eu
}

// SetNillableVideoCodec sets the "video_codec" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableVideoCodec(s *string) *EpisodeUpdate {
	if s != nil {
		eu.SetVideoCodec(*s)
	}
	return eu
}

// ClearVideoCodec clears the value of the "video_codec" field.
func (eu *EpisodeUpdate) ClearVideoCodec() *EpisodeUpdate {
	eu.mutation.ClearVideoCodec()
	return eu
}

// SetAudioCodec sets the "audio_codec" field.
func (eu *EpisodeUpdate) SetAudioCodec(s string) *EpisodeUpdate {
	eu.mutation.SetAudioCodec(s)
	return eu
}

// SetNillableAudioCodec sets the "audio_codec" field if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableAudioCodec(s *string) *EpisodeUpdate {
	if s != nil {
		eu.SetAudioCodec(*s)
	}
	return eu
}

// ClearAudioCodec clears the value of the "audio_codec" field.
func (eu *EpisodeUpdate) ClearAudioCodec() *EpisodeUpdate {
	eu.mutation.ClearAudioCodec()
	return eu
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (eu *EpisodeUpdate) SetItemID(id int) *EpisodeUpdate {
	eu.mutation.SetItemID(id)
	return eu
}

// SetItem sets the "item" edge to the Item entity.
func (eu *EpisodeUpdate) SetItem(i *Item) *EpisodeUpdate {
	return eu.SetItemID(i.ID)
}

// SetAnimeID sets the "anime" edge to the Anime entity by ID.
func (eu *EpisodeUpdate) SetAnimeID(id int) *EpisodeUpdate {
	eu.mutation.SetAnimeID(id)
	return eu
}

// SetAnime sets the "anime" edge to the Anime entity.
func (eu *EpisodeUpdate) SetAnime(a *Anime) *EpisodeUpdate {
	return eu.SetAnimeID(a.ID)
}

// SetReleaseGroupID sets the "release_group" edge to the ReleaseGroup entity by ID.
func (eu *EpisodeUpdate) SetReleaseGroupID(id int) *EpisodeUpdate {
	eu.mutation.SetReleaseGroupID(id)
	return eu
}

// SetNillableReleaseGroupID sets the "release_group" edge to the ReleaseGroup entity by ID if the given value is not nil.
func (eu *EpisodeUpdate) SetNillableReleaseGroupID(id *int) *EpisodeUpdate {
	if id != nil {
		eu = eu.SetReleaseGroupID(*id)
	}
	return eu
}

// SetReleaseGroup sets the "release_group" edge to the ReleaseGroup entity.
func (eu *EpisodeUpdate) SetReleaseGroup(r *ReleaseGroup) *EpisodeUpdate {
	return eu.SetReleaseGroupID(r.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (eu *EpisodeUpdate) Mutation() *EpisodeMutation {
	return eu.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (eu *EpisodeUpdate) ClearItem() *EpisodeUpdate {
	eu.mutation.ClearItem()
	return eu
}

// ClearAnime clears the "anime" edge to the Anime entity.
func (eu *EpisodeUpdate) ClearAnime() *EpisodeUpdate {
	eu.mutation.ClearAnime()
	return eu
}

// ClearReleaseGroup clears the "release_group" edge to the ReleaseGroup entity.
func (eu *EpisodeUpdate) ClearReleaseGroup() *EpisodeUpdate {
	eu.mutation.ClearReleaseGroup()
	return eu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EpisodeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, eu.sqlSave, eu.mutation, eu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EpisodeUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EpisodeUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EpisodeUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (eu *EpisodeUpdate) check() error {
	if v, ok := eu.mutation.EpisodeNumber(); ok {
		if err := episode.EpisodeNumberValidator(v); err != nil {
			return &ValidationError{Name: "episode_number", err: fmt.Errorf(`ent: validator failed for field "Episode.episode_number": %w`, err)}
		}
	}
	if v, ok := eu.mutation.AnimeSeason(); ok {
		if err := episode.AnimeSeasonValidator(v); err != nil {
			return &ValidationError{Name: "anime_season", err: fmt.Errorf(`ent: validator failed for field "Episode.anime_season": %w`, err)}
		}
	}
	if _, ok := eu.mutation.ItemID(); eu.mutation.ItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Episode.item"`)
	}
	if _, ok := eu.mutation.AnimeID(); eu.mutation.AnimeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Episode.anime"`)
	}
	return nil
}

func (eu *EpisodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := eu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(episode.Table, episode.Columns, sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt))
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.EpisodeNumber(); ok {
		_spec.SetField(episode.FieldEpisodeNumber, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedEpisodeNumber(); ok {
		_spec.AddField(episode.FieldEpisodeNumber, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AnimeSeason(); ok {
		_spec.SetField(episode.FieldAnimeSeason, field.TypeInt, value)
	}
	if value, ok := eu.mutation.AddedAnimeSeason(); ok {
		_spec.AddField(episode.FieldAnimeSeason, field.TypeInt, value)
	}
	if value, ok := eu.mutation.Resolution(); ok {
		_spec.SetField(episode.FieldResolution, field.TypeString, value)
	}
	if eu.mutation.ResolutionCleared() {
		_spec.ClearField(episode.FieldResolution, field.TypeString)
	}
	if value, ok := eu.mutation.VideoCodec(); ok {
		_spec.SetField(episode.FieldVideoCodec, field.TypeString, value)
	}
	if eu.mutation.VideoCodecCleared() {
		_spec.ClearField(episode.FieldVideoCodec, field.TypeString)
	}
	if value, ok := eu.mutation.AudioCodec(); ok {
		_spec.SetField(episode.FieldAudioCodec, field.TypeString, value)
	}
	if eu.mutation.AudioCodecCleared() {
		_spec.ClearField(episode.FieldAudioCodec, field.TypeString)
	}
	if eu.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   episode.ItemTable,
			Columns: []string{episode.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   episode.ItemTable,
			Columns: []string{episode.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.AnimeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.AnimeTable,
			Columns: []string{episode.AnimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.AnimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.AnimeTable,
			Columns: []string{episode.AnimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if eu.mutation.ReleaseGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ReleaseGroupTable,
			Columns: []string{episode.ReleaseGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(releasegroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := eu.mutation.ReleaseGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ReleaseGroupTable,
			Columns: []string{episode.ReleaseGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(releasegroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	eu.mutation.done = true
	return n, nil
}

// EpisodeUpdateOne is the builder for updating a single Episode entity.
type EpisodeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EpisodeMutation
}

// SetEpisodeNumber sets the "episode_number" field.
func (euo *EpisodeUpdateOne) SetEpisodeNumber(i int) *EpisodeUpdateOne {
	euo.mutation.ResetEpisodeNumber()
	euo.mutation.SetEpisodeNumber(i)
	return euo
}

// SetNillableEpisodeNumber sets the "episode_number" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableEpisodeNumber(i *int) *EpisodeUpdateOne {
	if i != nil {
		euo.SetEpisodeNumber(*i)
	}
	return euo
}

// AddEpisodeNumber adds i to the "episode_number" field.
func (euo *EpisodeUpdateOne) AddEpisodeNumber(i int) *EpisodeUpdateOne {
	euo.mutation.AddEpisodeNumber(i)
	return euo
}

// SetAnimeSeason sets the "anime_season" field.
func (euo *EpisodeUpdateOne) SetAnimeSeason(i int) *EpisodeUpdateOne {
	euo.mutation.ResetAnimeSeason()
	euo.mutation.SetAnimeSeason(i)
	return euo
}

// SetNillableAnimeSeason sets the "anime_season" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableAnimeSeason(i *int) *EpisodeUpdateOne {
	if i != nil {
		euo.SetAnimeSeason(*i)
	}
	return euo
}

// AddAnimeSeason adds i to the "anime_season" field.
func (euo *EpisodeUpdateOne) AddAnimeSeason(i int) *EpisodeUpdateOne {
	euo.mutation.AddAnimeSeason(i)
	return euo
}

// SetResolution sets the "resolution" field.
func (euo *EpisodeUpdateOne) SetResolution(s string) *EpisodeUpdateOne {
	euo.mutation.SetResolution(s)
	return euo
}

// SetNillableResolution sets the "resolution" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableResolution(s *string) *EpisodeUpdateOne {
	if s != nil {
		euo.SetResolution(*s)
	}
	return euo
}

// ClearResolution clears the value of the "resolution" field.
func (euo *EpisodeUpdateOne) ClearResolution() *EpisodeUpdateOne {
	euo.mutation.ClearResolution()
	return euo
}

// SetVideoCodec sets the "video_codec" field.
func (euo *EpisodeUpdateOne) SetVideoCodec(s string) *EpisodeUpdateOne {
	euo.mutation.SetVideoCodec(s)
	return euo
}

// SetNillableVideoCodec sets the "video_codec" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableVideoCodec(s *string) *EpisodeUpdateOne {
	if s != nil {
		euo.SetVideoCodec(*s)
	}
	return euo
}

// ClearVideoCodec clears the value of the "video_codec" field.
func (euo *EpisodeUpdateOne) ClearVideoCodec() *EpisodeUpdateOne {
	euo.mutation.ClearVideoCodec()
	return euo
}

// SetAudioCodec sets the "audio_codec" field.
func (euo *EpisodeUpdateOne) SetAudioCodec(s string) *EpisodeUpdateOne {
	euo.mutation.SetAudioCodec(s)
	return euo
}

// SetNillableAudioCodec sets the "audio_codec" field if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableAudioCodec(s *string) *EpisodeUpdateOne {
	if s != nil {
		euo.SetAudioCodec(*s)
	}
	return euo
}

// ClearAudioCodec clears the value of the "audio_codec" field.
func (euo *EpisodeUpdateOne) ClearAudioCodec() *EpisodeUpdateOne {
	euo.mutation.ClearAudioCodec()
	return euo
}

// SetItemID sets the "item" edge to the Item entity by ID.
func (euo *EpisodeUpdateOne) SetItemID(id int) *EpisodeUpdateOne {
	euo.mutation.SetItemID(id)
	return euo
}

// SetItem sets the "item" edge to the Item entity.
func (euo *EpisodeUpdateOne) SetItem(i *Item) *EpisodeUpdateOne {
	return euo.SetItemID(i.ID)
}

// SetAnimeID sets the "anime" edge to the Anime entity by ID.
func (euo *EpisodeUpdateOne) SetAnimeID(id int) *EpisodeUpdateOne {
	euo.mutation.SetAnimeID(id)
	return euo
}

// SetAnime sets the "anime" edge to the Anime entity.
func (euo *EpisodeUpdateOne) SetAnime(a *Anime) *EpisodeUpdateOne {
	return euo.SetAnimeID(a.ID)
}

// SetReleaseGroupID sets the "release_group" edge to the ReleaseGroup entity by ID.
func (euo *EpisodeUpdateOne) SetReleaseGroupID(id int) *EpisodeUpdateOne {
	euo.mutation.SetReleaseGroupID(id)
	return euo
}

// SetNillableReleaseGroupID sets the "release_group" edge to the ReleaseGroup entity by ID if the given value is not nil.
func (euo *EpisodeUpdateOne) SetNillableReleaseGroupID(id *int) *EpisodeUpdateOne {
	if id != nil {
		euo = euo.SetReleaseGroupID(*id)
	}
	return euo
}

// SetReleaseGroup sets the "release_group" edge to the ReleaseGroup entity.
func (euo *EpisodeUpdateOne) SetReleaseGroup(r *ReleaseGroup) *EpisodeUpdateOne {
	return euo.SetReleaseGroupID(r.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (euo *EpisodeUpdateOne) Mutation() *EpisodeMutation {
	return euo.mutation
}

// ClearItem clears the "item" edge to the Item entity.
func (euo *EpisodeUpdateOne) ClearItem() *EpisodeUpdateOne {
	euo.mutation.ClearItem()
	return euo
}

// ClearAnime clears the "anime" edge to the Anime entity.
func (euo *EpisodeUpdateOne) ClearAnime() *EpisodeUpdateOne {
	euo.mutation.ClearAnime()
	return euo
}

// ClearReleaseGroup clears the "release_group" edge to the ReleaseGroup entity.
func (euo *EpisodeUpdateOne) ClearReleaseGroup() *EpisodeUpdateOne {
	euo.mutation.ClearReleaseGroup()
	return euo
}

// Where appends a list predicates to the EpisodeUpdate builder.
func (euo *EpisodeUpdateOne) Where(ps ...predicate.Episode) *EpisodeUpdateOne {
	euo.mutation.Where(ps...)
	return euo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EpisodeUpdateOne) Select(field string, fields ...string) *EpisodeUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Episode entity.
func (euo *EpisodeUpdateOne) Save(ctx context.Context) (*Episode, error) {
	return withHooks(ctx, euo.sqlSave, euo.mutation, euo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EpisodeUpdateOne) SaveX(ctx context.Context) *Episode {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EpisodeUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EpisodeUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (euo *EpisodeUpdateOne) check() error {
	if v, ok := euo.mutation.EpisodeNumber(); ok {
		if err := episode.EpisodeNumberValidator(v); err != nil {
			return &ValidationError{Name: "episode_number", err: fmt.Errorf(`ent: validator failed for field "Episode.episode_number": %w`, err)}
		}
	}
	if v, ok := euo.mutation.AnimeSeason(); ok {
		if err := episode.AnimeSeasonValidator(v); err != nil {
			return &ValidationError{Name: "anime_season", err: fmt.Errorf(`ent: validator failed for field "Episode.anime_season": %w`, err)}
		}
	}
	if _, ok := euo.mutation.ItemID(); euo.mutation.ItemCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Episode.item"`)
	}
	if _, ok := euo.mutation.AnimeID(); euo.mutation.AnimeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Episode.anime"`)
	}
	return nil
}

func (euo *EpisodeUpdateOne) sqlSave(ctx context.Context) (_node *Episode, err error) {
	if err := euo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(episode.Table, episode.Columns, sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt))
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Episode.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, episode.FieldID)
		for _, f := range fields {
			if !episode.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != episode.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.EpisodeNumber(); ok {
		_spec.SetField(episode.FieldEpisodeNumber, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedEpisodeNumber(); ok {
		_spec.AddField(episode.FieldEpisodeNumber, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AnimeSeason(); ok {
		_spec.SetField(episode.FieldAnimeSeason, field.TypeInt, value)
	}
	if value, ok := euo.mutation.AddedAnimeSeason(); ok {
		_spec.AddField(episode.FieldAnimeSeason, field.TypeInt, value)
	}
	if value, ok := euo.mutation.Resolution(); ok {
		_spec.SetField(episode.FieldResolution, field.TypeString, value)
	}
	if euo.mutation.ResolutionCleared() {
		_spec.ClearField(episode.FieldResolution, field.TypeString)
	}
	if value, ok := euo.mutation.VideoCodec(); ok {
		_spec.SetField(episode.FieldVideoCodec, field.TypeString, value)
	}
	if euo.mutation.VideoCodecCleared() {
		_spec.ClearField(episode.FieldVideoCodec, field.TypeString)
	}
	if value, ok := euo.mutation.AudioCodec(); ok {
		_spec.SetField(episode.FieldAudioCodec, field.TypeString, value)
	}
	if euo.mutation.AudioCodecCleared() {
		_spec.ClearField(episode.FieldAudioCodec, field.TypeString)
	}
	if euo.mutation.ItemCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   episode.ItemTable,
			Columns: []string{episode.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ItemIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   episode.ItemTable,
			Columns: []string{episode.ItemColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.AnimeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.AnimeTable,
			Columns: []string{episode.AnimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.AnimeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.AnimeTable,
			Columns: []string{episode.AnimeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if euo.mutation.ReleaseGroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ReleaseGroupTable,
			Columns: []string{episode.ReleaseGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(releasegroup.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := euo.mutation.ReleaseGroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.ReleaseGroupTable,
			Columns: []string{episode.ReleaseGroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(releasegroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Episode{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{episode.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	euo.mutation.done = true
	return _node, nil
}
