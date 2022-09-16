// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/sergot/tibiago/ent/migrate"

	"github.com/sergot/tibiago/ent/boss"
	"github.com/sergot/tibiago/ent/bosslist"
	"github.com/sergot/tibiago/ent/instance"
	"github.com/sergot/tibiago/ent/instanceconfig"
	"github.com/sergot/tibiago/ent/participant"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Boss is the client for interacting with the Boss builders.
	Boss *BossClient
	// Bosslist is the client for interacting with the Bosslist builders.
	Bosslist *BosslistClient
	// Instance is the client for interacting with the Instance builders.
	Instance *InstanceClient
	// InstanceConfig is the client for interacting with the InstanceConfig builders.
	InstanceConfig *InstanceConfigClient
	// Participant is the client for interacting with the Participant builders.
	Participant *ParticipantClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Boss = NewBossClient(c.config)
	c.Bosslist = NewBosslistClient(c.config)
	c.Instance = NewInstanceClient(c.config)
	c.InstanceConfig = NewInstanceConfigClient(c.config)
	c.Participant = NewParticipantClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Boss:           NewBossClient(cfg),
		Bosslist:       NewBosslistClient(cfg),
		Instance:       NewInstanceClient(cfg),
		InstanceConfig: NewInstanceConfigClient(cfg),
		Participant:    NewParticipantClient(cfg),
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
		ctx:            ctx,
		config:         cfg,
		Boss:           NewBossClient(cfg),
		Bosslist:       NewBosslistClient(cfg),
		Instance:       NewInstanceClient(cfg),
		InstanceConfig: NewInstanceConfigClient(cfg),
		Participant:    NewParticipantClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Boss.
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
	c.Boss.Use(hooks...)
	c.Bosslist.Use(hooks...)
	c.Instance.Use(hooks...)
	c.InstanceConfig.Use(hooks...)
	c.Participant.Use(hooks...)
}

// BossClient is a client for the Boss schema.
type BossClient struct {
	config
}

// NewBossClient returns a client for the Boss from the given config.
func NewBossClient(c config) *BossClient {
	return &BossClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `boss.Hooks(f(g(h())))`.
func (c *BossClient) Use(hooks ...Hook) {
	c.hooks.Boss = append(c.hooks.Boss, hooks...)
}

// Create returns a builder for creating a Boss entity.
func (c *BossClient) Create() *BossCreate {
	mutation := newBossMutation(c.config, OpCreate)
	return &BossCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Boss entities.
func (c *BossClient) CreateBulk(builders ...*BossCreate) *BossCreateBulk {
	return &BossCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Boss.
func (c *BossClient) Update() *BossUpdate {
	mutation := newBossMutation(c.config, OpUpdate)
	return &BossUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BossClient) UpdateOne(b *Boss) *BossUpdateOne {
	mutation := newBossMutation(c.config, OpUpdateOne, withBoss(b))
	return &BossUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BossClient) UpdateOneID(id uuid.UUID) *BossUpdateOne {
	mutation := newBossMutation(c.config, OpUpdateOne, withBossID(id))
	return &BossUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Boss.
func (c *BossClient) Delete() *BossDelete {
	mutation := newBossMutation(c.config, OpDelete)
	return &BossDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BossClient) DeleteOne(b *Boss) *BossDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BossClient) DeleteOneID(id uuid.UUID) *BossDeleteOne {
	builder := c.Delete().Where(boss.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BossDeleteOne{builder}
}

// Query returns a query builder for Boss.
func (c *BossClient) Query() *BossQuery {
	return &BossQuery{
		config: c.config,
	}
}

// Get returns a Boss entity by its id.
func (c *BossClient) Get(ctx context.Context, id uuid.UUID) (*Boss, error) {
	return c.Query().Where(boss.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BossClient) GetX(ctx context.Context, id uuid.UUID) *Boss {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBosslists queries the bosslists edge of a Boss.
func (c *BossClient) QueryBosslists(b *Boss) *BosslistQuery {
	query := &BosslistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(boss.Table, boss.FieldID, id),
			sqlgraph.To(bosslist.Table, bosslist.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, boss.BosslistsTable, boss.BosslistsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BossClient) Hooks() []Hook {
	return c.hooks.Boss
}

// BosslistClient is a client for the Bosslist schema.
type BosslistClient struct {
	config
}

// NewBosslistClient returns a client for the Bosslist from the given config.
func NewBosslistClient(c config) *BosslistClient {
	return &BosslistClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `bosslist.Hooks(f(g(h())))`.
func (c *BosslistClient) Use(hooks ...Hook) {
	c.hooks.Bosslist = append(c.hooks.Bosslist, hooks...)
}

// Create returns a builder for creating a Bosslist entity.
func (c *BosslistClient) Create() *BosslistCreate {
	mutation := newBosslistMutation(c.config, OpCreate)
	return &BosslistCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Bosslist entities.
func (c *BosslistClient) CreateBulk(builders ...*BosslistCreate) *BosslistCreateBulk {
	return &BosslistCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Bosslist.
func (c *BosslistClient) Update() *BosslistUpdate {
	mutation := newBosslistMutation(c.config, OpUpdate)
	return &BosslistUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BosslistClient) UpdateOne(b *Bosslist) *BosslistUpdateOne {
	mutation := newBosslistMutation(c.config, OpUpdateOne, withBosslist(b))
	return &BosslistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BosslistClient) UpdateOneID(id uuid.UUID) *BosslistUpdateOne {
	mutation := newBosslistMutation(c.config, OpUpdateOne, withBosslistID(id))
	return &BosslistUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Bosslist.
func (c *BosslistClient) Delete() *BosslistDelete {
	mutation := newBosslistMutation(c.config, OpDelete)
	return &BosslistDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BosslistClient) DeleteOne(b *Bosslist) *BosslistDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BosslistClient) DeleteOneID(id uuid.UUID) *BosslistDeleteOne {
	builder := c.Delete().Where(bosslist.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BosslistDeleteOne{builder}
}

// Query returns a query builder for Bosslist.
func (c *BosslistClient) Query() *BosslistQuery {
	return &BosslistQuery{
		config: c.config,
	}
}

// Get returns a Bosslist entity by its id.
func (c *BosslistClient) Get(ctx context.Context, id uuid.UUID) (*Bosslist, error) {
	return c.Query().Where(bosslist.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BosslistClient) GetX(ctx context.Context, id uuid.UUID) *Bosslist {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBoss queries the boss edge of a Bosslist.
func (c *BosslistClient) QueryBoss(b *Bosslist) *BossQuery {
	query := &BossQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bosslist.Table, bosslist.FieldID, id),
			sqlgraph.To(boss.Table, boss.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, bosslist.BossTable, bosslist.BossColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryParticipants queries the participants edge of a Bosslist.
func (c *BosslistClient) QueryParticipants(b *Bosslist) *ParticipantQuery {
	query := &ParticipantQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := b.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(bosslist.Table, bosslist.FieldID, id),
			sqlgraph.To(participant.Table, participant.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, bosslist.ParticipantsTable, bosslist.ParticipantsColumn),
		)
		fromV = sqlgraph.Neighbors(b.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BosslistClient) Hooks() []Hook {
	return c.hooks.Bosslist
}

// InstanceClient is a client for the Instance schema.
type InstanceClient struct {
	config
}

// NewInstanceClient returns a client for the Instance from the given config.
func NewInstanceClient(c config) *InstanceClient {
	return &InstanceClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instance.Hooks(f(g(h())))`.
func (c *InstanceClient) Use(hooks ...Hook) {
	c.hooks.Instance = append(c.hooks.Instance, hooks...)
}

// Create returns a builder for creating a Instance entity.
func (c *InstanceClient) Create() *InstanceCreate {
	mutation := newInstanceMutation(c.config, OpCreate)
	return &InstanceCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Instance entities.
func (c *InstanceClient) CreateBulk(builders ...*InstanceCreate) *InstanceCreateBulk {
	return &InstanceCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Instance.
func (c *InstanceClient) Update() *InstanceUpdate {
	mutation := newInstanceMutation(c.config, OpUpdate)
	return &InstanceUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstanceClient) UpdateOne(i *Instance) *InstanceUpdateOne {
	mutation := newInstanceMutation(c.config, OpUpdateOne, withInstance(i))
	return &InstanceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstanceClient) UpdateOneID(id uuid.UUID) *InstanceUpdateOne {
	mutation := newInstanceMutation(c.config, OpUpdateOne, withInstanceID(id))
	return &InstanceUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Instance.
func (c *InstanceClient) Delete() *InstanceDelete {
	mutation := newInstanceMutation(c.config, OpDelete)
	return &InstanceDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *InstanceClient) DeleteOne(i *Instance) *InstanceDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *InstanceClient) DeleteOneID(id uuid.UUID) *InstanceDeleteOne {
	builder := c.Delete().Where(instance.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstanceDeleteOne{builder}
}

// Query returns a query builder for Instance.
func (c *InstanceClient) Query() *InstanceQuery {
	return &InstanceQuery{
		config: c.config,
	}
}

// Get returns a Instance entity by its id.
func (c *InstanceClient) Get(ctx context.Context, id uuid.UUID) (*Instance, error) {
	return c.Query().Where(instance.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstanceClient) GetX(ctx context.Context, id uuid.UUID) *Instance {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryConfig queries the config edge of a Instance.
func (c *InstanceClient) QueryConfig(i *Instance) *InstanceConfigQuery {
	query := &InstanceConfigQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := i.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instance.Table, instance.FieldID, id),
			sqlgraph.To(instanceconfig.Table, instanceconfig.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, instance.ConfigTable, instance.ConfigColumn),
		)
		fromV = sqlgraph.Neighbors(i.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstanceClient) Hooks() []Hook {
	return c.hooks.Instance
}

// InstanceConfigClient is a client for the InstanceConfig schema.
type InstanceConfigClient struct {
	config
}

// NewInstanceConfigClient returns a client for the InstanceConfig from the given config.
func NewInstanceConfigClient(c config) *InstanceConfigClient {
	return &InstanceConfigClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `instanceconfig.Hooks(f(g(h())))`.
func (c *InstanceConfigClient) Use(hooks ...Hook) {
	c.hooks.InstanceConfig = append(c.hooks.InstanceConfig, hooks...)
}

// Create returns a builder for creating a InstanceConfig entity.
func (c *InstanceConfigClient) Create() *InstanceConfigCreate {
	mutation := newInstanceConfigMutation(c.config, OpCreate)
	return &InstanceConfigCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of InstanceConfig entities.
func (c *InstanceConfigClient) CreateBulk(builders ...*InstanceConfigCreate) *InstanceConfigCreateBulk {
	return &InstanceConfigCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for InstanceConfig.
func (c *InstanceConfigClient) Update() *InstanceConfigUpdate {
	mutation := newInstanceConfigMutation(c.config, OpUpdate)
	return &InstanceConfigUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *InstanceConfigClient) UpdateOne(ic *InstanceConfig) *InstanceConfigUpdateOne {
	mutation := newInstanceConfigMutation(c.config, OpUpdateOne, withInstanceConfig(ic))
	return &InstanceConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *InstanceConfigClient) UpdateOneID(id uuid.UUID) *InstanceConfigUpdateOne {
	mutation := newInstanceConfigMutation(c.config, OpUpdateOne, withInstanceConfigID(id))
	return &InstanceConfigUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for InstanceConfig.
func (c *InstanceConfigClient) Delete() *InstanceConfigDelete {
	mutation := newInstanceConfigMutation(c.config, OpDelete)
	return &InstanceConfigDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *InstanceConfigClient) DeleteOne(ic *InstanceConfig) *InstanceConfigDeleteOne {
	return c.DeleteOneID(ic.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *InstanceConfigClient) DeleteOneID(id uuid.UUID) *InstanceConfigDeleteOne {
	builder := c.Delete().Where(instanceconfig.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &InstanceConfigDeleteOne{builder}
}

// Query returns a query builder for InstanceConfig.
func (c *InstanceConfigClient) Query() *InstanceConfigQuery {
	return &InstanceConfigQuery{
		config: c.config,
	}
}

// Get returns a InstanceConfig entity by its id.
func (c *InstanceConfigClient) Get(ctx context.Context, id uuid.UUID) (*InstanceConfig, error) {
	return c.Query().Where(instanceconfig.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *InstanceConfigClient) GetX(ctx context.Context, id uuid.UUID) *InstanceConfig {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryInstance queries the instance edge of a InstanceConfig.
func (c *InstanceConfigClient) QueryInstance(ic *InstanceConfig) *InstanceQuery {
	query := &InstanceQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ic.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(instanceconfig.Table, instanceconfig.FieldID, id),
			sqlgraph.To(instance.Table, instance.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, instanceconfig.InstanceTable, instanceconfig.InstanceColumn),
		)
		fromV = sqlgraph.Neighbors(ic.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *InstanceConfigClient) Hooks() []Hook {
	return c.hooks.InstanceConfig
}

// ParticipantClient is a client for the Participant schema.
type ParticipantClient struct {
	config
}

// NewParticipantClient returns a client for the Participant from the given config.
func NewParticipantClient(c config) *ParticipantClient {
	return &ParticipantClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `participant.Hooks(f(g(h())))`.
func (c *ParticipantClient) Use(hooks ...Hook) {
	c.hooks.Participant = append(c.hooks.Participant, hooks...)
}

// Create returns a builder for creating a Participant entity.
func (c *ParticipantClient) Create() *ParticipantCreate {
	mutation := newParticipantMutation(c.config, OpCreate)
	return &ParticipantCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Participant entities.
func (c *ParticipantClient) CreateBulk(builders ...*ParticipantCreate) *ParticipantCreateBulk {
	return &ParticipantCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Participant.
func (c *ParticipantClient) Update() *ParticipantUpdate {
	mutation := newParticipantMutation(c.config, OpUpdate)
	return &ParticipantUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ParticipantClient) UpdateOne(pa *Participant) *ParticipantUpdateOne {
	mutation := newParticipantMutation(c.config, OpUpdateOne, withParticipant(pa))
	return &ParticipantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ParticipantClient) UpdateOneID(id uuid.UUID) *ParticipantUpdateOne {
	mutation := newParticipantMutation(c.config, OpUpdateOne, withParticipantID(id))
	return &ParticipantUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Participant.
func (c *ParticipantClient) Delete() *ParticipantDelete {
	mutation := newParticipantMutation(c.config, OpDelete)
	return &ParticipantDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ParticipantClient) DeleteOne(pa *Participant) *ParticipantDeleteOne {
	return c.DeleteOneID(pa.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ParticipantClient) DeleteOneID(id uuid.UUID) *ParticipantDeleteOne {
	builder := c.Delete().Where(participant.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ParticipantDeleteOne{builder}
}

// Query returns a query builder for Participant.
func (c *ParticipantClient) Query() *ParticipantQuery {
	return &ParticipantQuery{
		config: c.config,
	}
}

// Get returns a Participant entity by its id.
func (c *ParticipantClient) Get(ctx context.Context, id uuid.UUID) (*Participant, error) {
	return c.Query().Where(participant.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ParticipantClient) GetX(ctx context.Context, id uuid.UUID) *Participant {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBosslist queries the bosslist edge of a Participant.
func (c *ParticipantClient) QueryBosslist(pa *Participant) *BosslistQuery {
	query := &BosslistQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := pa.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(participant.Table, participant.FieldID, id),
			sqlgraph.To(bosslist.Table, bosslist.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, participant.BosslistTable, participant.BosslistColumn),
		)
		fromV = sqlgraph.Neighbors(pa.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ParticipantClient) Hooks() []Hook {
	return c.hooks.Participant
}
