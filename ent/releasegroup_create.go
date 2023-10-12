// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/releasegroup"
)

// ReleaseGroupCreate is the builder for creating a ReleaseGroup entity.
type ReleaseGroupCreate struct {
	config
	mutation *ReleaseGroupMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (rgc *ReleaseGroupCreate) SetName(s string) *ReleaseGroupCreate {
	rgc.mutation.SetName(s)
	return rgc
}

// AddEpisodeIDs adds the "episodes" edge to the Episode entity by IDs.
func (rgc *ReleaseGroupCreate) AddEpisodeIDs(ids ...int) *ReleaseGroupCreate {
	rgc.mutation.AddEpisodeIDs(ids...)
	return rgc
}

// AddEpisodes adds the "episodes" edges to the Episode entity.
func (rgc *ReleaseGroupCreate) AddEpisodes(e ...*Episode) *ReleaseGroupCreate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return rgc.AddEpisodeIDs(ids...)
}

// Mutation returns the ReleaseGroupMutation object of the builder.
func (rgc *ReleaseGroupCreate) Mutation() *ReleaseGroupMutation {
	return rgc.mutation
}

// Save creates the ReleaseGroup in the database.
func (rgc *ReleaseGroupCreate) Save(ctx context.Context) (*ReleaseGroup, error) {
	return withHooks(ctx, rgc.sqlSave, rgc.mutation, rgc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rgc *ReleaseGroupCreate) SaveX(ctx context.Context) *ReleaseGroup {
	v, err := rgc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgc *ReleaseGroupCreate) Exec(ctx context.Context) error {
	_, err := rgc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgc *ReleaseGroupCreate) ExecX(ctx context.Context) {
	if err := rgc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rgc *ReleaseGroupCreate) check() error {
	if _, ok := rgc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "ReleaseGroup.name"`)}
	}
	if v, ok := rgc.mutation.Name(); ok {
		if err := releasegroup.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "ReleaseGroup.name": %w`, err)}
		}
	}
	return nil
}

func (rgc *ReleaseGroupCreate) sqlSave(ctx context.Context) (*ReleaseGroup, error) {
	if err := rgc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rgc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rgc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rgc.mutation.id = &_node.ID
	rgc.mutation.done = true
	return _node, nil
}

func (rgc *ReleaseGroupCreate) createSpec() (*ReleaseGroup, *sqlgraph.CreateSpec) {
	var (
		_node = &ReleaseGroup{config: rgc.config}
		_spec = sqlgraph.NewCreateSpec(releasegroup.Table, sqlgraph.NewFieldSpec(releasegroup.FieldID, field.TypeInt))
	)
	if value, ok := rgc.mutation.Name(); ok {
		_spec.SetField(releasegroup.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := rgc.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   releasegroup.EpisodesTable,
			Columns: []string{releasegroup.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReleaseGroupCreateBulk is the builder for creating many ReleaseGroup entities in bulk.
type ReleaseGroupCreateBulk struct {
	config
	err      error
	builders []*ReleaseGroupCreate
}

// Save creates the ReleaseGroup entities in the database.
func (rgcb *ReleaseGroupCreateBulk) Save(ctx context.Context) ([]*ReleaseGroup, error) {
	if rgcb.err != nil {
		return nil, rgcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rgcb.builders))
	nodes := make([]*ReleaseGroup, len(rgcb.builders))
	mutators := make([]Mutator, len(rgcb.builders))
	for i := range rgcb.builders {
		func(i int, root context.Context) {
			builder := rgcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReleaseGroupMutation)
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
					_, err = mutators[i+1].Mutate(root, rgcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rgcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rgcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rgcb *ReleaseGroupCreateBulk) SaveX(ctx context.Context) []*ReleaseGroup {
	v, err := rgcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rgcb *ReleaseGroupCreateBulk) Exec(ctx context.Context) error {
	_, err := rgcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rgcb *ReleaseGroupCreateBulk) ExecX(ctx context.Context) {
	if err := rgcb.Exec(ctx); err != nil {
		panic(err)
	}
}
