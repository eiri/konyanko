// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/eiri/konyanko/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/eiri/konyanko/ent/anime"
	"github.com/eiri/konyanko/ent/episode"
	"github.com/eiri/konyanko/ent/item"
	"github.com/eiri/konyanko/ent/releasegroup"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Anime is the client for interacting with the Anime builders.
	Anime *AnimeClient
	// Episode is the client for interacting with the Episode builders.
	Episode *EpisodeClient
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// ReleaseGroup is the client for interacting with the ReleaseGroup builders.
	ReleaseGroup *ReleaseGroupClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Anime = NewAnimeClient(c.config)
	c.Episode = NewEpisodeClient(c.config)
	c.Item = NewItemClient(c.config)
	c.ReleaseGroup = NewReleaseGroupClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("ent: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Anime:        NewAnimeClient(cfg),
		Episode:      NewEpisodeClient(cfg),
		Item:         NewItemClient(cfg),
		ReleaseGroup: NewReleaseGroupClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:          ctx,
		config:       cfg,
		Anime:        NewAnimeClient(cfg),
		Episode:      NewEpisodeClient(cfg),
		Item:         NewItemClient(cfg),
		ReleaseGroup: NewReleaseGroupClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Anime.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Anime.Use(hooks...)
	c.Episode.Use(hooks...)
	c.Item.Use(hooks...)
	c.ReleaseGroup.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Anime.Intercept(interceptors...)
	c.Episode.Intercept(interceptors...)
	c.Item.Intercept(interceptors...)
	c.ReleaseGroup.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *AnimeMutation:
		return c.Anime.mutate(ctx, m)
	case *EpisodeMutation:
		return c.Episode.mutate(ctx, m)
	case *ItemMutation:
		return c.Item.mutate(ctx, m)
	case *ReleaseGroupMutation:
		return c.ReleaseGroup.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// AnimeClient is a client for the Anime schema.
type AnimeClient struct {
	config
}

// NewAnimeClient returns a client for the Anime from the given config.
func NewAnimeClient(c config) *AnimeClient {
	return &AnimeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `anime.Hooks(f(g(h())))`.
func (c *AnimeClient) Use(hooks ...Hook) {
	c.hooks.Anime = append(c.hooks.Anime, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `anime.Intercept(f(g(h())))`.
func (c *AnimeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Anime = append(c.inters.Anime, interceptors...)
}

// Create returns a builder for creating a Anime entity.
func (c *AnimeClient) Create() *AnimeCreate {
	mutation := newAnimeMutation(c.config, OpCreate)
	return &AnimeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Anime entities.
func (c *AnimeClient) CreateBulk(builders ...*AnimeCreate) *AnimeCreateBulk {
	return &AnimeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *AnimeClient) MapCreateBulk(slice any, setFunc func(*AnimeCreate, int)) *AnimeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &AnimeCreateBulk{err: fmt.Errorf("calling to AnimeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*AnimeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &AnimeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Anime.
func (c *AnimeClient) Update() *AnimeUpdate {
	mutation := newAnimeMutation(c.config, OpUpdate)
	return &AnimeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *AnimeClient) UpdateOne(a *Anime) *AnimeUpdateOne {
	mutation := newAnimeMutation(c.config, OpUpdateOne, withAnime(a))
	return &AnimeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *AnimeClient) UpdateOneID(id int) *AnimeUpdateOne {
	mutation := newAnimeMutation(c.config, OpUpdateOne, withAnimeID(id))
	return &AnimeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Anime.
func (c *AnimeClient) Delete() *AnimeDelete {
	mutation := newAnimeMutation(c.config, OpDelete)
	return &AnimeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *AnimeClient) DeleteOne(a *Anime) *AnimeDeleteOne {
	return c.DeleteOneID(a.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *AnimeClient) DeleteOneID(id int) *AnimeDeleteOne {
	builder := c.Delete().Where(anime.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &AnimeDeleteOne{builder}
}

// Query returns a query builder for Anime.
func (c *AnimeClient) Query() *AnimeQuery {
	return &AnimeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeAnime},
		inters: c.Interceptors(),
	}
}

// Get returns a Anime entity by its id.
func (c *AnimeClient) Get(ctx context.Context, id int) (*Anime, error) {
	return c.Query().Where(anime.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *AnimeClient) GetX(ctx context.Context, id int) *Anime {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEpisodes queries the episodes edge of a Anime.
func (c *AnimeClient) QueryEpisodes(a *Anime) *EpisodeQuery {
	query := (&EpisodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := a.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(anime.Table, anime.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, anime.EpisodesTable, anime.EpisodesColumn),
		)
		fromV = sqlgraph.Neighbors(a.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *AnimeClient) Hooks() []Hook {
	return c.hooks.Anime
}

// Interceptors returns the client interceptors.
func (c *AnimeClient) Interceptors() []Interceptor {
	return c.inters.Anime
}

func (c *AnimeClient) mutate(ctx context.Context, m *AnimeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&AnimeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&AnimeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&AnimeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&AnimeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Anime mutation op: %q", m.Op())
	}
}

// EpisodeClient is a client for the Episode schema.
type EpisodeClient struct {
	config
}

// NewEpisodeClient returns a client for the Episode from the given config.
func NewEpisodeClient(c config) *EpisodeClient {
	return &EpisodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `episode.Hooks(f(g(h())))`.
func (c *EpisodeClient) Use(hooks ...Hook) {
	c.hooks.Episode = append(c.hooks.Episode, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `episode.Intercept(f(g(h())))`.
func (c *EpisodeClient) Intercept(interceptors ...Interceptor) {
	c.inters.Episode = append(c.inters.Episode, interceptors...)
}

// Create returns a builder for creating a Episode entity.
func (c *EpisodeClient) Create() *EpisodeCreate {
	mutation := newEpisodeMutation(c.config, OpCreate)
	return &EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Episode entities.
func (c *EpisodeClient) CreateBulk(builders ...*EpisodeCreate) *EpisodeCreateBulk {
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *EpisodeClient) MapCreateBulk(slice any, setFunc func(*EpisodeCreate, int)) *EpisodeCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &EpisodeCreateBulk{err: fmt.Errorf("calling to EpisodeClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*EpisodeCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &EpisodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Episode.
func (c *EpisodeClient) Update() *EpisodeUpdate {
	mutation := newEpisodeMutation(c.config, OpUpdate)
	return &EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *EpisodeClient) UpdateOne(e *Episode) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisode(e))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *EpisodeClient) UpdateOneID(id int) *EpisodeUpdateOne {
	mutation := newEpisodeMutation(c.config, OpUpdateOne, withEpisodeID(id))
	return &EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Episode.
func (c *EpisodeClient) Delete() *EpisodeDelete {
	mutation := newEpisodeMutation(c.config, OpDelete)
	return &EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *EpisodeClient) DeleteOne(e *Episode) *EpisodeDeleteOne {
	return c.DeleteOneID(e.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *EpisodeClient) DeleteOneID(id int) *EpisodeDeleteOne {
	builder := c.Delete().Where(episode.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &EpisodeDeleteOne{builder}
}

// Query returns a query builder for Episode.
func (c *EpisodeClient) Query() *EpisodeQuery {
	return &EpisodeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeEpisode},
		inters: c.Interceptors(),
	}
}

// Get returns a Episode entity by its id.
func (c *EpisodeClient) Get(ctx context.Context, id int) (*Episode, error) {
	return c.Query().Where(episode.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *EpisodeClient) GetX(ctx context.Context, id int) *Episode {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryItem queries the item edge of a Episode.
func (c *EpisodeClient) QueryItem(e *Episode) *ItemQuery {
	query := (&ItemClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(item.Table, item.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, episode.ItemTable, episode.ItemColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTitle queries the title edge of a Episode.
func (c *EpisodeClient) QueryTitle(e *Episode) *AnimeQuery {
	query := (&AnimeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(anime.Table, anime.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, episode.TitleTable, episode.TitleColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryReleaseGroup queries the release_group edge of a Episode.
func (c *EpisodeClient) QueryReleaseGroup(e *Episode) *ReleaseGroupQuery {
	query := (&ReleaseGroupClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := e.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(episode.Table, episode.FieldID, id),
			sqlgraph.To(releasegroup.Table, releasegroup.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, episode.ReleaseGroupTable, episode.ReleaseGroupColumn),
		)
		fromV = sqlgraph.Neighbors(e.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *EpisodeClient) Hooks() []Hook {
	return c.hooks.Episode
}

// Interceptors returns the client interceptors.
func (c *EpisodeClient) Interceptors() []Interceptor {
	return c.inters.Episode
}

func (c *EpisodeClient) mutate(ctx context.Context, m *EpisodeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&EpisodeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&EpisodeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&EpisodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&EpisodeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Episode mutation op: %q", m.Op())
	}
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `item.Hooks(f(g(h())))`.
func (c *ItemClient) Use(hooks ...Hook) {
	c.hooks.Item = append(c.hooks.Item, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `item.Intercept(f(g(h())))`.
func (c *ItemClient) Intercept(interceptors ...Interceptor) {
	c.inters.Item = append(c.inters.Item, interceptors...)
}

// Create returns a builder for creating a Item entity.
func (c *ItemClient) Create() *ItemCreate {
	mutation := newItemMutation(c.config, OpCreate)
	return &ItemCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Item entities.
func (c *ItemClient) CreateBulk(builders ...*ItemCreate) *ItemCreateBulk {
	return &ItemCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ItemClient) MapCreateBulk(slice any, setFunc func(*ItemCreate, int)) *ItemCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ItemCreateBulk{err: fmt.Errorf("calling to ItemClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ItemCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ItemCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	mutation := newItemMutation(c.config, OpUpdate)
	return &ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItem(i))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id int) *ItemUpdateOne {
	mutation := newItemMutation(c.config, OpUpdateOne, withItemID(id))
	return &ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	mutation := newItemMutation(c.config, OpDelete)
	return &ItemDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ItemClient) DeleteOneID(id int) *ItemDeleteOne {
	builder := c.Delete().Where(item.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ItemDeleteOne{builder}
}

// Query returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeItem},
		inters: c.Interceptors(),
	}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id int) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id int) *Item {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEpisodes queries the episodes edge of a Item.
func (c *ItemClient) QueryEpisodes(i *Item) *EpisodeQuery {
	query := (&EpisodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(item.Table, item.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, item.EpisodesTable, item.EpisodesColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ItemClient) Hooks() []Hook {
	return c.hooks.Item
}

// Interceptors returns the client interceptors.
func (c *ItemClient) Interceptors() []Interceptor {
	return c.inters.Item
}

func (c *ItemClient) mutate(ctx context.Context, m *ItemMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ItemCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ItemUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ItemUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ItemDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Item mutation op: %q", m.Op())
	}
}

// ReleaseGroupClient is a client for the ReleaseGroup schema.
type ReleaseGroupClient struct {
	config
}

// NewReleaseGroupClient returns a client for the ReleaseGroup from the given config.
func NewReleaseGroupClient(c config) *ReleaseGroupClient {
	return &ReleaseGroupClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `releasegroup.Hooks(f(g(h())))`.
func (c *ReleaseGroupClient) Use(hooks ...Hook) {
	c.hooks.ReleaseGroup = append(c.hooks.ReleaseGroup, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `releasegroup.Intercept(f(g(h())))`.
func (c *ReleaseGroupClient) Intercept(interceptors ...Interceptor) {
	c.inters.ReleaseGroup = append(c.inters.ReleaseGroup, interceptors...)
}

// Create returns a builder for creating a ReleaseGroup entity.
func (c *ReleaseGroupClient) Create() *ReleaseGroupCreate {
	mutation := newReleaseGroupMutation(c.config, OpCreate)
	return &ReleaseGroupCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ReleaseGroup entities.
func (c *ReleaseGroupClient) CreateBulk(builders ...*ReleaseGroupCreate) *ReleaseGroupCreateBulk {
	return &ReleaseGroupCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ReleaseGroupClient) MapCreateBulk(slice any, setFunc func(*ReleaseGroupCreate, int)) *ReleaseGroupCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ReleaseGroupCreateBulk{err: fmt.Errorf("calling to ReleaseGroupClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ReleaseGroupCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ReleaseGroupCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ReleaseGroup.
func (c *ReleaseGroupClient) Update() *ReleaseGroupUpdate {
	mutation := newReleaseGroupMutation(c.config, OpUpdate)
	return &ReleaseGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReleaseGroupClient) UpdateOne(rg *ReleaseGroup) *ReleaseGroupUpdateOne {
	mutation := newReleaseGroupMutation(c.config, OpUpdateOne, withReleaseGroup(rg))
	return &ReleaseGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReleaseGroupClient) UpdateOneID(id int) *ReleaseGroupUpdateOne {
	mutation := newReleaseGroupMutation(c.config, OpUpdateOne, withReleaseGroupID(id))
	return &ReleaseGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ReleaseGroup.
func (c *ReleaseGroupClient) Delete() *ReleaseGroupDelete {
	mutation := newReleaseGroupMutation(c.config, OpDelete)
	return &ReleaseGroupDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReleaseGroupClient) DeleteOne(rg *ReleaseGroup) *ReleaseGroupDeleteOne {
	return c.DeleteOneID(rg.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ReleaseGroupClient) DeleteOneID(id int) *ReleaseGroupDeleteOne {
	builder := c.Delete().Where(releasegroup.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReleaseGroupDeleteOne{builder}
}

// Query returns a query builder for ReleaseGroup.
func (c *ReleaseGroupClient) Query() *ReleaseGroupQuery {
	return &ReleaseGroupQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeReleaseGroup},
		inters: c.Interceptors(),
	}
}

// Get returns a ReleaseGroup entity by its id.
func (c *ReleaseGroupClient) Get(ctx context.Context, id int) (*ReleaseGroup, error) {
	return c.Query().Where(releasegroup.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReleaseGroupClient) GetX(ctx context.Context, id int) *ReleaseGroup {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryEpisodes queries the episodes edge of a ReleaseGroup.
func (c *ReleaseGroupClient) QueryEpisodes(rg *ReleaseGroup) *EpisodeQuery {
	query := (&EpisodeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rg.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(releasegroup.Table, releasegroup.FieldID, id),
			sqlgraph.To(episode.Table, episode.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, releasegroup.EpisodesTable, releasegroup.EpisodesColumn),
		)
		fromV = sqlgraph.Neighbors(rg.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReleaseGroupClient) Hooks() []Hook {
	return c.hooks.ReleaseGroup
}

// Interceptors returns the client interceptors.
func (c *ReleaseGroupClient) Interceptors() []Interceptor {
	return c.inters.ReleaseGroup
}

func (c *ReleaseGroupClient) mutate(ctx context.Context, m *ReleaseGroupMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ReleaseGroupCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ReleaseGroupUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ReleaseGroupUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ReleaseGroupDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown ReleaseGroup mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Anime, Episode, Item, ReleaseGroup []ent.Hook
	}
	inters struct {
		Anime, Episode, Item, ReleaseGroup []ent.Interceptor
	}
)
