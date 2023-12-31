// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"moonlogs/ent/migrate"

	"moonlogs/ent/logrecord"
	"moonlogs/ent/logschema"
	"moonlogs/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// LogRecord is the client for interacting with the LogRecord builders.
	LogRecord *LogRecordClient
	// LogSchema is the client for interacting with the LogSchema builders.
	LogSchema *LogSchemaClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.LogRecord = NewLogRecordClient(c.config)
	c.LogSchema = NewLogSchemaClient(c.config)
	c.User = NewUserClient(c.config)
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
		ctx:       ctx,
		config:    cfg,
		LogRecord: NewLogRecordClient(cfg),
		LogSchema: NewLogSchemaClient(cfg),
		User:      NewUserClient(cfg),
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
		ctx:       ctx,
		config:    cfg,
		LogRecord: NewLogRecordClient(cfg),
		LogSchema: NewLogSchemaClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		LogRecord.
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
	c.LogRecord.Use(hooks...)
	c.LogSchema.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.LogRecord.Intercept(interceptors...)
	c.LogSchema.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LogRecordMutation:
		return c.LogRecord.mutate(ctx, m)
	case *LogSchemaMutation:
		return c.LogSchema.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// LogRecordClient is a client for the LogRecord schema.
type LogRecordClient struct {
	config
}

// NewLogRecordClient returns a client for the LogRecord from the given config.
func NewLogRecordClient(c config) *LogRecordClient {
	return &LogRecordClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `logrecord.Hooks(f(g(h())))`.
func (c *LogRecordClient) Use(hooks ...Hook) {
	c.hooks.LogRecord = append(c.hooks.LogRecord, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `logrecord.Intercept(f(g(h())))`.
func (c *LogRecordClient) Intercept(interceptors ...Interceptor) {
	c.inters.LogRecord = append(c.inters.LogRecord, interceptors...)
}

// Create returns a builder for creating a LogRecord entity.
func (c *LogRecordClient) Create() *LogRecordCreate {
	mutation := newLogRecordMutation(c.config, OpCreate)
	return &LogRecordCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LogRecord entities.
func (c *LogRecordClient) CreateBulk(builders ...*LogRecordCreate) *LogRecordCreateBulk {
	return &LogRecordCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LogRecordClient) MapCreateBulk(slice any, setFunc func(*LogRecordCreate, int)) *LogRecordCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LogRecordCreateBulk{err: fmt.Errorf("calling to LogRecordClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LogRecordCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LogRecordCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LogRecord.
func (c *LogRecordClient) Update() *LogRecordUpdate {
	mutation := newLogRecordMutation(c.config, OpUpdate)
	return &LogRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LogRecordClient) UpdateOne(lr *LogRecord) *LogRecordUpdateOne {
	mutation := newLogRecordMutation(c.config, OpUpdateOne, withLogRecord(lr))
	return &LogRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LogRecordClient) UpdateOneID(id int) *LogRecordUpdateOne {
	mutation := newLogRecordMutation(c.config, OpUpdateOne, withLogRecordID(id))
	return &LogRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LogRecord.
func (c *LogRecordClient) Delete() *LogRecordDelete {
	mutation := newLogRecordMutation(c.config, OpDelete)
	return &LogRecordDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LogRecordClient) DeleteOne(lr *LogRecord) *LogRecordDeleteOne {
	return c.DeleteOneID(lr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LogRecordClient) DeleteOneID(id int) *LogRecordDeleteOne {
	builder := c.Delete().Where(logrecord.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LogRecordDeleteOne{builder}
}

// Query returns a query builder for LogRecord.
func (c *LogRecordClient) Query() *LogRecordQuery {
	return &LogRecordQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLogRecord},
		inters: c.Interceptors(),
	}
}

// Get returns a LogRecord entity by its id.
func (c *LogRecordClient) Get(ctx context.Context, id int) (*LogRecord, error) {
	return c.Query().Where(logrecord.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LogRecordClient) GetX(ctx context.Context, id int) *LogRecord {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LogRecordClient) Hooks() []Hook {
	return c.hooks.LogRecord
}

// Interceptors returns the client interceptors.
func (c *LogRecordClient) Interceptors() []Interceptor {
	return c.inters.LogRecord
}

func (c *LogRecordClient) mutate(ctx context.Context, m *LogRecordMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LogRecordCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LogRecordUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LogRecordUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LogRecordDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LogRecord mutation op: %q", m.Op())
	}
}

// LogSchemaClient is a client for the LogSchema schema.
type LogSchemaClient struct {
	config
}

// NewLogSchemaClient returns a client for the LogSchema from the given config.
func NewLogSchemaClient(c config) *LogSchemaClient {
	return &LogSchemaClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `logschema.Hooks(f(g(h())))`.
func (c *LogSchemaClient) Use(hooks ...Hook) {
	c.hooks.LogSchema = append(c.hooks.LogSchema, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `logschema.Intercept(f(g(h())))`.
func (c *LogSchemaClient) Intercept(interceptors ...Interceptor) {
	c.inters.LogSchema = append(c.inters.LogSchema, interceptors...)
}

// Create returns a builder for creating a LogSchema entity.
func (c *LogSchemaClient) Create() *LogSchemaCreate {
	mutation := newLogSchemaMutation(c.config, OpCreate)
	return &LogSchemaCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of LogSchema entities.
func (c *LogSchemaClient) CreateBulk(builders ...*LogSchemaCreate) *LogSchemaCreateBulk {
	return &LogSchemaCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *LogSchemaClient) MapCreateBulk(slice any, setFunc func(*LogSchemaCreate, int)) *LogSchemaCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &LogSchemaCreateBulk{err: fmt.Errorf("calling to LogSchemaClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*LogSchemaCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &LogSchemaCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for LogSchema.
func (c *LogSchemaClient) Update() *LogSchemaUpdate {
	mutation := newLogSchemaMutation(c.config, OpUpdate)
	return &LogSchemaUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LogSchemaClient) UpdateOne(ls *LogSchema) *LogSchemaUpdateOne {
	mutation := newLogSchemaMutation(c.config, OpUpdateOne, withLogSchema(ls))
	return &LogSchemaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LogSchemaClient) UpdateOneID(id int) *LogSchemaUpdateOne {
	mutation := newLogSchemaMutation(c.config, OpUpdateOne, withLogSchemaID(id))
	return &LogSchemaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for LogSchema.
func (c *LogSchemaClient) Delete() *LogSchemaDelete {
	mutation := newLogSchemaMutation(c.config, OpDelete)
	return &LogSchemaDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LogSchemaClient) DeleteOne(ls *LogSchema) *LogSchemaDeleteOne {
	return c.DeleteOneID(ls.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LogSchemaClient) DeleteOneID(id int) *LogSchemaDeleteOne {
	builder := c.Delete().Where(logschema.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LogSchemaDeleteOne{builder}
}

// Query returns a query builder for LogSchema.
func (c *LogSchemaClient) Query() *LogSchemaQuery {
	return &LogSchemaQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLogSchema},
		inters: c.Interceptors(),
	}
}

// Get returns a LogSchema entity by its id.
func (c *LogSchemaClient) Get(ctx context.Context, id int) (*LogSchema, error) {
	return c.Query().Where(logschema.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LogSchemaClient) GetX(ctx context.Context, id int) *LogSchema {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *LogSchemaClient) Hooks() []Hook {
	return c.hooks.LogSchema
}

// Interceptors returns the client interceptors.
func (c *LogSchemaClient) Interceptors() []Interceptor {
	return c.inters.LogSchema
}

func (c *LogSchemaClient) mutate(ctx context.Context, m *LogSchemaMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LogSchemaCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LogSchemaUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LogSchemaUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LogSchemaDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown LogSchema mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *UserClient) MapCreateBulk(slice any, setFunc func(*UserCreate, int)) *UserCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &UserCreateBulk{err: fmt.Errorf("calling to UserClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*UserCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		LogRecord, LogSchema, User []ent.Hook
	}
	inters struct {
		LogRecord, LogSchema, User []ent.Interceptor
	}
)
