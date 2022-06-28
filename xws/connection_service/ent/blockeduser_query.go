// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"dislinkt/connection_service/ent/blockeduser"
	"dislinkt/connection_service/ent/predicate"
	"dislinkt/connection_service/ent/user"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BlockedUserQuery is the builder for querying BlockedUser entities.
type BlockedUserQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.BlockedUser
	// eager-loading edges.
	withBlockedBy *UserQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the BlockedUserQuery builder.
func (buq *BlockedUserQuery) Where(ps ...predicate.BlockedUser) *BlockedUserQuery {
	buq.predicates = append(buq.predicates, ps...)
	return buq
}

// Limit adds a limit step to the query.
func (buq *BlockedUserQuery) Limit(limit int) *BlockedUserQuery {
	buq.limit = &limit
	return buq
}

// Offset adds an offset step to the query.
func (buq *BlockedUserQuery) Offset(offset int) *BlockedUserQuery {
	buq.offset = &offset
	return buq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (buq *BlockedUserQuery) Unique(unique bool) *BlockedUserQuery {
	buq.unique = &unique
	return buq
}

// Order adds an order step to the query.
func (buq *BlockedUserQuery) Order(o ...OrderFunc) *BlockedUserQuery {
	buq.order = append(buq.order, o...)
	return buq
}

// QueryBlockedBy chains the current query on the "blocked_by" edge.
func (buq *BlockedUserQuery) QueryBlockedBy() *UserQuery {
	query := &UserQuery{config: buq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := buq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := buq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(blockeduser.Table, blockeduser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, blockeduser.BlockedByTable, blockeduser.BlockedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(buq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first BlockedUser entity from the query.
// Returns a *NotFoundError when no BlockedUser was found.
func (buq *BlockedUserQuery) First(ctx context.Context) (*BlockedUser, error) {
	nodes, err := buq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{blockeduser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (buq *BlockedUserQuery) FirstX(ctx context.Context) *BlockedUser {
	node, err := buq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first BlockedUser ID from the query.
// Returns a *NotFoundError when no BlockedUser ID was found.
func (buq *BlockedUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = buq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{blockeduser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (buq *BlockedUserQuery) FirstIDX(ctx context.Context) int {
	id, err := buq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single BlockedUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one BlockedUser entity is found.
// Returns a *NotFoundError when no BlockedUser entities are found.
func (buq *BlockedUserQuery) Only(ctx context.Context) (*BlockedUser, error) {
	nodes, err := buq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{blockeduser.Label}
	default:
		return nil, &NotSingularError{blockeduser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (buq *BlockedUserQuery) OnlyX(ctx context.Context) *BlockedUser {
	node, err := buq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only BlockedUser ID in the query.
// Returns a *NotSingularError when more than one BlockedUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (buq *BlockedUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = buq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{blockeduser.Label}
	default:
		err = &NotSingularError{blockeduser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (buq *BlockedUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := buq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of BlockedUsers.
func (buq *BlockedUserQuery) All(ctx context.Context) ([]*BlockedUser, error) {
	if err := buq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return buq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (buq *BlockedUserQuery) AllX(ctx context.Context) []*BlockedUser {
	nodes, err := buq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of BlockedUser IDs.
func (buq *BlockedUserQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := buq.Select(blockeduser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (buq *BlockedUserQuery) IDsX(ctx context.Context) []int {
	ids, err := buq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (buq *BlockedUserQuery) Count(ctx context.Context) (int, error) {
	if err := buq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return buq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (buq *BlockedUserQuery) CountX(ctx context.Context) int {
	count, err := buq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (buq *BlockedUserQuery) Exist(ctx context.Context) (bool, error) {
	if err := buq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return buq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (buq *BlockedUserQuery) ExistX(ctx context.Context) bool {
	exist, err := buq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the BlockedUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (buq *BlockedUserQuery) Clone() *BlockedUserQuery {
	if buq == nil {
		return nil
	}
	return &BlockedUserQuery{
		config:        buq.config,
		limit:         buq.limit,
		offset:        buq.offset,
		order:         append([]OrderFunc{}, buq.order...),
		predicates:    append([]predicate.BlockedUser{}, buq.predicates...),
		withBlockedBy: buq.withBlockedBy.Clone(),
		// clone intermediate query.
		sql:    buq.sql.Clone(),
		path:   buq.path,
		unique: buq.unique,
	}
}

// WithBlockedBy tells the query-builder to eager-load the nodes that are connected to
// the "blocked_by" edge. The optional arguments are used to configure the query builder of the edge.
func (buq *BlockedUserQuery) WithBlockedBy(opts ...func(*UserQuery)) *BlockedUserQuery {
	query := &UserQuery{config: buq.config}
	for _, opt := range opts {
		opt(query)
	}
	buq.withBlockedBy = query
	return buq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.BlockedUser.Query().
//		GroupBy(blockeduser.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (buq *BlockedUserQuery) GroupBy(field string, fields ...string) *BlockedUserGroupBy {
	grbuild := &BlockedUserGroupBy{config: buq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := buq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return buq.sqlQuery(ctx), nil
	}
	grbuild.label = blockeduser.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.BlockedUser.Query().
//		Select(blockeduser.FieldCreatedAt).
//		Scan(ctx, &v)
//
func (buq *BlockedUserQuery) Select(fields ...string) *BlockedUserSelect {
	buq.fields = append(buq.fields, fields...)
	selbuild := &BlockedUserSelect{BlockedUserQuery: buq}
	selbuild.label = blockeduser.Label
	selbuild.flds, selbuild.scan = &buq.fields, selbuild.Scan
	return selbuild
}

func (buq *BlockedUserQuery) prepareQuery(ctx context.Context) error {
	for _, f := range buq.fields {
		if !blockeduser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if buq.path != nil {
		prev, err := buq.path(ctx)
		if err != nil {
			return err
		}
		buq.sql = prev
	}
	return nil
}

func (buq *BlockedUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*BlockedUser, error) {
	var (
		nodes       = []*BlockedUser{}
		withFKs     = buq.withFKs
		_spec       = buq.querySpec()
		loadedTypes = [1]bool{
			buq.withBlockedBy != nil,
		}
	)
	if buq.withBlockedBy != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, blockeduser.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*BlockedUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &BlockedUser{config: buq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, buq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := buq.withBlockedBy; query != nil {
		ids := make([]int, 0, len(nodes))
		nodeids := make(map[int][]*BlockedUser)
		for i := range nodes {
			if nodes[i].user_block == nil {
				continue
			}
			fk := *nodes[i].user_block
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(user.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "user_block" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.BlockedBy = n
			}
		}
	}

	return nodes, nil
}

func (buq *BlockedUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := buq.querySpec()
	_spec.Node.Columns = buq.fields
	if len(buq.fields) > 0 {
		_spec.Unique = buq.unique != nil && *buq.unique
	}
	return sqlgraph.CountNodes(ctx, buq.driver, _spec)
}

func (buq *BlockedUserQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := buq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (buq *BlockedUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   blockeduser.Table,
			Columns: blockeduser.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: blockeduser.FieldID,
			},
		},
		From:   buq.sql,
		Unique: true,
	}
	if unique := buq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := buq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, blockeduser.FieldID)
		for i := range fields {
			if fields[i] != blockeduser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := buq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := buq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := buq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := buq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (buq *BlockedUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(buq.driver.Dialect())
	t1 := builder.Table(blockeduser.Table)
	columns := buq.fields
	if len(columns) == 0 {
		columns = blockeduser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if buq.sql != nil {
		selector = buq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if buq.unique != nil && *buq.unique {
		selector.Distinct()
	}
	for _, p := range buq.predicates {
		p(selector)
	}
	for _, p := range buq.order {
		p(selector)
	}
	if offset := buq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := buq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// BlockedUserGroupBy is the group-by builder for BlockedUser entities.
type BlockedUserGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (bugb *BlockedUserGroupBy) Aggregate(fns ...AggregateFunc) *BlockedUserGroupBy {
	bugb.fns = append(bugb.fns, fns...)
	return bugb
}

// Scan applies the group-by query and scans the result into the given value.
func (bugb *BlockedUserGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := bugb.path(ctx)
	if err != nil {
		return err
	}
	bugb.sql = query
	return bugb.sqlScan(ctx, v)
}

func (bugb *BlockedUserGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range bugb.fields {
		if !blockeduser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := bugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := bugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (bugb *BlockedUserGroupBy) sqlQuery() *sql.Selector {
	selector := bugb.sql.Select()
	aggregation := make([]string, 0, len(bugb.fns))
	for _, fn := range bugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(bugb.fields)+len(bugb.fns))
		for _, f := range bugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(bugb.fields...)...)
}

// BlockedUserSelect is the builder for selecting fields of BlockedUser entities.
type BlockedUserSelect struct {
	*BlockedUserQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (bus *BlockedUserSelect) Scan(ctx context.Context, v interface{}) error {
	if err := bus.prepareQuery(ctx); err != nil {
		return err
	}
	bus.sql = bus.BlockedUserQuery.sqlQuery(ctx)
	return bus.sqlScan(ctx, v)
}

func (bus *BlockedUserSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := bus.sql.Query()
	if err := bus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
