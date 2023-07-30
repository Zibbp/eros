// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/zibbp/eros/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/zibbp/eros/ent/report"
	"github.com/zibbp/eros/ent/script"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Report is the client for interacting with the Report builders.
	Report *ReportClient
	// Script is the client for interacting with the Script builders.
	Script *ScriptClient
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
	c.Report = NewReportClient(c.config)
	c.Script = NewScriptClient(c.config)
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
		ctx:    ctx,
		config: cfg,
		Report: NewReportClient(cfg),
		Script: NewScriptClient(cfg),
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
		ctx:    ctx,
		config: cfg,
		Report: NewReportClient(cfg),
		Script: NewScriptClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Report.
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
	c.Report.Use(hooks...)
	c.Script.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Report.Intercept(interceptors...)
	c.Script.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ReportMutation:
		return c.Report.mutate(ctx, m)
	case *ScriptMutation:
		return c.Script.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// ReportClient is a client for the Report schema.
type ReportClient struct {
	config
}

// NewReportClient returns a client for the Report from the given config.
func NewReportClient(c config) *ReportClient {
	return &ReportClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `report.Hooks(f(g(h())))`.
func (c *ReportClient) Use(hooks ...Hook) {
	c.hooks.Report = append(c.hooks.Report, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `report.Intercept(f(g(h())))`.
func (c *ReportClient) Intercept(interceptors ...Interceptor) {
	c.inters.Report = append(c.inters.Report, interceptors...)
}

// Create returns a builder for creating a Report entity.
func (c *ReportClient) Create() *ReportCreate {
	mutation := newReportMutation(c.config, OpCreate)
	return &ReportCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Report entities.
func (c *ReportClient) CreateBulk(builders ...*ReportCreate) *ReportCreateBulk {
	return &ReportCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Report.
func (c *ReportClient) Update() *ReportUpdate {
	mutation := newReportMutation(c.config, OpUpdate)
	return &ReportUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ReportClient) UpdateOne(r *Report) *ReportUpdateOne {
	mutation := newReportMutation(c.config, OpUpdateOne, withReport(r))
	return &ReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ReportClient) UpdateOneID(id uuid.UUID) *ReportUpdateOne {
	mutation := newReportMutation(c.config, OpUpdateOne, withReportID(id))
	return &ReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Report.
func (c *ReportClient) Delete() *ReportDelete {
	mutation := newReportMutation(c.config, OpDelete)
	return &ReportDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ReportClient) DeleteOne(r *Report) *ReportDeleteOne {
	return c.DeleteOneID(r.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ReportClient) DeleteOneID(id uuid.UUID) *ReportDeleteOne {
	builder := c.Delete().Where(report.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ReportDeleteOne{builder}
}

// Query returns a query builder for Report.
func (c *ReportClient) Query() *ReportQuery {
	return &ReportQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeReport},
		inters: c.Interceptors(),
	}
}

// Get returns a Report entity by its id.
func (c *ReportClient) Get(ctx context.Context, id uuid.UUID) (*Report, error) {
	return c.Query().Where(report.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ReportClient) GetX(ctx context.Context, id uuid.UUID) *Report {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryScript queries the script edge of a Report.
func (c *ReportClient) QueryScript(r *Report) *ScriptQuery {
	query := (&ScriptClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := r.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(report.Table, report.FieldID, id),
			sqlgraph.To(script.Table, script.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, report.ScriptTable, report.ScriptColumn),
		)
		fromV = sqlgraph.Neighbors(r.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ReportClient) Hooks() []Hook {
	return c.hooks.Report
}

// Interceptors returns the client interceptors.
func (c *ReportClient) Interceptors() []Interceptor {
	return c.inters.Report
}

func (c *ReportClient) mutate(ctx context.Context, m *ReportMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ReportCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ReportUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ReportUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ReportDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Report mutation op: %q", m.Op())
	}
}

// ScriptClient is a client for the Script schema.
type ScriptClient struct {
	config
}

// NewScriptClient returns a client for the Script from the given config.
func NewScriptClient(c config) *ScriptClient {
	return &ScriptClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `script.Hooks(f(g(h())))`.
func (c *ScriptClient) Use(hooks ...Hook) {
	c.hooks.Script = append(c.hooks.Script, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `script.Intercept(f(g(h())))`.
func (c *ScriptClient) Intercept(interceptors ...Interceptor) {
	c.inters.Script = append(c.inters.Script, interceptors...)
}

// Create returns a builder for creating a Script entity.
func (c *ScriptClient) Create() *ScriptCreate {
	mutation := newScriptMutation(c.config, OpCreate)
	return &ScriptCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Script entities.
func (c *ScriptClient) CreateBulk(builders ...*ScriptCreate) *ScriptCreateBulk {
	return &ScriptCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Script.
func (c *ScriptClient) Update() *ScriptUpdate {
	mutation := newScriptMutation(c.config, OpUpdate)
	return &ScriptUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ScriptClient) UpdateOne(s *Script) *ScriptUpdateOne {
	mutation := newScriptMutation(c.config, OpUpdateOne, withScript(s))
	return &ScriptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ScriptClient) UpdateOneID(id uuid.UUID) *ScriptUpdateOne {
	mutation := newScriptMutation(c.config, OpUpdateOne, withScriptID(id))
	return &ScriptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Script.
func (c *ScriptClient) Delete() *ScriptDelete {
	mutation := newScriptMutation(c.config, OpDelete)
	return &ScriptDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ScriptClient) DeleteOne(s *Script) *ScriptDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ScriptClient) DeleteOneID(id uuid.UUID) *ScriptDeleteOne {
	builder := c.Delete().Where(script.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ScriptDeleteOne{builder}
}

// Query returns a query builder for Script.
func (c *ScriptClient) Query() *ScriptQuery {
	return &ScriptQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeScript},
		inters: c.Interceptors(),
	}
}

// Get returns a Script entity by its id.
func (c *ScriptClient) Get(ctx context.Context, id uuid.UUID) (*Script, error) {
	return c.Query().Where(script.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ScriptClient) GetX(ctx context.Context, id uuid.UUID) *Script {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryReports queries the reports edge of a Script.
func (c *ScriptClient) QueryReports(s *Script) *ReportQuery {
	query := (&ReportClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(script.Table, script.FieldID, id),
			sqlgraph.To(report.Table, report.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, script.ReportsTable, script.ReportsColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ScriptClient) Hooks() []Hook {
	return c.hooks.Script
}

// Interceptors returns the client interceptors.
func (c *ScriptClient) Interceptors() []Interceptor {
	return c.inters.Script
}

func (c *ScriptClient) mutate(ctx context.Context, m *ScriptMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ScriptCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ScriptUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ScriptUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ScriptDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Script mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Report, Script []ent.Hook
	}
	inters struct {
		Report, Script []ent.Interceptor
	}
)