// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/predicate"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetViewURL sets the "view_url" field.
func (iu *ItemUpdate) SetViewURL(s string) *ItemUpdate {
	iu.mutation.SetViewURL(s)
	return iu
}

// SetDownloadURL sets the "download_url" field.
func (iu *ItemUpdate) SetDownloadURL(s string) *ItemUpdate {
	iu.mutation.SetDownloadURL(s)
	return iu
}

// SetFileName sets the "file_name" field.
func (iu *ItemUpdate) SetFileName(s string) *ItemUpdate {
	iu.mutation.SetFileName(s)
	return iu
}

// SetFileSize sets the "file_size" field.
func (iu *ItemUpdate) SetFileSize(i int) *ItemUpdate {
	iu.mutation.ResetFileSize()
	iu.mutation.SetFileSize(i)
	return iu
}

// AddFileSize adds i to the "file_size" field.
func (iu *ItemUpdate) AddFileSize(i int) *ItemUpdate {
	iu.mutation.AddFileSize(i)
	return iu
}

// SetPublishDate sets the "publish_date" field.
func (iu *ItemUpdate) SetPublishDate(t time.Time) *ItemUpdate {
	iu.mutation.SetPublishDate(t)
	return iu
}

// SetNillablePublishDate sets the "publish_date" field if the given value is not nil.
func (iu *ItemUpdate) SetNillablePublishDate(t *time.Time) *ItemUpdate {
	if t != nil {
		iu.SetPublishDate(*t)
	}
	return iu
}

// SetEpisodesID sets the "episodes" edge to the Episode entity by ID.
func (iu *ItemUpdate) SetEpisodesID(id int) *ItemUpdate {
	iu.mutation.SetEpisodesID(id)
	return iu
}

// SetNillableEpisodesID sets the "episodes" edge to the Episode entity by ID if the given value is not nil.
func (iu *ItemUpdate) SetNillableEpisodesID(id *int) *ItemUpdate {
	if id != nil {
		iu = iu.SetEpisodesID(*id)
	}
	return iu
}

// SetEpisodes sets the "episodes" edge to the Episode entity.
func (iu *ItemUpdate) SetEpisodes(e *Episode) *ItemUpdate {
	return iu.SetEpisodesID(e.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearEpisodes clears the "episodes" edge to the Episode entity.
func (iu *ItemUpdate) ClearEpisodes() *ItemUpdate {
	iu.mutation.ClearEpisodes()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, iu.sqlSave, iu.mutation, iu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iu *ItemUpdate) check() error {
	if v, ok := iu.mutation.FileName(); ok {
		if err := item.FileNameValidator(v); err != nil {
			return &ValidationError{Name: "file_name", err: fmt.Errorf(`ent: validator failed for field "Item.file_name": %w`, err)}
		}
	}
	if v, ok := iu.mutation.FileSize(); ok {
		if err := item.FileSizeValidator(v); err != nil {
			return &ValidationError{Name: "file_size", err: fmt.Errorf(`ent: validator failed for field "Item.file_size": %w`, err)}
		}
	}
	return nil
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.ViewURL(); ok {
		_spec.SetField(item.FieldViewURL, field.TypeString, value)
	}
	if value, ok := iu.mutation.DownloadURL(); ok {
		_spec.SetField(item.FieldDownloadURL, field.TypeString, value)
	}
	if value, ok := iu.mutation.FileName(); ok {
		_spec.SetField(item.FieldFileName, field.TypeString, value)
	}
	if value, ok := iu.mutation.FileSize(); ok {
		_spec.SetField(item.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := iu.mutation.AddedFileSize(); ok {
		_spec.AddField(item.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := iu.mutation.PublishDate(); ok {
		_spec.SetField(item.FieldPublishDate, field.TypeTime, value)
	}
	if iu.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   item.EpisodesTable,
			Columns: []string{item.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   item.EpisodesTable,
			Columns: []string{item.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iu.mutation.done = true
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetViewURL sets the "view_url" field.
func (iuo *ItemUpdateOne) SetViewURL(s string) *ItemUpdateOne {
	iuo.mutation.SetViewURL(s)
	return iuo
}

// SetDownloadURL sets the "download_url" field.
func (iuo *ItemUpdateOne) SetDownloadURL(s string) *ItemUpdateOne {
	iuo.mutation.SetDownloadURL(s)
	return iuo
}

// SetFileName sets the "file_name" field.
func (iuo *ItemUpdateOne) SetFileName(s string) *ItemUpdateOne {
	iuo.mutation.SetFileName(s)
	return iuo
}

// SetFileSize sets the "file_size" field.
func (iuo *ItemUpdateOne) SetFileSize(i int) *ItemUpdateOne {
	iuo.mutation.ResetFileSize()
	iuo.mutation.SetFileSize(i)
	return iuo
}

// AddFileSize adds i to the "file_size" field.
func (iuo *ItemUpdateOne) AddFileSize(i int) *ItemUpdateOne {
	iuo.mutation.AddFileSize(i)
	return iuo
}

// SetPublishDate sets the "publish_date" field.
func (iuo *ItemUpdateOne) SetPublishDate(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetPublishDate(t)
	return iuo
}

// SetNillablePublishDate sets the "publish_date" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillablePublishDate(t *time.Time) *ItemUpdateOne {
	if t != nil {
		iuo.SetPublishDate(*t)
	}
	return iuo
}

// SetEpisodesID sets the "episodes" edge to the Episode entity by ID.
func (iuo *ItemUpdateOne) SetEpisodesID(id int) *ItemUpdateOne {
	iuo.mutation.SetEpisodesID(id)
	return iuo
}

// SetNillableEpisodesID sets the "episodes" edge to the Episode entity by ID if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableEpisodesID(id *int) *ItemUpdateOne {
	if id != nil {
		iuo = iuo.SetEpisodesID(*id)
	}
	return iuo
}

// SetEpisodes sets the "episodes" edge to the Episode entity.
func (iuo *ItemUpdateOne) SetEpisodes(e *Episode) *ItemUpdateOne {
	return iuo.SetEpisodesID(e.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearEpisodes clears the "episodes" edge to the Episode entity.
func (iuo *ItemUpdateOne) ClearEpisodes() *ItemUpdateOne {
	iuo.mutation.ClearEpisodes()
	return iuo
}

// Where appends a list predicates to the ItemUpdate builder.
func (iuo *ItemUpdateOne) Where(ps ...predicate.Item) *ItemUpdateOne {
	iuo.mutation.Where(ps...)
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	return withHooks(ctx, iuo.sqlSave, iuo.mutation, iuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iuo *ItemUpdateOne) check() error {
	if v, ok := iuo.mutation.FileName(); ok {
		if err := item.FileNameValidator(v); err != nil {
			return &ValidationError{Name: "file_name", err: fmt.Errorf(`ent: validator failed for field "Item.file_name": %w`, err)}
		}
	}
	if v, ok := iuo.mutation.FileSize(); ok {
		if err := item.FileSizeValidator(v); err != nil {
			return &ValidationError{Name: "file_size", err: fmt.Errorf(`ent: validator failed for field "Item.file_size": %w`, err)}
		}
	}
	return nil
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	if err := iuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(item.Table, item.Columns, sqlgraph.NewFieldSpec(item.FieldID, field.TypeInt))
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.ViewURL(); ok {
		_spec.SetField(item.FieldViewURL, field.TypeString, value)
	}
	if value, ok := iuo.mutation.DownloadURL(); ok {
		_spec.SetField(item.FieldDownloadURL, field.TypeString, value)
	}
	if value, ok := iuo.mutation.FileName(); ok {
		_spec.SetField(item.FieldFileName, field.TypeString, value)
	}
	if value, ok := iuo.mutation.FileSize(); ok {
		_spec.SetField(item.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.AddedFileSize(); ok {
		_spec.AddField(item.FieldFileSize, field.TypeInt, value)
	}
	if value, ok := iuo.mutation.PublishDate(); ok {
		_spec.SetField(item.FieldPublishDate, field.TypeTime, value)
	}
	if iuo.mutation.EpisodesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   item.EpisodesTable,
			Columns: []string{item.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.EpisodesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   item.EpisodesTable,
			Columns: []string{item.EpisodesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(episode.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iuo.mutation.done = true
	return _node, nil
}