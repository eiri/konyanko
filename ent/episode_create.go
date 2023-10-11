// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/releasegroup"
)

// EpisodeCreate is the builder for creating a Episode entity.
type EpisodeCreate struct {
	config
	mutation *EpisodeMutation
	hooks    []Hook
}

// SetEpisodeNumber sets the "episode_number" field.
func (ec *EpisodeCreate) SetEpisodeNumber(i int) *EpisodeCreate {
	ec.mutation.SetEpisodeNumber(i)
	return ec
}

// SetNillableEpisodeNumber sets the "episode_number" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableEpisodeNumber(i *int) *EpisodeCreate {
	if i != nil {
		ec.SetEpisodeNumber(*i)
	}
	return ec
}

// SetAnimeSeason sets the "anime_season" field.
func (ec *EpisodeCreate) SetAnimeSeason(i int) *EpisodeCreate {
	ec.mutation.SetAnimeSeason(i)
	return ec
}

// SetNillableAnimeSeason sets the "anime_season" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableAnimeSeason(i *int) *EpisodeCreate {
	if i != nil {
		ec.SetAnimeSeason(*i)
	}
	return ec
}

// SetViewURL sets the "view_url" field.
func (ec *EpisodeCreate) SetViewURL(s string) *EpisodeCreate {
	ec.mutation.SetViewURL(s)
	return ec
}

// SetDownloadURL sets the "download_url" field.
func (ec *EpisodeCreate) SetDownloadURL(s string) *EpisodeCreate {
	ec.mutation.SetDownloadURL(s)
	return ec
}

// SetFileName sets the "file_name" field.
func (ec *EpisodeCreate) SetFileName(s string) *EpisodeCreate {
	ec.mutation.SetFileName(s)
	return ec
}

// SetFileSize sets the "file_size" field.
func (ec *EpisodeCreate) SetFileSize(i int) *EpisodeCreate {
	ec.mutation.SetFileSize(i)
	return ec
}

// SetResolution sets the "resolution" field.
func (ec *EpisodeCreate) SetResolution(s string) *EpisodeCreate {
	ec.mutation.SetResolution(s)
	return ec
}

// SetNillableResolution sets the "resolution" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableResolution(s *string) *EpisodeCreate {
	if s != nil {
		ec.SetResolution(*s)
	}
	return ec
}

// SetVideoCodec sets the "video_codec" field.
func (ec *EpisodeCreate) SetVideoCodec(s string) *EpisodeCreate {
	ec.mutation.SetVideoCodec(s)
	return ec
}

// SetNillableVideoCodec sets the "video_codec" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableVideoCodec(s *string) *EpisodeCreate {
	if s != nil {
		ec.SetVideoCodec(*s)
	}
	return ec
}

// SetAudioCodec sets the "audio_codec" field.
func (ec *EpisodeCreate) SetAudioCodec(s string) *EpisodeCreate {
	ec.mutation.SetAudioCodec(s)
	return ec
}

// SetNillableAudioCodec sets the "audio_codec" field if the given value is not nil.
func (ec *EpisodeCreate) SetNillableAudioCodec(s *string) *EpisodeCreate {
	if s != nil {
		ec.SetAudioCodec(*s)
	}
	return ec
}

// SetTitleID sets the "title" edge to the Anime entity by ID.
func (ec *EpisodeCreate) SetTitleID(id int) *EpisodeCreate {
	ec.mutation.SetTitleID(id)
	return ec
}

// SetTitle sets the "title" edge to the Anime entity.
func (ec *EpisodeCreate) SetTitle(a *Anime) *EpisodeCreate {
	return ec.SetTitleID(a.ID)
}

// SetReleaseGroupID sets the "release_group" edge to the ReleaseGroup entity by ID.
func (ec *EpisodeCreate) SetReleaseGroupID(id int) *EpisodeCreate {
	ec.mutation.SetReleaseGroupID(id)
	return ec
}

// SetNillableReleaseGroupID sets the "release_group" edge to the ReleaseGroup entity by ID if the given value is not nil.
func (ec *EpisodeCreate) SetNillableReleaseGroupID(id *int) *EpisodeCreate {
	if id != nil {
		ec = ec.SetReleaseGroupID(*id)
	}
	return ec
}

// SetReleaseGroup sets the "release_group" edge to the ReleaseGroup entity.
func (ec *EpisodeCreate) SetReleaseGroup(r *ReleaseGroup) *EpisodeCreate {
	return ec.SetReleaseGroupID(r.ID)
}

// Mutation returns the EpisodeMutation object of the builder.
func (ec *EpisodeCreate) Mutation() *EpisodeMutation {
	return ec.mutation
}

// Save creates the Episode in the database.
func (ec *EpisodeCreate) Save(ctx context.Context) (*Episode, error) {
	ec.defaults()
	return withHooks(ctx, ec.sqlSave, ec.mutation, ec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ec *EpisodeCreate) SaveX(ctx context.Context) *Episode {
	v, err := ec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ec *EpisodeCreate) Exec(ctx context.Context) error {
	_, err := ec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ec *EpisodeCreate) ExecX(ctx context.Context) {
	if err := ec.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ec *EpisodeCreate) defaults() {
	if _, ok := ec.mutation.EpisodeNumber(); !ok {
		v := episode.DefaultEpisodeNumber
		ec.mutation.SetEpisodeNumber(v)
	}
	if _, ok := ec.mutation.AnimeSeason(); !ok {
		v := episode.DefaultAnimeSeason
		ec.mutation.SetAnimeSeason(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ec *EpisodeCreate) check() error {
	if _, ok := ec.mutation.EpisodeNumber(); !ok {
		return &ValidationError{Name: "episode_number", err: errors.New(`ent: missing required field "Episode.episode_number"`)}
	}
	if v, ok := ec.mutation.EpisodeNumber(); ok {
		if err := episode.EpisodeNumberValidator(v); err != nil {
			return &ValidationError{Name: "episode_number", err: fmt.Errorf(`ent: validator failed for field "Episode.episode_number": %w`, err)}
		}
	}
	if _, ok := ec.mutation.AnimeSeason(); !ok {
		return &ValidationError{Name: "anime_season", err: errors.New(`ent: missing required field "Episode.anime_season"`)}
	}
	if v, ok := ec.mutation.AnimeSeason(); ok {
		if err := episode.AnimeSeasonValidator(v); err != nil {
			return &ValidationError{Name: "anime_season", err: fmt.Errorf(`ent: validator failed for field "Episode.anime_season": %w`, err)}
		}
	}
	if _, ok := ec.mutation.ViewURL(); !ok {
		return &ValidationError{Name: "view_url", err: errors.New(`ent: missing required field "Episode.view_url"`)}
	}
	if _, ok := ec.mutation.DownloadURL(); !ok {
		return &ValidationError{Name: "download_url", err: errors.New(`ent: missing required field "Episode.download_url"`)}
	}
	if _, ok := ec.mutation.FileName(); !ok {
		return &ValidationError{Name: "file_name", err: errors.New(`ent: missing required field "Episode.file_name"`)}
	}
	if v, ok := ec.mutation.FileName(); ok {
		if err := episode.FileNameValidator(v); err != nil {
			return &ValidationError{Name: "file_name", err: fmt.Errorf(`ent: validator failed for field "Episode.file_name": %w`, err)}
		}
	}
	if _, ok := ec.mutation.FileSize(); !ok {
		return &ValidationError{Name: "file_size", err: errors.New(`ent: missing required field "Episode.file_size"`)}
	}
	if v, ok := ec.mutation.FileSize(); ok {
		if err := episode.FileSizeValidator(v); err != nil {
			return &ValidationError{Name: "file_size", err: fmt.Errorf(`ent: validator failed for field "Episode.file_size": %w`, err)}
		}
	}
	if _, ok := ec.mutation.TitleID(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required edge "Episode.title"`)}
	}
	return nil
}

func (ec *EpisodeCreate) sqlSave(ctx context.Context) (*Episode, error) {
	if err := ec.check(); err != nil {
		return nil, err
	}
	_node, _spec := ec.createSpec()
	if err := sqlgraph.CreateNode(ctx, ec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ec.mutation.id = &_node.ID
	ec.mutation.done = true
	return _node, nil
}

func (ec *EpisodeCreate) createSpec() (*Episode, *sqlgraph.CreateSpec) {
	var (
		_node = &Episode{config: ec.config}
		_spec = sqlgraph.NewCreateSpec(episode.Table, sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt))
	)
	if value, ok := ec.mutation.EpisodeNumber(); ok {
		_spec.SetField(episode.FieldEpisodeNumber, field.TypeInt, value)
		_node.EpisodeNumber = value
	}
	if value, ok := ec.mutation.AnimeSeason(); ok {
		_spec.SetField(episode.FieldAnimeSeason, field.TypeInt, value)
		_node.AnimeSeason = value
	}
	if value, ok := ec.mutation.ViewURL(); ok {
		_spec.SetField(episode.FieldViewURL, field.TypeString, value)
		_node.ViewURL = value
	}
	if value, ok := ec.mutation.DownloadURL(); ok {
		_spec.SetField(episode.FieldDownloadURL, field.TypeString, value)
		_node.DownloadURL = value
	}
	if value, ok := ec.mutation.FileName(); ok {
		_spec.SetField(episode.FieldFileName, field.TypeString, value)
		_node.FileName = value
	}
	if value, ok := ec.mutation.FileSize(); ok {
		_spec.SetField(episode.FieldFileSize, field.TypeInt, value)
		_node.FileSize = value
	}
	if value, ok := ec.mutation.Resolution(); ok {
		_spec.SetField(episode.FieldResolution, field.TypeString, value)
		_node.Resolution = value
	}
	if value, ok := ec.mutation.VideoCodec(); ok {
		_spec.SetField(episode.FieldVideoCodec, field.TypeString, value)
		_node.VideoCodec = value
	}
	if value, ok := ec.mutation.AudioCodec(); ok {
		_spec.SetField(episode.FieldAudioCodec, field.TypeString, value)
		_node.AudioCodec = value
	}
	if nodes := ec.mutation.TitleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   episode.TitleTable,
			Columns: []string{episode.TitleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(anime.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.anime_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ec.mutation.ReleaseGroupIDs(); len(nodes) > 0 {
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
		_node.release_group_id = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// EpisodeCreateBulk is the builder for creating many Episode entities in bulk.
type EpisodeCreateBulk struct {
	config
	err      error
	builders []*EpisodeCreate
}

// Save creates the Episode entities in the database.
func (ecb *EpisodeCreateBulk) Save(ctx context.Context) ([]*Episode, error) {
	if ecb.err != nil {
		return nil, ecb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ecb.builders))
	nodes := make([]*Episode, len(ecb.builders))
	mutators := make([]Mutator, len(ecb.builders))
	for i := range ecb.builders {
		func(i int, root context.Context) {
			builder := ecb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*EpisodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ecb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ecb *EpisodeCreateBulk) SaveX(ctx context.Context) []*Episode {
	v, err := ecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ecb *EpisodeCreateBulk) Exec(ctx context.Context) error {
	_, err := ecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ecb *EpisodeCreateBulk) ExecX(ctx context.Context) {
	if err := ecb.Exec(ctx); err != nil {
		panic(err)
	}
}
