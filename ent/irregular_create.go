// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eiri/konyanko/ent/irregular"
)

// IrregularCreate is the builder for creating a Irregular entity.
type IrregularCreate struct {
	config
	mutation *IrregularMutation
	hooks    []Hook
}

// SetViewURL sets the "view_url" field.
func (ic *IrregularCreate) SetViewURL(s string) *IrregularCreate {
	ic.mutation.SetViewURL(s)
	return ic
}

// SetDownloadURL sets the "download_url" field.
func (ic *IrregularCreate) SetDownloadURL(s string) *IrregularCreate {
	ic.mutation.SetDownloadURL(s)
	return ic
}

// SetFileName sets the "file_name" field.
func (ic *IrregularCreate) SetFileName(s string) *IrregularCreate {
	ic.mutation.SetFileName(s)
	return ic
}

// SetFileSize sets the "file_size" field.
func (ic *IrregularCreate) SetFileSize(i int) *IrregularCreate {
	ic.mutation.SetFileSize(i)
	return ic
}

// Mutation returns the IrregularMutation object of the builder.
func (ic *IrregularCreate) Mutation() *IrregularMutation {
	return ic.mutation
}

// Save creates the Irregular in the database.
func (ic *IrregularCreate) Save(ctx context.Context) (*Irregular, error) {
	return withHooks(ctx, ic.sqlSave, ic.mutation, ic.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *IrregularCreate) SaveX(ctx context.Context) *Irregular {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *IrregularCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *IrregularCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *IrregularCreate) check() error {
	if _, ok := ic.mutation.ViewURL(); !ok {
		return &ValidationError{Name: "view_url", err: errors.New(`ent: missing required field "Irregular.view_url"`)}
	}
	if _, ok := ic.mutation.DownloadURL(); !ok {
		return &ValidationError{Name: "download_url", err: errors.New(`ent: missing required field "Irregular.download_url"`)}
	}
	if _, ok := ic.mutation.FileName(); !ok {
		return &ValidationError{Name: "file_name", err: errors.New(`ent: missing required field "Irregular.file_name"`)}
	}
	if v, ok := ic.mutation.FileName(); ok {
		if err := irregular.FileNameValidator(v); err != nil {
			return &ValidationError{Name: "file_name", err: fmt.Errorf(`ent: validator failed for field "Irregular.file_name": %w`, err)}
		}
	}
	if _, ok := ic.mutation.FileSize(); !ok {
		return &ValidationError{Name: "file_size", err: errors.New(`ent: missing required field "Irregular.file_size"`)}
	}
	if v, ok := ic.mutation.FileSize(); ok {
		if err := irregular.FileSizeValidator(v); err != nil {
			return &ValidationError{Name: "file_size", err: fmt.Errorf(`ent: validator failed for field "Irregular.file_size": %w`, err)}
		}
	}
	return nil
}

func (ic *IrregularCreate) sqlSave(ctx context.Context) (*Irregular, error) {
	if err := ic.check(); err != nil {
		return nil, err
	}
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ic.mutation.id = &_node.ID
	ic.mutation.done = true
	return _node, nil
}

func (ic *IrregularCreate) createSpec() (*Irregular, *sqlgraph.CreateSpec) {
	var (
		_node = &Irregular{config: ic.config}
		_spec = sqlgraph.NewCreateSpec(irregular.Table, sqlgraph.NewFieldSpec(irregular.FieldID, field.TypeInt))
	)
	if value, ok := ic.mutation.ViewURL(); ok {
		_spec.SetField(irregular.FieldViewURL, field.TypeString, value)
		_node.ViewURL = value
	}
	if value, ok := ic.mutation.DownloadURL(); ok {
		_spec.SetField(irregular.FieldDownloadURL, field.TypeString, value)
		_node.DownloadURL = value
	}
	if value, ok := ic.mutation.FileName(); ok {
		_spec.SetField(irregular.FieldFileName, field.TypeString, value)
		_node.FileName = value
	}
	if value, ok := ic.mutation.FileSize(); ok {
		_spec.SetField(irregular.FieldFileSize, field.TypeInt, value)
		_node.FileSize = value
	}
	return _node, _spec
}

// IrregularCreateBulk is the builder for creating many Irregular entities in bulk.
type IrregularCreateBulk struct {
	config
	err      error
	builders []*IrregularCreate
}

// Save creates the Irregular entities in the database.
func (icb *IrregularCreateBulk) Save(ctx context.Context) ([]*Irregular, error) {
	if icb.err != nil {
		return nil, icb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Irregular, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IrregularMutation)
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
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *IrregularCreateBulk) SaveX(ctx context.Context) []*Irregular {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *IrregularCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *IrregularCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}
