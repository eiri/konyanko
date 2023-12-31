// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/predicate"
	"github.com/eiri/konyanko/ent/releasegroup"
)

// ReleaseGroupQuery is the builder for querying ReleaseGroup entities.
type ReleaseGroupQuery struct {
	config
	ctx               *QueryContext
	order             []releasegroup.OrderOption
	inters            []Interceptor
	predicates        []predicate.ReleaseGroup
	withEpisodes      *EpisodeQuery
	modifiers         []func(*sql.Selector)
	loadTotal         []func(context.Context, []*ReleaseGroup) error
	withNamedEpisodes map[string]*EpisodeQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ReleaseGroupQuery builder.
func (rgq *ReleaseGroupQuery) Where(ps ...predicate.ReleaseGroup) *ReleaseGroupQuery {
	rgq.predicates = append(rgq.predicates, ps...)
	return rgq
}

// Limit the number of records to be returned by this query.
func (rgq *ReleaseGroupQuery) Limit(limit int) *ReleaseGroupQuery {
	rgq.ctx.Limit = &limit
	return rgq
}

// Offset to start from.
func (rgq *ReleaseGroupQuery) Offset(offset int) *ReleaseGroupQuery {
	rgq.ctx.Offset = &offset
	return rgq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rgq *ReleaseGroupQuery) Unique(unique bool) *ReleaseGroupQuery {
	rgq.ctx.Unique = &unique
	return rgq
}

// Order specifies how the records should be ordered.
func (rgq *ReleaseGroupQuery) Order(o ...releasegroup.OrderOption) *ReleaseGroupQuery {
	rgq.order = append(rgq.order, o...)
	return rgq
}

// QueryEpisodes chains the current query on the "episodes" edge.
func (rgq *ReleaseGroupQuery) QueryEpisodes() *EpisodeQuery {
	query := (&EpisodeClient{config: rgq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rgq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rgq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(releasegroup.Table, releasegroup.FieldID, selector),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, releasegroup.EpisodesTable, releasegroup.EpisodesColumn),
		)
		fromU = sqlgraph.SetNeighbors(rgq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ReleaseGroup entity from the query.
// Returns a *NotFoundError when no ReleaseGroup was found.
func (rgq *ReleaseGroupQuery) First(ctx context.Context) (*ReleaseGroup, error) {
	nodes, err := rgq.Limit(1).All(setContextOp(ctx, rgq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{releasegroup.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) FirstX(ctx context.Context) *ReleaseGroup {
	node, err := rgq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ReleaseGroup ID from the query.
// Returns a *NotFoundError when no ReleaseGroup ID was found.
func (rgq *ReleaseGroupQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rgq.Limit(1).IDs(setContextOp(ctx, rgq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{releasegroup.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) FirstIDX(ctx context.Context) int {
	id, err := rgq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ReleaseGroup entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ReleaseGroup entity is found.
// Returns a *NotFoundError when no ReleaseGroup entities are found.
func (rgq *ReleaseGroupQuery) Only(ctx context.Context) (*ReleaseGroup, error) {
	nodes, err := rgq.Limit(2).All(setContextOp(ctx, rgq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{releasegroup.Label}
	default:
		return nil, &NotSingularError{releasegroup.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) OnlyX(ctx context.Context) *ReleaseGroup {
	node, err := rgq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ReleaseGroup ID in the query.
// Returns a *NotSingularError when more than one ReleaseGroup ID is found.
// Returns a *NotFoundError when no entities are found.
func (rgq *ReleaseGroupQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = rgq.Limit(2).IDs(setContextOp(ctx, rgq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{releasegroup.Label}
	default:
		err = &NotSingularError{releasegroup.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) OnlyIDX(ctx context.Context) int {
	id, err := rgq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ReleaseGroups.
func (rgq *ReleaseGroupQuery) All(ctx context.Context) ([]*ReleaseGroup, error) {
	ctx = setContextOp(ctx, rgq.ctx, "All")
	if err := rgq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ReleaseGroup, *ReleaseGroupQuery]()
	return withInterceptors[[]*ReleaseGroup](ctx, rgq, qr, rgq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) AllX(ctx context.Context) []*ReleaseGroup {
	nodes, err := rgq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ReleaseGroup IDs.
func (rgq *ReleaseGroupQuery) IDs(ctx context.Context) (ids []int, err error) {
	if rgq.ctx.Unique == nil && rgq.path != nil {
		rgq.Unique(true)
	}
	ctx = setContextOp(ctx, rgq.ctx, "IDs")
	if err = rgq.Select(releasegroup.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) IDsX(ctx context.Context) []int {
	ids, err := rgq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rgq *ReleaseGroupQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rgq.ctx, "Count")
	if err := rgq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rgq, querierCount[*ReleaseGroupQuery](), rgq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) CountX(ctx context.Context) int {
	count, err := rgq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rgq *ReleaseGroupQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rgq.ctx, "Exist")
	switch _, err := rgq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rgq *ReleaseGroupQuery) ExistX(ctx context.Context) bool {
	exist, err := rgq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ReleaseGroupQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rgq *ReleaseGroupQuery) Clone() *ReleaseGroupQuery {
	if rgq == nil {
		return nil
	}
	return &ReleaseGroupQuery{
		config:       rgq.config,
		ctx:          rgq.ctx.Clone(),
		order:        append([]releasegroup.OrderOption{}, rgq.order...),
		inters:       append([]Interceptor{}, rgq.inters...),
		predicates:   append([]predicate.ReleaseGroup{}, rgq.predicates...),
		withEpisodes: rgq.withEpisodes.Clone(),
		// clone intermediate query.
		sql:  rgq.sql.Clone(),
		path: rgq.path,
	}
}

// WithEpisodes tells the query-builder to eager-load the nodes that are connected to
// the "episodes" edge. The optional arguments are used to configure the query builder of the edge.
func (rgq *ReleaseGroupQuery) WithEpisodes(opts ...func(*EpisodeQuery)) *ReleaseGroupQuery {
	query := (&EpisodeClient{config: rgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rgq.withEpisodes = query
	return rgq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ReleaseGroup.Query().
//		GroupBy(releasegroup.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (rgq *ReleaseGroupQuery) GroupBy(field string, fields ...string) *ReleaseGroupGroupBy {
	rgq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ReleaseGroupGroupBy{build: rgq}
	grbuild.flds = &rgq.ctx.Fields
	grbuild.label = releasegroup.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.ReleaseGroup.Query().
//		Select(releasegroup.FieldName).
//		Scan(ctx, &v)
func (rgq *ReleaseGroupQuery) Select(fields ...string) *ReleaseGroupSelect {
	rgq.ctx.Fields = append(rgq.ctx.Fields, fields...)
	sbuild := &ReleaseGroupSelect{ReleaseGroupQuery: rgq}
	sbuild.label = releasegroup.Label
	sbuild.flds, sbuild.scan = &rgq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ReleaseGroupSelect configured with the given aggregations.
func (rgq *ReleaseGroupQuery) Aggregate(fns ...AggregateFunc) *ReleaseGroupSelect {
	return rgq.Select().Aggregate(fns...)
}

func (rgq *ReleaseGroupQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rgq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rgq); err != nil {
				return err
			}
		}
	}
	for _, f := range rgq.ctx.Fields {
		if !releasegroup.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if rgq.path != nil {
		prev, err := rgq.path(ctx)
		if err != nil {
			return err
		}
		rgq.sql = prev
	}
	return nil
}

func (rgq *ReleaseGroupQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ReleaseGroup, error) {
	var (
		nodes       = []*ReleaseGroup{}
		_spec       = rgq.querySpec()
		loadedTypes = [1]bool{
			rgq.withEpisodes != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ReleaseGroup).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ReleaseGroup{config: rgq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(rgq.modifiers) > 0 {
		_spec.Modifiers = rgq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rgq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rgq.withEpisodes; query != nil {
		if err := rgq.loadEpisodes(ctx, query, nodes,
			func(n *ReleaseGroup) { n.Edges.Episodes = []*Episode{} },
			func(n *ReleaseGroup, e *Episode) { n.Edges.Episodes = append(n.Edges.Episodes, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range rgq.withNamedEpisodes {
		if err := rgq.loadEpisodes(ctx, query, nodes,
			func(n *ReleaseGroup) { n.appendNamedEpisodes(name) },
			func(n *ReleaseGroup, e *Episode) { n.appendNamedEpisodes(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range rgq.loadTotal {
		if err := rgq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rgq *ReleaseGroupQuery) loadEpisodes(ctx context.Context, query *EpisodeQuery, nodes []*ReleaseGroup, init func(*ReleaseGroup), assign func(*ReleaseGroup, *Episode)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*ReleaseGroup)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Episode(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(releasegroup.EpisodesColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.release_group_id
		if fk == nil {
			return fmt.Errorf(`foreign-key "release_group_id" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "release_group_id" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (rgq *ReleaseGroupQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rgq.querySpec()
	if len(rgq.modifiers) > 0 {
		_spec.Modifiers = rgq.modifiers
	}
	_spec.Node.Columns = rgq.ctx.Fields
	if len(rgq.ctx.Fields) > 0 {
		_spec.Unique = rgq.ctx.Unique != nil && *rgq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rgq.driver, _spec)
}

func (rgq *ReleaseGroupQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(releasegroup.Table, releasegroup.Columns, sqlgraph.NewFieldSpec(releasegroup.FieldID, field.TypeInt))
	_spec.From = rgq.sql
	if unique := rgq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rgq.path != nil {
		_spec.Unique = true
	}
	if fields := rgq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, releasegroup.FieldID)
		for i := range fields {
			if fields[i] != releasegroup.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rgq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rgq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rgq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rgq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rgq *ReleaseGroupQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rgq.driver.Dialect())
	t1 := builder.Table(releasegroup.Table)
	columns := rgq.ctx.Fields
	if len(columns) == 0 {
		columns = releasegroup.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rgq.sql != nil {
		selector = rgq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rgq.ctx.Unique != nil && *rgq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rgq.predicates {
		p(selector)
	}
	for _, p := range rgq.order {
		p(selector)
	}
	if offset := rgq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rgq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedEpisodes tells the query-builder to eager-load the nodes that are connected to the "episodes"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (rgq *ReleaseGroupQuery) WithNamedEpisodes(name string, opts ...func(*EpisodeQuery)) *ReleaseGroupQuery {
	query := (&EpisodeClient{config: rgq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if rgq.withNamedEpisodes == nil {
		rgq.withNamedEpisodes = make(map[string]*EpisodeQuery)
	}
	rgq.withNamedEpisodes[name] = query
	return rgq
}

// ReleaseGroupGroupBy is the group-by builder for ReleaseGroup entities.
type ReleaseGroupGroupBy struct {
	selector
	build *ReleaseGroupQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rggb *ReleaseGroupGroupBy) Aggregate(fns ...AggregateFunc) *ReleaseGroupGroupBy {
	rggb.fns = append(rggb.fns, fns...)
	return rggb
}

// Scan applies the selector query and scans the result into the given value.
func (rggb *ReleaseGroupGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rggb.build.ctx, "GroupBy")
	if err := rggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReleaseGroupQuery, *ReleaseGroupGroupBy](ctx, rggb.build, rggb, rggb.build.inters, v)
}

func (rggb *ReleaseGroupGroupBy) sqlScan(ctx context.Context, root *ReleaseGroupQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rggb.fns))
	for _, fn := range rggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rggb.flds)+len(rggb.fns))
		for _, f := range *rggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ReleaseGroupSelect is the builder for selecting fields of ReleaseGroup entities.
type ReleaseGroupSelect struct {
	*ReleaseGroupQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rgs *ReleaseGroupSelect) Aggregate(fns ...AggregateFunc) *ReleaseGroupSelect {
	rgs.fns = append(rgs.fns, fns...)
	return rgs
}

// Scan applies the selector query and scans the result into the given value.
func (rgs *ReleaseGroupSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rgs.ctx, "Select")
	if err := rgs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ReleaseGroupQuery, *ReleaseGroupSelect](ctx, rgs.ReleaseGroupQuery, rgs, rgs.inters, v)
}

func (rgs *ReleaseGroupSelect) sqlScan(ctx context.Context, root *ReleaseGroupQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rgs.fns))
	for _, fn := range rgs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rgs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rgs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
