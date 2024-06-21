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
	"freefrom.space/nobot/core/twitternote"
)

// TwitterNoteQuery is the builder for querying TwitterNote entities.
type TwitterNoteQuery struct {
	config
	ctx        *QueryContext
	order      []twitternote.OrderOption
	inters     []Interceptor
	predicates []predicate.TwitterNote
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TwitterNoteQuery builder.
func (tnq *TwitterNoteQuery) Where(ps ...predicate.TwitterNote) *TwitterNoteQuery {
	tnq.predicates = append(tnq.predicates, ps...)
	return tnq
}

// Limit the number of records to be returned by this query.
func (tnq *TwitterNoteQuery) Limit(limit int) *TwitterNoteQuery {
	tnq.ctx.Limit = &limit
	return tnq
}

// Offset to start from.
func (tnq *TwitterNoteQuery) Offset(offset int) *TwitterNoteQuery {
	tnq.ctx.Offset = &offset
	return tnq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tnq *TwitterNoteQuery) Unique(unique bool) *TwitterNoteQuery {
	tnq.ctx.Unique = &unique
	return tnq
}

// Order specifies how the records should be ordered.
func (tnq *TwitterNoteQuery) Order(o ...twitternote.OrderOption) *TwitterNoteQuery {
	tnq.order = append(tnq.order, o...)
	return tnq
}

// First returns the first TwitterNote entity from the query.
// Returns a *NotFoundError when no TwitterNote was found.
func (tnq *TwitterNoteQuery) First(ctx context.Context) (*TwitterNote, error) {
	nodes, err := tnq.Limit(1).All(setContextOp(ctx, tnq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{twitternote.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tnq *TwitterNoteQuery) FirstX(ctx context.Context) *TwitterNote {
	node, err := tnq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TwitterNote ID from the query.
// Returns a *NotFoundError when no TwitterNote ID was found.
func (tnq *TwitterNoteQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tnq.Limit(1).IDs(setContextOp(ctx, tnq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{twitternote.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tnq *TwitterNoteQuery) FirstIDX(ctx context.Context) string {
	id, err := tnq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TwitterNote entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TwitterNote entity is found.
// Returns a *NotFoundError when no TwitterNote entities are found.
func (tnq *TwitterNoteQuery) Only(ctx context.Context) (*TwitterNote, error) {
	nodes, err := tnq.Limit(2).All(setContextOp(ctx, tnq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{twitternote.Label}
	default:
		return nil, &NotSingularError{twitternote.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tnq *TwitterNoteQuery) OnlyX(ctx context.Context) *TwitterNote {
	node, err := tnq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TwitterNote ID in the query.
// Returns a *NotSingularError when more than one TwitterNote ID is found.
// Returns a *NotFoundError when no entities are found.
func (tnq *TwitterNoteQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tnq.Limit(2).IDs(setContextOp(ctx, tnq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{twitternote.Label}
	default:
		err = &NotSingularError{twitternote.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tnq *TwitterNoteQuery) OnlyIDX(ctx context.Context) string {
	id, err := tnq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TwitterNotes.
func (tnq *TwitterNoteQuery) All(ctx context.Context) ([]*TwitterNote, error) {
	ctx = setContextOp(ctx, tnq.ctx, "All")
	if err := tnq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TwitterNote, *TwitterNoteQuery]()
	return withInterceptors[[]*TwitterNote](ctx, tnq, qr, tnq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tnq *TwitterNoteQuery) AllX(ctx context.Context) []*TwitterNote {
	nodes, err := tnq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TwitterNote IDs.
func (tnq *TwitterNoteQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tnq.ctx.Unique == nil && tnq.path != nil {
		tnq.Unique(true)
	}
	ctx = setContextOp(ctx, tnq.ctx, "IDs")
	if err = tnq.Select(twitternote.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tnq *TwitterNoteQuery) IDsX(ctx context.Context) []string {
	ids, err := tnq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tnq *TwitterNoteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tnq.ctx, "Count")
	if err := tnq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tnq, querierCount[*TwitterNoteQuery](), tnq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tnq *TwitterNoteQuery) CountX(ctx context.Context) int {
	count, err := tnq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tnq *TwitterNoteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tnq.ctx, "Exist")
	switch _, err := tnq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("core: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tnq *TwitterNoteQuery) ExistX(ctx context.Context) bool {
	exist, err := tnq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TwitterNoteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tnq *TwitterNoteQuery) Clone() *TwitterNoteQuery {
	if tnq == nil {
		return nil
	}
	return &TwitterNoteQuery{
		config:     tnq.config,
		ctx:        tnq.ctx.Clone(),
		order:      append([]twitternote.OrderOption{}, tnq.order...),
		inters:     append([]Interceptor{}, tnq.inters...),
		predicates: append([]predicate.TwitterNote{}, tnq.predicates...),
		// clone intermediate query.
		sql:  tnq.sql.Clone(),
		path: tnq.path,
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
//	client.TwitterNote.Query().
//		GroupBy(twitternote.FieldUserID).
//		Aggregate(core.Count()).
//		Scan(ctx, &v)
func (tnq *TwitterNoteQuery) GroupBy(field string, fields ...string) *TwitterNoteGroupBy {
	tnq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TwitterNoteGroupBy{build: tnq}
	grbuild.flds = &tnq.ctx.Fields
	grbuild.label = twitternote.Label
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
//	client.TwitterNote.Query().
//		Select(twitternote.FieldUserID).
//		Scan(ctx, &v)
func (tnq *TwitterNoteQuery) Select(fields ...string) *TwitterNoteSelect {
	tnq.ctx.Fields = append(tnq.ctx.Fields, fields...)
	sbuild := &TwitterNoteSelect{TwitterNoteQuery: tnq}
	sbuild.label = twitternote.Label
	sbuild.flds, sbuild.scan = &tnq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TwitterNoteSelect configured with the given aggregations.
func (tnq *TwitterNoteQuery) Aggregate(fns ...AggregateFunc) *TwitterNoteSelect {
	return tnq.Select().Aggregate(fns...)
}

func (tnq *TwitterNoteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tnq.inters {
		if inter == nil {
			return fmt.Errorf("core: uninitialized interceptor (forgotten import core/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tnq); err != nil {
				return err
			}
		}
	}
	for _, f := range tnq.ctx.Fields {
		if !twitternote.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("core: invalid field %q for query", f)}
		}
	}
	if tnq.path != nil {
		prev, err := tnq.path(ctx)
		if err != nil {
			return err
		}
		tnq.sql = prev
	}
	return nil
}

func (tnq *TwitterNoteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TwitterNote, error) {
	var (
		nodes = []*TwitterNote{}
		_spec = tnq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TwitterNote).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TwitterNote{config: tnq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = tnq.schemaConfig.TwitterNote
	ctx = internal.NewSchemaConfigContext(ctx, tnq.schemaConfig)
	if len(tnq.modifiers) > 0 {
		_spec.Modifiers = tnq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tnq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tnq *TwitterNoteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tnq.querySpec()
	_spec.Node.Schema = tnq.schemaConfig.TwitterNote
	ctx = internal.NewSchemaConfigContext(ctx, tnq.schemaConfig)
	if len(tnq.modifiers) > 0 {
		_spec.Modifiers = tnq.modifiers
	}
	_spec.Node.Columns = tnq.ctx.Fields
	if len(tnq.ctx.Fields) > 0 {
		_spec.Unique = tnq.ctx.Unique != nil && *tnq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tnq.driver, _spec)
}

func (tnq *TwitterNoteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(twitternote.Table, twitternote.Columns, sqlgraph.NewFieldSpec(twitternote.FieldID, field.TypeString))
	_spec.From = tnq.sql
	if unique := tnq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tnq.path != nil {
		_spec.Unique = true
	}
	if fields := tnq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, twitternote.FieldID)
		for i := range fields {
			if fields[i] != twitternote.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tnq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tnq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tnq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tnq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tnq *TwitterNoteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tnq.driver.Dialect())
	t1 := builder.Table(twitternote.Table)
	columns := tnq.ctx.Fields
	if len(columns) == 0 {
		columns = twitternote.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tnq.sql != nil {
		selector = tnq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tnq.ctx.Unique != nil && *tnq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(tnq.schemaConfig.TwitterNote)
	ctx = internal.NewSchemaConfigContext(ctx, tnq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range tnq.modifiers {
		m(selector)
	}
	for _, p := range tnq.predicates {
		p(selector)
	}
	for _, p := range tnq.order {
		p(selector)
	}
	if offset := tnq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tnq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (tnq *TwitterNoteQuery) ForUpdate(opts ...sql.LockOption) *TwitterNoteQuery {
	if tnq.driver.Dialect() == dialect.Postgres {
		tnq.Unique(false)
	}
	tnq.modifiers = append(tnq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return tnq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (tnq *TwitterNoteQuery) ForShare(opts ...sql.LockOption) *TwitterNoteQuery {
	if tnq.driver.Dialect() == dialect.Postgres {
		tnq.Unique(false)
	}
	tnq.modifiers = append(tnq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return tnq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tnq *TwitterNoteQuery) Modify(modifiers ...func(s *sql.Selector)) *TwitterNoteSelect {
	tnq.modifiers = append(tnq.modifiers, modifiers...)
	return tnq.Select()
}

// TwitterNoteGroupBy is the group-by builder for TwitterNote entities.
type TwitterNoteGroupBy struct {
	selector
	build *TwitterNoteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tngb *TwitterNoteGroupBy) Aggregate(fns ...AggregateFunc) *TwitterNoteGroupBy {
	tngb.fns = append(tngb.fns, fns...)
	return tngb
}

// Scan applies the selector query and scans the result into the given value.
func (tngb *TwitterNoteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tngb.build.ctx, "GroupBy")
	if err := tngb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterNoteQuery, *TwitterNoteGroupBy](ctx, tngb.build, tngb, tngb.build.inters, v)
}

func (tngb *TwitterNoteGroupBy) sqlScan(ctx context.Context, root *TwitterNoteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tngb.fns))
	for _, fn := range tngb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tngb.flds)+len(tngb.fns))
		for _, f := range *tngb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tngb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tngb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TwitterNoteSelect is the builder for selecting fields of TwitterNote entities.
type TwitterNoteSelect struct {
	*TwitterNoteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tns *TwitterNoteSelect) Aggregate(fns ...AggregateFunc) *TwitterNoteSelect {
	tns.fns = append(tns.fns, fns...)
	return tns
}

// Scan applies the selector query and scans the result into the given value.
func (tns *TwitterNoteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tns.ctx, "Select")
	if err := tns.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TwitterNoteQuery, *TwitterNoteSelect](ctx, tns.TwitterNoteQuery, tns, tns.inters, v)
}

func (tns *TwitterNoteSelect) sqlScan(ctx context.Context, root *TwitterNoteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tns.fns))
	for _, fn := range tns.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tns.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tns.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (tns *TwitterNoteSelect) Modify(modifiers ...func(s *sql.Selector)) *TwitterNoteSelect {
	tns.modifiers = append(tns.modifiers, modifiers...)
	return tns
}
