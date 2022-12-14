// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/velcro-xiv/velcro/ent/logevent"
	"github.com/velcro-xiv/velcro/ent/predicate"
)

// LogEventQuery is the builder for querying LogEvent entities.
type LogEventQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.LogEvent
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LogEventQuery builder.
func (leq *LogEventQuery) Where(ps ...predicate.LogEvent) *LogEventQuery {
	leq.predicates = append(leq.predicates, ps...)
	return leq
}

// Limit adds a limit step to the query.
func (leq *LogEventQuery) Limit(limit int) *LogEventQuery {
	leq.limit = &limit
	return leq
}

// Offset adds an offset step to the query.
func (leq *LogEventQuery) Offset(offset int) *LogEventQuery {
	leq.offset = &offset
	return leq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (leq *LogEventQuery) Unique(unique bool) *LogEventQuery {
	leq.unique = &unique
	return leq
}

// Order adds an order step to the query.
func (leq *LogEventQuery) Order(o ...OrderFunc) *LogEventQuery {
	leq.order = append(leq.order, o...)
	return leq
}

// First returns the first LogEvent entity from the query.
// Returns a *NotFoundError when no LogEvent was found.
func (leq *LogEventQuery) First(ctx context.Context) (*LogEvent, error) {
	nodes, err := leq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{logevent.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (leq *LogEventQuery) FirstX(ctx context.Context) *LogEvent {
	node, err := leq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first LogEvent ID from the query.
// Returns a *NotFoundError when no LogEvent ID was found.
func (leq *LogEventQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = leq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{logevent.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (leq *LogEventQuery) FirstIDX(ctx context.Context) int {
	id, err := leq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single LogEvent entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one LogEvent entity is found.
// Returns a *NotFoundError when no LogEvent entities are found.
func (leq *LogEventQuery) Only(ctx context.Context) (*LogEvent, error) {
	nodes, err := leq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{logevent.Label}
	default:
		return nil, &NotSingularError{logevent.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (leq *LogEventQuery) OnlyX(ctx context.Context) *LogEvent {
	node, err := leq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only LogEvent ID in the query.
// Returns a *NotSingularError when more than one LogEvent ID is found.
// Returns a *NotFoundError when no entities are found.
func (leq *LogEventQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = leq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{logevent.Label}
	default:
		err = &NotSingularError{logevent.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (leq *LogEventQuery) OnlyIDX(ctx context.Context) int {
	id, err := leq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of LogEvents.
func (leq *LogEventQuery) All(ctx context.Context) ([]*LogEvent, error) {
	if err := leq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return leq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (leq *LogEventQuery) AllX(ctx context.Context) []*LogEvent {
	nodes, err := leq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of LogEvent IDs.
func (leq *LogEventQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := leq.Select(logevent.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (leq *LogEventQuery) IDsX(ctx context.Context) []int {
	ids, err := leq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (leq *LogEventQuery) Count(ctx context.Context) (int, error) {
	if err := leq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return leq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (leq *LogEventQuery) CountX(ctx context.Context) int {
	count, err := leq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (leq *LogEventQuery) Exist(ctx context.Context) (bool, error) {
	if err := leq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return leq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (leq *LogEventQuery) ExistX(ctx context.Context) bool {
	exist, err := leq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LogEventQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (leq *LogEventQuery) Clone() *LogEventQuery {
	if leq == nil {
		return nil
	}
	return &LogEventQuery{
		config:     leq.config,
		limit:      leq.limit,
		offset:     leq.offset,
		order:      append([]OrderFunc{}, leq.order...),
		predicates: append([]predicate.LogEvent{}, leq.predicates...),
		// clone intermediate query.
		sql:    leq.sql.Clone(),
		path:   leq.path,
		unique: leq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.LogEvent.Query().
//		GroupBy(logevent.FieldTimestamp).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (leq *LogEventQuery) GroupBy(field string, fields ...string) *LogEventGroupBy {
	grbuild := &LogEventGroupBy{config: leq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := leq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return leq.sqlQuery(ctx), nil
	}
	grbuild.label = logevent.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Timestamp time.Time `json:"timestamp,omitempty"`
//	}
//
//	client.LogEvent.Query().
//		Select(logevent.FieldTimestamp).
//		Scan(ctx, &v)
func (leq *LogEventQuery) Select(fields ...string) *LogEventSelect {
	leq.fields = append(leq.fields, fields...)
	selbuild := &LogEventSelect{LogEventQuery: leq}
	selbuild.label = logevent.Label
	selbuild.flds, selbuild.scan = &leq.fields, selbuild.Scan
	return selbuild
}

func (leq *LogEventQuery) prepareQuery(ctx context.Context) error {
	for _, f := range leq.fields {
		if !logevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if leq.path != nil {
		prev, err := leq.path(ctx)
		if err != nil {
			return err
		}
		leq.sql = prev
	}
	return nil
}

func (leq *LogEventQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*LogEvent, error) {
	var (
		nodes = []*LogEvent{}
		_spec = leq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*LogEvent).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &LogEvent{config: leq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, leq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (leq *LogEventQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := leq.querySpec()
	_spec.Node.Columns = leq.fields
	if len(leq.fields) > 0 {
		_spec.Unique = leq.unique != nil && *leq.unique
	}
	return sqlgraph.CountNodes(ctx, leq.driver, _spec)
}

func (leq *LogEventQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := leq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (leq *LogEventQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   logevent.Table,
			Columns: logevent.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: logevent.FieldID,
			},
		},
		From:   leq.sql,
		Unique: true,
	}
	if unique := leq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := leq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, logevent.FieldID)
		for i := range fields {
			if fields[i] != logevent.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := leq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := leq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := leq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := leq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (leq *LogEventQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(leq.driver.Dialect())
	t1 := builder.Table(logevent.Table)
	columns := leq.fields
	if len(columns) == 0 {
		columns = logevent.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if leq.sql != nil {
		selector = leq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if leq.unique != nil && *leq.unique {
		selector.Distinct()
	}
	for _, p := range leq.predicates {
		p(selector)
	}
	for _, p := range leq.order {
		p(selector)
	}
	if offset := leq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := leq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// LogEventGroupBy is the group-by builder for LogEvent entities.
type LogEventGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (legb *LogEventGroupBy) Aggregate(fns ...AggregateFunc) *LogEventGroupBy {
	legb.fns = append(legb.fns, fns...)
	return legb
}

// Scan applies the group-by query and scans the result into the given value.
func (legb *LogEventGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := legb.path(ctx)
	if err != nil {
		return err
	}
	legb.sql = query
	return legb.sqlScan(ctx, v)
}

func (legb *LogEventGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range legb.fields {
		if !logevent.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := legb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := legb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (legb *LogEventGroupBy) sqlQuery() *sql.Selector {
	selector := legb.sql.Select()
	aggregation := make([]string, 0, len(legb.fns))
	for _, fn := range legb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(legb.fields)+len(legb.fns))
		for _, f := range legb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(legb.fields...)...)
}

// LogEventSelect is the builder for selecting fields of LogEvent entities.
type LogEventSelect struct {
	*LogEventQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (les *LogEventSelect) Scan(ctx context.Context, v interface{}) error {
	if err := les.prepareQuery(ctx); err != nil {
		return err
	}
	les.sql = les.LogEventQuery.sqlQuery(ctx)
	return les.sqlScan(ctx, v)
}

func (les *LogEventSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := les.sql.Query()
	if err := les.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
