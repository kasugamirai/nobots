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
	"freefrom.space/nobot/core/twitterfetchhistory"
)

// TwitterFetchHistoryQuery is the builder for querying TwitterFetchHistory entities.
type TwitterFetchHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []twitterfetchhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.TwitterFetchHistory
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TwitterFetchHistoryQuery builder.
func (tfhq *TwitterFetchHistoryQuery) Where(ps ...predicate.TwitterFetchHistory) *TwitterFetchHistoryQuery {
	tfhq.predicates = append(tfhq.predicates, ps...)
	return tfhq
}

// Limit the number of records to be returned by this query.
func (tfhq *TwitterFetchHistoryQuery) Limit(limit int) *TwitterFetchHistoryQuery {
	tfhq.ctx.Limit = &limit
	return tfhq
}

// Offset to start from.
func (tfhq *TwitterFetchHistoryQuery) Offset(offset int) *TwitterFetchHistoryQuery {
	tfhq.ctx.Offset = &offset
	return tfhq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tfhq *TwitterFetchHistoryQuery) Unique(unique bool) *TwitterFetchHistoryQuery {
	tfhq.ctx.Unique = &unique
	return tfhq
}

// Order specifies how the records should be ordered.
func (tfhq *TwitterFetchHistoryQuery) Order(o ...twitterfetchhistory.OrderOption) *TwitterFetchHistoryQuery {
	tfhq.order = append(tfhq.order, o...)
	return tfhq
}

// First returns the first TwitterFetchHistory entity from the query.
// Returns a *NotFoundError when no TwitterFetchHistory was found.
func (tfhq *TwitterFetchHistoryQuery) First(ctx context.Context) (*TwitterFetchHistory, error) {
	nodes, err := tfhq.Limit(1).All(setContextOp(ctx, tfhq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{twitterfetchhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) FirstX(ctx context.Context) *TwitterFetchHistory {
	node, err := tfhq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TwitterFetchHistory ID from the query.
// Returns a *NotFoundError when no TwitterFetchHistory ID was found.
func (tfhq *TwitterFetchHistoryQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tfhq.Limit(1).IDs(setContextOp(ctx, tfhq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{twitterfetchhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) FirstIDX(ctx context.Context) int {
	id, err := tfhq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TwitterFetchHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TwitterFetchHistory entity is found.
// Returns a *NotFoundError when no TwitterFetchHistory entities are found.
func (tfhq *TwitterFetchHistoryQuery) Only(ctx context.Context) (*TwitterFetchHistory, error) {
	nodes, err := tfhq.Limit(2).All(setContextOp(ctx, tfhq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{twitterfetchhistory.Label}
	default:
		return nil, &NotSingularError{twitterfetchhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) OnlyX(ctx context.Context) *TwitterFetchHistory {
	node, err := tfhq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TwitterFetchHistory ID in the query.
// Returns a *NotSingularError when more than one TwitterFetchHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (tfhq *TwitterFetchHistoryQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tfhq.Limit(2).IDs(setContextOp(ctx, tfhq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{twitterfetchhistory.Label}
	default:
		err = &NotSingularError{twitterfetchhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) OnlyIDX(ctx context.Context) int {
	id, err := tfhq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TwitterFetchHistories.
func (tfhq *TwitterFetchHistoryQuery) All(ctx context.Context) ([]*TwitterFetchHistory, error) {
	ctx = setContextOp(ctx, tfhq.ctx, "All")
	if err := tfhq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TwitterFetchHistory, *TwitterFetchHistoryQuery]()
	return withInterceptors[[]*TwitterFetchHistory](ctx, tfhq, qr, tfhq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) AllX(ctx context.Context) []*TwitterFetchHistory {
	nodes, err := tfhq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TwitterFetchHistory IDs.
func (tfhq *TwitterFetchHistoryQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tfhq.ctx.Unique == nil && tfhq.path != nil {
		tfhq.Unique(true)
	}
	ctx = setContextOp(ctx, tfhq.ctx, "IDs")
	if err = tfhq.Select(twitterfetchhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) IDsX(ctx context.Context) []int {
	ids, err := tfhq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tfhq *TwitterFetchHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tfhq.ctx, "Count")
	if err := tfhq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tfhq, querierCount[*TwitterFetchHistoryQuery](), tfhq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) CountX(ctx context.Context) int {
	count, err := tfhq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tfhq *TwitterFetchHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tfhq.ctx, "Exist")
	switch _, err := tfhq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("core: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tfhq *TwitterFetchHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := tfhq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TwitterFetchHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tfhq *TwitterFetchHistoryQuery) Clone() *TwitterFetchHistoryQuery {
	if tfhq == nil {
		return nil
	}
	return &TwitterFetchHistoryQuery{
		config:     tfhq.config,
		ctx:        tfhq.ctx.Clone(),
		order:      append([]twitterfetchhistory.OrderOption{}, tfhq.order...),
		inters:     append([]Interceptor{}, tfhq.inters...),
		predicates: append([]predicate.TwitterFetchHistory{}, tfhq.predicates...),
		// clone intermediate query.
		sql:  tfhq.sql.Clone(),
		path: tfhq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ItemID string `json:"item_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.TwitterFetchHistory.Query().
//		GroupBy(twitterfetchhistory.FieldItemID).
//		Aggregate(core.Count()).
//		Scan(ctx, &v)
func (tfhq *TwitterFetchHistoryQuery) GroupBy(field string, fields ...string) *TwitterFetchHistoryGroupBy {
	tfhq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TwitterFetchHistoryGroupBy{build: tfhq}
	grbuild.flds = &tfhq.ctx.Fields
	grbuild.label = twitterfetchhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ItemID string `json:"item_id,omitempty"`
//	}
//
//	client.TwitterFetchHistory.Query().
//		Select(twitterfetchhistory.FieldItemID).
//		Scan(ctx, &v)
func (tfhq *TwitterFetchHistoryQuery) Select(fields ...string) *TwitterFetchHistorySelect {
	tfhq.ctx.Fields = append(tfhq.ctx.Fields, fields...)
	sbuild := &TwitterFetchHistorySelect{TwitterFetchHistoryQuery: tfhq}
	sbuild.label = twitterfetchhistory.Label
	sbuild.flds, sbuild.scan = &tfhq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TwitterFetchHistorySelect configured with the given aggregations.
func (tfhq *TwitterFetchHistoryQuery) Aggregate(fns ...AggregateFunc) *TwitterFetchHistorySelect {
	return tfhq.Select().Aggregate(fns...)
}

func (tfhq *TwitterFetchHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tfhq.inters {
		if inter == nil {
			return fmt.Errorf("core: uninitialized interceptor (forgotten import core/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tfhq); err != nil {
				return err
			}
		}
	}
	for _, f := range tfhq.ctx.Fields {
		if !twitterfetchhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("core: invalid field %q for query", f)}
		}
	}
	if tfhq.path != nil {
		prev, err := tfhq.path(ctx)
		if err != nil {
			return err
		}
		tfhq.sql = prev
	}
	return nil
}

func (tfhq *TwitterFetchHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TwitterFetchHistory, error) {
	var (
		nodes = []*TwitterFetchHistory{}
		_spec = tfhq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TwitterFetchHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TwitterFetchHistory{config: tfhq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tfhq.schemaConfig.TwitterFetchHistory
	ctx = internal.NewSchemaConfigContext(ctx, tfhq.schemaConfig)
	if len(tfhq.modifiers) > 0 {
		_spec.Modifiers = tfhq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tfhq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tfhq *TwitterFetchHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tfhq.querySpec()
	_spec.Node.Schema = tfhq.schemaConfig.TwitterFetchHistory
	ctx = internal.NewSchemaConfigContext(ctx, tfhq.schemaConfig)
	if len(tfhq.modifiers) > 0 {
		_spec.Modifiers = tfhq.modifiers
	}
	_spec.Node.Columns = tfhq.ctx.Fields
	if len(tfhq.ctx.Fields) > 0 {
		_spec.Unique = tfhq.ctx.Unique != nil && *tfhq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tfhq.driver, _spec)
}

func (tfhq *TwitterFetchHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(twitterfetchhistory.Table, twitterfetchhistory.Columns, sqlgraph.NewFieldSpec(twitterfetchhistory.FieldID, field.TypeInt))
	_spec.From = tfhq.sql
	if unique := tfhq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tfhq.path != nil {
		_spec.Unique = true
	}
	if fields := tfhq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twitterfetchhistory.FieldID)
		for i := range fields {
			if fields[i] != twitterfetchhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tfhq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tfhq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tfhq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tfhq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tfhq *TwitterFetchHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tfhq.driver.Dialect())
	t1 := builder.Table(twitterfetchhistory.Table)
	columns := tfhq.ctx.Fields
	if len(columns) == 0 {
		columns = twitterfetchhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tfhq.sql != nil {
		selector = tfhq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tfhq.ctx.Unique != nil && *tfhq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tfhq.schemaConfig.TwitterFetchHistory)
	ctx = internal.NewSchemaConfigContext(ctx, tfhq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range tfhq.modifiers {
		m(selector)
	}
	for _, p := range tfhq.predicates {
		p(selector)
	}
	for _, p := range tfhq.order {
		p(selector)
	}
	if offset := tfhq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tfhq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tfhq *TwitterFetchHistoryQuery) ForUpdate(opts ...sql.LockOption) *TwitterFetchHistoryQuery {
	if tfhq.driver.Dialect() == dialect.Postgres {
		tfhq.Unique(false)
	}
	tfhq.modifiers = append(tfhq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tfhq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tfhq *TwitterFetchHistoryQuery) ForShare(opts ...sql.LockOption) *TwitterFetchHistoryQuery {
	if tfhq.driver.Dialect() == dialect.Postgres {
		tfhq.Unique(false)
	}
	tfhq.modifiers = append(tfhq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tfhq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tfhq *TwitterFetchHistoryQuery) Modify(modifiers ...func(s *sql.Selector)) *TwitterFetchHistorySelect {
	tfhq.modifiers = append(tfhq.modifiers, modifiers...)
	return tfhq.Select()
}

// TwitterFetchHistoryGroupBy is the group-by builder for TwitterFetchHistory entities.
type TwitterFetchHistoryGroupBy struct {
	selector
	build *TwitterFetchHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tfhgb *TwitterFetchHistoryGroupBy) Aggregate(fns ...AggregateFunc) *TwitterFetchHistoryGroupBy {
	tfhgb.fns = append(tfhgb.fns, fns...)
	return tfhgb
}

// Scan applies the selector query and scans the result into the given value.
func (tfhgb *TwitterFetchHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tfhgb.build.ctx, "GroupBy")
	if err := tfhgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterFetchHistoryQuery, *TwitterFetchHistoryGroupBy](ctx, tfhgb.build, tfhgb, tfhgb.build.inters, v)
}

func (tfhgb *TwitterFetchHistoryGroupBy) sqlScan(ctx context.Context, root *TwitterFetchHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tfhgb.fns))
	for _, fn := range tfhgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tfhgb.flds)+len(tfhgb.fns))
		for _, f := range *tfhgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tfhgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tfhgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TwitterFetchHistorySelect is the builder for selecting fields of TwitterFetchHistory entities.
type TwitterFetchHistorySelect struct {
	*TwitterFetchHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tfhs *TwitterFetchHistorySelect) Aggregate(fns ...AggregateFunc) *TwitterFetchHistorySelect {
	tfhs.fns = append(tfhs.fns, fns...)
	return tfhs
}

// Scan applies the selector query and scans the result into the given value.
func (tfhs *TwitterFetchHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tfhs.ctx, "Select")
	if err := tfhs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterFetchHistoryQuery, *TwitterFetchHistorySelect](ctx, tfhs.TwitterFetchHistoryQuery, tfhs, tfhs.inters, v)
}

func (tfhs *TwitterFetchHistorySelect) sqlScan(ctx context.Context, root *TwitterFetchHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tfhs.fns))
	for _, fn := range tfhs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tfhs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tfhs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tfhs *TwitterFetchHistorySelect) Modify(modifiers ...func(s *sql.Selector)) *TwitterFetchHistorySelect {
	tfhs.modifiers = append(tfhs.modifiers, modifiers...)
	return tfhs
}
