// Code generated by ent, DO NOT EDIT.

package core

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"freefrom.space/nobot/core/internal"
	"freefrom.space/nobot/core/predicate"
	"freefrom.space/nobot/core/twitternotesfetchhistory"
)

// TwitterNotesFetchHistoryQuery is the builder for querying TwitterNotesFetchHistory entities.
type TwitterNotesFetchHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []twitternotesfetchhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.TwitterNotesFetchHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TwitterNotesFetchHistoryQuery builder.
func (tnfhq *TwitterNotesFetchHistoryQuery) Where(ps ...predicate.TwitterNotesFetchHistory) *TwitterNotesFetchHistoryQuery {
	tnfhq.predicates = append(tnfhq.predicates, ps...)
	return tnfhq
}

// Limit the number of records to be returned by this query.
func (tnfhq *TwitterNotesFetchHistoryQuery) Limit(limit int) *TwitterNotesFetchHistoryQuery {
	tnfhq.ctx.Limit = &limit
	return tnfhq
}

// Offset to start from.
func (tnfhq *TwitterNotesFetchHistoryQuery) Offset(offset int) *TwitterNotesFetchHistoryQuery {
	tnfhq.ctx.Offset = &offset
	return tnfhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tnfhq *TwitterNotesFetchHistoryQuery) Unique(unique bool) *TwitterNotesFetchHistoryQuery {
	tnfhq.ctx.Unique = &unique
	return tnfhq
}

// Order specifies how the records should be ordered.
func (tnfhq *TwitterNotesFetchHistoryQuery) Order(o ...twitternotesfetchhistory.OrderOption) *TwitterNotesFetchHistoryQuery {
	tnfhq.order = append(tnfhq.order, o...)
	return tnfhq
}

// First returns the first TwitterNotesFetchHistory entity from the query.
// Returns a *NotFoundError when no TwitterNotesFetchHistory was found.
func (tnfhq *TwitterNotesFetchHistoryQuery) First(ctx context.Context) (*TwitterNotesFetchHistory, error) {
	nodes, err := tnfhq.Limit(1).All(setContextOp(ctx, tnfhq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{twitternotesfetchhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) FirstX(ctx context.Context) *TwitterNotesFetchHistory {
	node, err := tnfhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TwitterNotesFetchHistory ID from the query.
// Returns a *NotFoundError when no TwitterNotesFetchHistory ID was found.
func (tnfhq *TwitterNotesFetchHistoryQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tnfhq.Limit(1).IDs(setContextOp(ctx, tnfhq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{twitternotesfetchhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) FirstIDX(ctx context.Context) int64 {
	id, err := tnfhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TwitterNotesFetchHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TwitterNotesFetchHistory entity is found.
// Returns a *NotFoundError when no TwitterNotesFetchHistory entities are found.
func (tnfhq *TwitterNotesFetchHistoryQuery) Only(ctx context.Context) (*TwitterNotesFetchHistory, error) {
	nodes, err := tnfhq.Limit(2).All(setContextOp(ctx, tnfhq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{twitternotesfetchhistory.Label}
	default:
		return nil, &NotSingularError{twitternotesfetchhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) OnlyX(ctx context.Context) *TwitterNotesFetchHistory {
	node, err := tnfhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TwitterNotesFetchHistory ID in the query.
// Returns a *NotSingularError when more than one TwitterNotesFetchHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (tnfhq *TwitterNotesFetchHistoryQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = tnfhq.Limit(2).IDs(setContextOp(ctx, tnfhq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{twitternotesfetchhistory.Label}
	default:
		err = &NotSingularError{twitternotesfetchhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := tnfhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TwitterNotesFetchHistories.
func (tnfhq *TwitterNotesFetchHistoryQuery) All(ctx context.Context) ([]*TwitterNotesFetchHistory, error) {
	ctx = setContextOp(ctx, tnfhq.ctx, "All")
	if err := tnfhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TwitterNotesFetchHistory, *TwitterNotesFetchHistoryQuery]()
	return withInterceptors[[]*TwitterNotesFetchHistory](ctx, tnfhq, qr, tnfhq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) AllX(ctx context.Context) []*TwitterNotesFetchHistory {
	nodes, err := tnfhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TwitterNotesFetchHistory IDs.
func (tnfhq *TwitterNotesFetchHistoryQuery) IDs(ctx context.Context) (ids []int64, err error) {
	if tnfhq.ctx.Unique == nil && tnfhq.path != nil {
		tnfhq.Unique(true)
	}
	ctx = setContextOp(ctx, tnfhq.ctx, "IDs")
	if err = tnfhq.Select(twitternotesfetchhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) IDsX(ctx context.Context) []int64 {
	ids, err := tnfhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tnfhq *TwitterNotesFetchHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tnfhq.ctx, "Count")
	if err := tnfhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tnfhq, querierCount[*TwitterNotesFetchHistoryQuery](), tnfhq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) CountX(ctx context.Context) int {
	count, err := tnfhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tnfhq *TwitterNotesFetchHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tnfhq.ctx, "Exist")
	switch _, err := tnfhq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("core: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tnfhq *TwitterNotesFetchHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := tnfhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TwitterNotesFetchHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tnfhq *TwitterNotesFetchHistoryQuery) Clone() *TwitterNotesFetchHistoryQuery {
	if tnfhq == nil {
		return nil
	}
	return &TwitterNotesFetchHistoryQuery{
		config:     tnfhq.config,
		ctx:        tnfhq.ctx.Clone(),
		order:      append([]twitternotesfetchhistory.OrderOption{}, tnfhq.order...),
		inters:     append([]Interceptor{}, tnfhq.inters...),
		predicates: append([]predicate.TwitterNotesFetchHistory{}, tnfhq.predicates...),
		// clone intermediate query.
		sql:  tnfhq.sql.Clone(),
		path: tnfhq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TwitterNotesFetchHistory.Query().
//		GroupBy(twitternotesfetchhistory.FieldUserID).
//		Aggregate(core.Count()).
//		Scan(ctx, &v)
func (tnfhq *TwitterNotesFetchHistoryQuery) GroupBy(field string, fields ...string) *TwitterNotesFetchHistoryGroupBy {
	tnfhq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TwitterNotesFetchHistoryGroupBy{build: tnfhq}
	grbuild.flds = &tnfhq.ctx.Fields
	grbuild.label = twitternotesfetchhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID string `json:"user_id,omitempty"`
//	}
//
//	client.TwitterNotesFetchHistory.Query().
//		Select(twitternotesfetchhistory.FieldUserID).
//		Scan(ctx, &v)
func (tnfhq *TwitterNotesFetchHistoryQuery) Select(fields ...string) *TwitterNotesFetchHistorySelect {
	tnfhq.ctx.Fields = append(tnfhq.ctx.Fields, fields...)
	sbuild := &TwitterNotesFetchHistorySelect{TwitterNotesFetchHistoryQuery: tnfhq}
	sbuild.label = twitternotesfetchhistory.Label
	sbuild.flds, sbuild.scan = &tnfhq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TwitterNotesFetchHistorySelect configured with the given aggregations.
func (tnfhq *TwitterNotesFetchHistoryQuery) Aggregate(fns ...AggregateFunc) *TwitterNotesFetchHistorySelect {
	return tnfhq.Select().Aggregate(fns...)
}

func (tnfhq *TwitterNotesFetchHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tnfhq.inters {
		if inter == nil {
			return fmt.Errorf("core: uninitialized interceptor (forgotten import core/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tnfhq); err != nil {
				return err
			}
		}
	}
	for _, f := range tnfhq.ctx.Fields {
		if !twitternotesfetchhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("core: invalid field %q for query", f)}
		}
	}
	if tnfhq.path != nil {
		prev, err := tnfhq.path(ctx)
		if err != nil {
			return err
		}
		tnfhq.sql = prev
	}
	return nil
}

func (tnfhq *TwitterNotesFetchHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TwitterNotesFetchHistory, error) {
	var (
		nodes = []*TwitterNotesFetchHistory{}
		_spec = tnfhq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TwitterNotesFetchHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TwitterNotesFetchHistory{config: tnfhq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tnfhq.schemaConfig.TwitterNotesFetchHistory
	ctx = internal.NewSchemaConfigContext(ctx, tnfhq.schemaConfig)
	if len(tnfhq.modifiers) > 0 {
		_spec.Modifiers = tnfhq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tnfhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tnfhq *TwitterNotesFetchHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tnfhq.querySpec()
	_spec.Node.Schema = tnfhq.schemaConfig.TwitterNotesFetchHistory
	ctx = internal.NewSchemaConfigContext(ctx, tnfhq.schemaConfig)
	if len(tnfhq.modifiers) > 0 {
		_spec.Modifiers = tnfhq.modifiers
	}
	_spec.Node.Columns = tnfhq.ctx.Fields
	if len(tnfhq.ctx.Fields) > 0 {
		_spec.Unique = tnfhq.ctx.Unique != nil && *tnfhq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tnfhq.driver, _spec)
}

func (tnfhq *TwitterNotesFetchHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(twitternotesfetchhistory.Table, twitternotesfetchhistory.Columns, sqlgraph.NewFieldSpec(twitternotesfetchhistory.FieldID, field.TypeInt64))
	_spec.From = tnfhq.sql
	if unique := tnfhq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tnfhq.path != nil {
		_spec.Unique = true
	}
	if fields := tnfhq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twitternotesfetchhistory.FieldID)
		for i := range fields {
			if fields[i] != twitternotesfetchhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tnfhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tnfhq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tnfhq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tnfhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tnfhq *TwitterNotesFetchHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tnfhq.driver.Dialect())
	t1 := builder.Table(twitternotesfetchhistory.Table)
	columns := tnfhq.ctx.Fields
	if len(columns) == 0 {
		columns = twitternotesfetchhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tnfhq.sql != nil {
		selector = tnfhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tnfhq.ctx.Unique != nil && *tnfhq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tnfhq.schemaConfig.TwitterNotesFetchHistory)
	ctx = internal.NewSchemaConfigContext(ctx, tnfhq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range tnfhq.modifiers {
		m(selector)
	}
	for _, p := range tnfhq.predicates {
		p(selector)
	}
	for _, p := range tnfhq.order {
		p(selector)
	}
	if offset := tnfhq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tnfhq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tnfhq *TwitterNotesFetchHistoryQuery) ForUpdate(opts ...sql.LockOption) *TwitterNotesFetchHistoryQuery {
	if tnfhq.driver.Dialect() == dialect.Postgres {
		tnfhq.Unique(false)
	}
	tnfhq.modifiers = append(tnfhq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tnfhq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tnfhq *TwitterNotesFetchHistoryQuery) ForShare(opts ...sql.LockOption) *TwitterNotesFetchHistoryQuery {
	if tnfhq.driver.Dialect() == dialect.Postgres {
		tnfhq.Unique(false)
	}
	tnfhq.modifiers = append(tnfhq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tnfhq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tnfhq *TwitterNotesFetchHistoryQuery) Modify(modifiers ...func(s *sql.Selector)) *TwitterNotesFetchHistorySelect {
	tnfhq.modifiers = append(tnfhq.modifiers, modifiers...)
	return tnfhq.Select()
}

// TwitterNotesFetchHistoryGroupBy is the group-by builder for TwitterNotesFetchHistory entities.
type TwitterNotesFetchHistoryGroupBy struct {
	selector
	build *TwitterNotesFetchHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tnfhgb *TwitterNotesFetchHistoryGroupBy) Aggregate(fns ...AggregateFunc) *TwitterNotesFetchHistoryGroupBy {
	tnfhgb.fns = append(tnfhgb.fns, fns...)
	return tnfhgb
}

// Scan applies the selector query and scans the result into the given value.
func (tnfhgb *TwitterNotesFetchHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tnfhgb.build.ctx, "GroupBy")
	if err := tnfhgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterNotesFetchHistoryQuery, *TwitterNotesFetchHistoryGroupBy](ctx, tnfhgb.build, tnfhgb, tnfhgb.build.inters, v)
}

func (tnfhgb *TwitterNotesFetchHistoryGroupBy) sqlScan(ctx context.Context, root *TwitterNotesFetchHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tnfhgb.fns))
	for _, fn := range tnfhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tnfhgb.flds)+len(tnfhgb.fns))
		for _, f := range *tnfhgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tnfhgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tnfhgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TwitterNotesFetchHistorySelect is the builder for selecting fields of TwitterNotesFetchHistory entities.
type TwitterNotesFetchHistorySelect struct {
	*TwitterNotesFetchHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tnfhs *TwitterNotesFetchHistorySelect) Aggregate(fns ...AggregateFunc) *TwitterNotesFetchHistorySelect {
	tnfhs.fns = append(tnfhs.fns, fns...)
	return tnfhs
}

// Scan applies the selector query and scans the result into the given value.
func (tnfhs *TwitterNotesFetchHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tnfhs.ctx, "Select")
	if err := tnfhs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterNotesFetchHistoryQuery, *TwitterNotesFetchHistorySelect](ctx, tnfhs.TwitterNotesFetchHistoryQuery, tnfhs, tnfhs.inters, v)
}

func (tnfhs *TwitterNotesFetchHistorySelect) sqlScan(ctx context.Context, root *TwitterNotesFetchHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tnfhs.fns))
	for _, fn := range tnfhs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tnfhs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tnfhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tnfhs *TwitterNotesFetchHistorySelect) Modify(modifiers ...func(s *sql.Selector)) *TwitterNotesFetchHistorySelect {
	tnfhs.modifiers = append(tnfhs.modifiers, modifiers...)
	return tnfhs
}
