// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"dislinkt/connection_service/ent/migrate"

	"dislinkt/connection_service/ent/blockeduser"
	"dislinkt/connection_service/ent/connection"
	"dislinkt/connection_service/ent/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// BlockedUser is the client for interacting with the BlockedUser builders.
	BlockedUser *BlockedUserClient
	// Connection is the client for interacting with the Connection builders.
	Connection *ConnectionClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.BlockedUser = NewBlockedUserClient(c.config)
	c.Connection = NewConnectionClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		BlockedUser: NewBlockedUserClient(cfg),
		Connection:  NewConnectionClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
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
		ctx:         ctx,
		config:      cfg,
		BlockedUser: NewBlockedUserClient(cfg),
		Connection:  NewConnectionClient(cfg),
		User:        NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		BlockedUser.
//		Query().
//		Count(ctx)
//
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
	c.BlockedUser.Use(hooks...)
	c.Connection.Use(hooks...)
	c.User.Use(hooks...)
}

// BlockedUserClient is a client for the BlockedUser schema.
type BlockedUserClient struct {
	config
}

// NewBlockedUserClient returns a client for the BlockedUser from the given config.
func NewBlockedUserClient(c config) *BlockedUserClient {
	return &BlockedUserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `blockeduser.Hooks(f(g(h())))`.
func (c *BlockedUserClient) Use(hooks ...Hook) {
	c.hooks.BlockedUser = append(c.hooks.BlockedUser, hooks...)
}

// Create returns a builder for creating a BlockedUser entity.
func (c *BlockedUserClient) Create() *BlockedUserCreate {
	mutation := newBlockedUserMutation(c.config, OpCreate)
	return &BlockedUserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of BlockedUser entities.
func (c *BlockedUserClient) CreateBulk(builders ...*BlockedUserCreate) *BlockedUserCreateBulk {
	return &BlockedUserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for BlockedUser.
func (c *BlockedUserClient) Update() *BlockedUserUpdate {
	mutation := newBlockedUserMutation(c.config, OpUpdate)
	return &BlockedUserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *BlockedUserClient) UpdateOne(bu *BlockedUser) *BlockedUserUpdateOne {
	mutation := newBlockedUserMutation(c.config, OpUpdateOne, withBlockedUser(bu))
	return &BlockedUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *BlockedUserClient) UpdateOneID(id int) *BlockedUserUpdateOne {
	mutation := newBlockedUserMutation(c.config, OpUpdateOne, withBlockedUserID(id))
	return &BlockedUserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for BlockedUser.
func (c *BlockedUserClient) Delete() *BlockedUserDelete {
	mutation := newBlockedUserMutation(c.config, OpDelete)
	return &BlockedUserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *BlockedUserClient) DeleteOne(bu *BlockedUser) *BlockedUserDeleteOne {
	return c.DeleteOneID(bu.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *BlockedUserClient) DeleteOneID(id int) *BlockedUserDeleteOne {
	builder := c.Delete().Where(blockeduser.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &BlockedUserDeleteOne{builder}
}

// Query returns a query builder for BlockedUser.
func (c *BlockedUserClient) Query() *BlockedUserQuery {
	return &BlockedUserQuery{
		config: c.config,
	}
}

// Get returns a BlockedUser entity by its id.
func (c *BlockedUserClient) Get(ctx context.Context, id int) (*BlockedUser, error) {
	return c.Query().Where(blockeduser.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *BlockedUserClient) GetX(ctx context.Context, id int) *BlockedUser {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryBlockedBy queries the blocked_by edge of a BlockedUser.
func (c *BlockedUserClient) QueryBlockedBy(bu *BlockedUser) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := bu.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(blockeduser.Table, blockeduser.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, blockeduser.BlockedByTable, blockeduser.BlockedByColumn),
		)
		fromV = sqlgraph.Neighbors(bu.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *BlockedUserClient) Hooks() []Hook {
	return c.hooks.BlockedUser
}

// ConnectionClient is a client for the Connection schema.
type ConnectionClient struct {
	config
}

// NewConnectionClient returns a client for the Connection from the given config.
func NewConnectionClient(c config) *ConnectionClient {
	return &ConnectionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `connection.Hooks(f(g(h())))`.
func (c *ConnectionClient) Use(hooks ...Hook) {
	c.hooks.Connection = append(c.hooks.Connection, hooks...)
}

// Create returns a builder for creating a Connection entity.
func (c *ConnectionClient) Create() *ConnectionCreate {
	mutation := newConnectionMutation(c.config, OpCreate)
	return &ConnectionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Connection entities.
func (c *ConnectionClient) CreateBulk(builders ...*ConnectionCreate) *ConnectionCreateBulk {
	return &ConnectionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Connection.
func (c *ConnectionClient) Update() *ConnectionUpdate {
	mutation := newConnectionMutation(c.config, OpUpdate)
	return &ConnectionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConnectionClient) UpdateOne(co *Connection) *ConnectionUpdateOne {
	mutation := newConnectionMutation(c.config, OpUpdateOne, withConnection(co))
	return &ConnectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConnectionClient) UpdateOneID(id int) *ConnectionUpdateOne {
	mutation := newConnectionMutation(c.config, OpUpdateOne, withConnectionID(id))
	return &ConnectionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Connection.
func (c *ConnectionClient) Delete() *ConnectionDelete {
	mutation := newConnectionMutation(c.config, OpDelete)
	return &ConnectionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ConnectionClient) DeleteOne(co *Connection) *ConnectionDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *ConnectionClient) DeleteOneID(id int) *ConnectionDeleteOne {
	builder := c.Delete().Where(connection.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConnectionDeleteOne{builder}
}

// Query returns a query builder for Connection.
func (c *ConnectionClient) Query() *ConnectionQuery {
	return &ConnectionQuery{
		config: c.config,
	}
}

// Get returns a Connection entity by its id.
func (c *ConnectionClient) Get(ctx context.Context, id int) (*Connection, error) {
	return c.Query().Where(connection.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConnectionClient) GetX(ctx context.Context, id int) *Connection {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Connection.
func (c *ConnectionClient) QueryUser(co *Connection) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(connection.Table, connection.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, connection.UserTable, connection.UserColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryConnection queries the connection edge of a Connection.
func (c *ConnectionClient) QueryConnection(co *Connection) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := co.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(connection.Table, connection.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, connection.ConnectionTable, connection.ConnectionColumn),
		)
		fromV = sqlgraph.Neighbors(co.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ConnectionClient) Hooks() []Hook {
	return c.hooks.Connection
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

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
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

// DeleteOne returns a builder for deleting the given entity by its id.
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

// QueryConnection queries the connection edge of a User.
func (c *UserClient) QueryConnection(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, user.ConnectionTable, user.ConnectionPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryBlock queries the block edge of a User.
func (c *UserClient) QueryBlock(u *User) *BlockedUserQuery {
	query := &BlockedUserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(blockeduser.Table, blockeduser.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.BlockTable, user.BlockColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryConnections queries the connections edge of a User.
func (c *UserClient) QueryConnections(u *User) *ConnectionQuery {
	query := &ConnectionQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := u.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(connection.Table, connection.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, user.ConnectionsTable, user.ConnectionsColumn),
		)
		fromV = sqlgraph.Neighbors(u.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
