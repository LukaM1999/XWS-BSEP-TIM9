// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"dislinkt/job_offer_service/ent/joboffer"
	"dislinkt/job_offer_service/ent/predicate"
	"dislinkt/job_offer_service/ent/skill"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// JobOfferQuery is the builder for querying JobOffer entities.
type JobOfferQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.JobOffer
	// eager-loading edges.
	withRequires *SkillQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the JobOfferQuery builder.
func (joq *JobOfferQuery) Where(ps ...predicate.JobOffer) *JobOfferQuery {
	joq.predicates = append(joq.predicates, ps...)
	return joq
}

// Limit adds a limit step to the query.
func (joq *JobOfferQuery) Limit(limit int) *JobOfferQuery {
	joq.limit = &limit
	return joq
}

// Offset adds an offset step to the query.
func (joq *JobOfferQuery) Offset(offset int) *JobOfferQuery {
	joq.offset = &offset
	return joq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (joq *JobOfferQuery) Unique(unique bool) *JobOfferQuery {
	joq.unique = &unique
	return joq
}

// Order adds an order step to the query.
func (joq *JobOfferQuery) Order(o ...OrderFunc) *JobOfferQuery {
	joq.order = append(joq.order, o...)
	return joq
}

// QueryRequires chains the current query on the "requires" edge.
func (joq *JobOfferQuery) QueryRequires() *SkillQuery {
	query := &SkillQuery{config: joq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := joq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := joq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(joboffer.Table, joboffer.FieldID, selector),
			sqlgraph.To(skill.Table, skill.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, joboffer.RequiresTable, joboffer.RequiresPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(joq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first JobOffer entity from the query.
// Returns a *NotFoundError when no JobOffer was found.
func (joq *JobOfferQuery) First(ctx context.Context) (*JobOffer, error) {
	nodes, err := joq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{joboffer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (joq *JobOfferQuery) FirstX(ctx context.Context) *JobOffer {
	node, err := joq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first JobOffer ID from the query.
// Returns a *NotFoundError when no JobOffer ID was found.
func (joq *JobOfferQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = joq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{joboffer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (joq *JobOfferQuery) FirstIDX(ctx context.Context) int {
	id, err := joq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single JobOffer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one JobOffer entity is found.
// Returns a *NotFoundError when no JobOffer entities are found.
func (joq *JobOfferQuery) Only(ctx context.Context) (*JobOffer, error) {
	nodes, err := joq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{joboffer.Label}
	default:
		return nil, &NotSingularError{joboffer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (joq *JobOfferQuery) OnlyX(ctx context.Context) *JobOffer {
	node, err := joq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only JobOffer ID in the query.
// Returns a *NotSingularError when more than one JobOffer ID is found.
// Returns a *NotFoundError when no entities are found.
func (joq *JobOfferQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = joq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = &NotSingularError{joboffer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (joq *JobOfferQuery) OnlyIDX(ctx context.Context) int {
	id, err := joq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of JobOffers.
func (joq *JobOfferQuery) All(ctx context.Context) ([]*JobOffer, error) {
	if err := joq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return joq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (joq *JobOfferQuery) AllX(ctx context.Context) []*JobOffer {
	nodes, err := joq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of JobOffer IDs.
func (joq *JobOfferQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := joq.Select(joboffer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (joq *JobOfferQuery) IDsX(ctx context.Context) []int {
	ids, err := joq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (joq *JobOfferQuery) Count(ctx context.Context) (int, error) {
	if err := joq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return joq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (joq *JobOfferQuery) CountX(ctx context.Context) int {
	count, err := joq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (joq *JobOfferQuery) Exist(ctx context.Context) (bool, error) {
	if err := joq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return joq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (joq *JobOfferQuery) ExistX(ctx context.Context) bool {
	exist, err := joq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the JobOfferQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (joq *JobOfferQuery) Clone() *JobOfferQuery {
	if joq == nil {
		return nil
	}
	return &JobOfferQuery{
		config:       joq.config,
		limit:        joq.limit,
		offset:       joq.offset,
		order:        append([]OrderFunc{}, joq.order...),
		predicates:   append([]predicate.JobOffer{}, joq.predicates...),
		withRequires: joq.withRequires.Clone(),
		// clone intermediate query.
		sql:    joq.sql.Clone(),
		path:   joq.path,
		unique: joq.unique,
	}
}

// WithRequires tells the query-builder to eager-load the nodes that are connected to
// the "requires" edge. The optional arguments are used to configure the query builder of the edge.
func (joq *JobOfferQuery) WithRequires(opts ...func(*SkillQuery)) *JobOfferQuery {
	query := &SkillQuery{config: joq.config}
	for _, opt := range opts {
		opt(query)
	}
	joq.withRequires = query
	return joq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProfileID string `json:"profile_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.JobOffer.Query().
//		GroupBy(joboffer.FieldProfileID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (joq *JobOfferQuery) GroupBy(field string, fields ...string) *JobOfferGroupBy {
	group := &JobOfferGroupBy{config: joq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := joq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return joq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProfileID string `json:"profile_id,omitempty"`
//	}
//
//	client.JobOffer.Query().
//		Select(joboffer.FieldProfileID).
//		Scan(ctx, &v)
func (joq *JobOfferQuery) Select(fields ...string) *JobOfferSelect {
	joq.fields = append(joq.fields, fields...)
	return &JobOfferSelect{JobOfferQuery: joq}
}

func (joq *JobOfferQuery) prepareQuery(ctx context.Context) error {
	for _, f := range joq.fields {
		if !joboffer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if joq.path != nil {
		prev, err := joq.path(ctx)
		if err != nil {
			return err
		}
		joq.sql = prev
	}
	return nil
}

func (joq *JobOfferQuery) sqlAll(ctx context.Context) ([]*JobOffer, error) {
	var (
		nodes       = []*JobOffer{}
		_spec       = joq.querySpec()
		loadedTypes = [1]bool{
			joq.withRequires != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &JobOffer{config: joq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("ent: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, joq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := joq.withRequires; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		ids := make(map[int]*JobOffer, len(nodes))
		for _, node := range nodes {
			ids[node.ID] = node
			fks = append(fks, node.ID)
			node.Edges.Requires = []*Skill{}
		}
		var (
			edgeids []int
			edges   = make(map[int][]*JobOffer)
		)
		_spec := &sqlgraph.EdgeQuerySpec{
			Edge: &sqlgraph.EdgeSpec{
				Inverse: false,
				Table:   joboffer.RequiresTable,
				Columns: joboffer.RequiresPrimaryKey,
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues(joboffer.RequiresPrimaryKey[0], fks...))
			},
			ScanValues: func() [2]interface{} {
				return [2]interface{}{new(sql.NullInt64), new(sql.NullInt64)}
			},
			Assign: func(out, in interface{}) error {
				eout, ok := out.(*sql.NullInt64)
				if !ok || eout == nil {
					return fmt.Errorf("unexpected id value for edge-out")
				}
				ein, ok := in.(*sql.NullInt64)
				if !ok || ein == nil {
					return fmt.Errorf("unexpected id value for edge-in")
				}
				outValue := int(eout.Int64)
				inValue := int(ein.Int64)
				node, ok := ids[outValue]
				if !ok {
					return fmt.Errorf("unexpected node id in edges: %v", outValue)
				}
				if _, ok := edges[inValue]; !ok {
					edgeids = append(edgeids, inValue)
				}
				edges[inValue] = append(edges[inValue], node)
				return nil
			},
		}
		if err := sqlgraph.QueryEdges(ctx, joq.driver, _spec); err != nil {
			return nil, fmt.Errorf(`query edges "requires": %w`, err)
		}
		query.Where(skill.IDIn(edgeids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := edges[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected "requires" node returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Requires = append(nodes[i].Edges.Requires, n)
			}
		}
	}

	return nodes, nil
}

func (joq *JobOfferQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := joq.querySpec()
	_spec.Node.Columns = joq.fields
	if len(joq.fields) > 0 {
		_spec.Unique = joq.unique != nil && *joq.unique
	}
	return sqlgraph.CountNodes(ctx, joq.driver, _spec)
}

func (joq *JobOfferQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := joq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (joq *JobOfferQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   joboffer.Table,
			Columns: joboffer.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: joboffer.FieldID,
			},
		},
		From:   joq.sql,
		Unique: true,
	}
	if unique := joq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := joq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, joboffer.FieldID)
		for i := range fields {
			if fields[i] != joboffer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := joq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := joq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := joq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := joq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (joq *JobOfferQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(joq.driver.Dialect())
	t1 := builder.Table(joboffer.Table)
	columns := joq.fields
	if len(columns) == 0 {
		columns = joboffer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if joq.sql != nil {
		selector = joq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if joq.unique != nil && *joq.unique {
		selector.Distinct()
	}
	for _, p := range joq.predicates {
		p(selector)
	}
	for _, p := range joq.order {
		p(selector)
	}
	if offset := joq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := joq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// JobOfferGroupBy is the group-by builder for JobOffer entities.
type JobOfferGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (jogb *JobOfferGroupBy) Aggregate(fns ...AggregateFunc) *JobOfferGroupBy {
	jogb.fns = append(jogb.fns, fns...)
	return jogb
}

// Scan applies the group-by query and scans the result into the given value.
func (jogb *JobOfferGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := jogb.path(ctx)
	if err != nil {
		return err
	}
	jogb.sql = query
	return jogb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jogb *JobOfferGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := jogb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(jogb.fields) > 1 {
		return nil, errors.New("ent: JobOfferGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := jogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jogb *JobOfferGroupBy) StringsX(ctx context.Context) []string {
	v, err := jogb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jogb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jogb *JobOfferGroupBy) StringX(ctx context.Context) string {
	v, err := jogb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(jogb.fields) > 1 {
		return nil, errors.New("ent: JobOfferGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := jogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jogb *JobOfferGroupBy) IntsX(ctx context.Context) []int {
	v, err := jogb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jogb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jogb *JobOfferGroupBy) IntX(ctx context.Context) int {
	v, err := jogb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(jogb.fields) > 1 {
		return nil, errors.New("ent: JobOfferGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := jogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jogb *JobOfferGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := jogb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jogb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jogb *JobOfferGroupBy) Float64X(ctx context.Context) float64 {
	v, err := jogb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(jogb.fields) > 1 {
		return nil, errors.New("ent: JobOfferGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := jogb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jogb *JobOfferGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := jogb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (jogb *JobOfferGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jogb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jogb *JobOfferGroupBy) BoolX(ctx context.Context) bool {
	v, err := jogb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jogb *JobOfferGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range jogb.fields {
		if !joboffer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := jogb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := jogb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (jogb *JobOfferGroupBy) sqlQuery() *sql.Selector {
	selector := jogb.sql.Select()
	aggregation := make([]string, 0, len(jogb.fns))
	for _, fn := range jogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(jogb.fields)+len(jogb.fns))
		for _, f := range jogb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(jogb.fields...)...)
}

// JobOfferSelect is the builder for selecting fields of JobOffer entities.
type JobOfferSelect struct {
	*JobOfferQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (jos *JobOfferSelect) Scan(ctx context.Context, v interface{}) error {
	if err := jos.prepareQuery(ctx); err != nil {
		return err
	}
	jos.sql = jos.JobOfferQuery.sqlQuery(ctx)
	return jos.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (jos *JobOfferSelect) ScanX(ctx context.Context, v interface{}) {
	if err := jos.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) Strings(ctx context.Context) ([]string, error) {
	if len(jos.fields) > 1 {
		return nil, errors.New("ent: JobOfferSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := jos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (jos *JobOfferSelect) StringsX(ctx context.Context) []string {
	v, err := jos.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = jos.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (jos *JobOfferSelect) StringX(ctx context.Context) string {
	v, err := jos.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) Ints(ctx context.Context) ([]int, error) {
	if len(jos.fields) > 1 {
		return nil, errors.New("ent: JobOfferSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := jos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (jos *JobOfferSelect) IntsX(ctx context.Context) []int {
	v, err := jos.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = jos.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (jos *JobOfferSelect) IntX(ctx context.Context) int {
	v, err := jos.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(jos.fields) > 1 {
		return nil, errors.New("ent: JobOfferSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := jos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (jos *JobOfferSelect) Float64sX(ctx context.Context) []float64 {
	v, err := jos.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = jos.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (jos *JobOfferSelect) Float64X(ctx context.Context) float64 {
	v, err := jos.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(jos.fields) > 1 {
		return nil, errors.New("ent: JobOfferSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := jos.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (jos *JobOfferSelect) BoolsX(ctx context.Context) []bool {
	v, err := jos.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (jos *JobOfferSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = jos.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{joboffer.Label}
	default:
		err = fmt.Errorf("ent: JobOfferSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (jos *JobOfferSelect) BoolX(ctx context.Context) bool {
	v, err := jos.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (jos *JobOfferSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := jos.sql.Query()
	if err := jos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
