// Code generated by ent, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/theopenlane/core/internal/ent/generated/predicate"
	"github.com/theopenlane/core/internal/ent/generated/webhookhistory"

	"github.com/theopenlane/core/internal/ent/generated/internal"
)

// WebhookHistoryQuery is the builder for querying WebhookHistory entities.
type WebhookHistoryQuery struct {
	config
	ctx        *QueryContext
	order      []webhookhistory.OrderOption
	inters     []Interceptor
	predicates []predicate.WebhookHistory
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*WebhookHistory) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the WebhookHistoryQuery builder.
func (whq *WebhookHistoryQuery) Where(ps ...predicate.WebhookHistory) *WebhookHistoryQuery {
	whq.predicates = append(whq.predicates, ps...)
	return whq
}

// Limit the number of records to be returned by this query.
func (whq *WebhookHistoryQuery) Limit(limit int) *WebhookHistoryQuery {
	whq.ctx.Limit = &limit
	return whq
}

// Offset to start from.
func (whq *WebhookHistoryQuery) Offset(offset int) *WebhookHistoryQuery {
	whq.ctx.Offset = &offset
	return whq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (whq *WebhookHistoryQuery) Unique(unique bool) *WebhookHistoryQuery {
	whq.ctx.Unique = &unique
	return whq
}

// Order specifies how the records should be ordered.
func (whq *WebhookHistoryQuery) Order(o ...webhookhistory.OrderOption) *WebhookHistoryQuery {
	whq.order = append(whq.order, o...)
	return whq
}

// First returns the first WebhookHistory entity from the query.
// Returns a *NotFoundError when no WebhookHistory was found.
func (whq *WebhookHistoryQuery) First(ctx context.Context) (*WebhookHistory, error) {
	nodes, err := whq.Limit(1).All(setContextOp(ctx, whq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{webhookhistory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (whq *WebhookHistoryQuery) FirstX(ctx context.Context) *WebhookHistory {
	node, err := whq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first WebhookHistory ID from the query.
// Returns a *NotFoundError when no WebhookHistory ID was found.
func (whq *WebhookHistoryQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = whq.Limit(1).IDs(setContextOp(ctx, whq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{webhookhistory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (whq *WebhookHistoryQuery) FirstIDX(ctx context.Context) string {
	id, err := whq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single WebhookHistory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one WebhookHistory entity is found.
// Returns a *NotFoundError when no WebhookHistory entities are found.
func (whq *WebhookHistoryQuery) Only(ctx context.Context) (*WebhookHistory, error) {
	nodes, err := whq.Limit(2).All(setContextOp(ctx, whq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{webhookhistory.Label}
	default:
		return nil, &NotSingularError{webhookhistory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (whq *WebhookHistoryQuery) OnlyX(ctx context.Context) *WebhookHistory {
	node, err := whq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only WebhookHistory ID in the query.
// Returns a *NotSingularError when more than one WebhookHistory ID is found.
// Returns a *NotFoundError when no entities are found.
func (whq *WebhookHistoryQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = whq.Limit(2).IDs(setContextOp(ctx, whq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{webhookhistory.Label}
	default:
		err = &NotSingularError{webhookhistory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (whq *WebhookHistoryQuery) OnlyIDX(ctx context.Context) string {
	id, err := whq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of WebhookHistories.
func (whq *WebhookHistoryQuery) All(ctx context.Context) ([]*WebhookHistory, error) {
	ctx = setContextOp(ctx, whq.ctx, ent.OpQueryAll)
	if err := whq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*WebhookHistory, *WebhookHistoryQuery]()
	return withInterceptors[[]*WebhookHistory](ctx, whq, qr, whq.inters)
}

// AllX is like All, but panics if an error occurs.
func (whq *WebhookHistoryQuery) AllX(ctx context.Context) []*WebhookHistory {
	nodes, err := whq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of WebhookHistory IDs.
func (whq *WebhookHistoryQuery) IDs(ctx context.Context) (ids []string, err error) {
	if whq.ctx.Unique == nil && whq.path != nil {
		whq.Unique(true)
	}
	ctx = setContextOp(ctx, whq.ctx, ent.OpQueryIDs)
	if err = whq.Select(webhookhistory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (whq *WebhookHistoryQuery) IDsX(ctx context.Context) []string {
	ids, err := whq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (whq *WebhookHistoryQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, whq.ctx, ent.OpQueryCount)
	if err := whq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, whq, querierCount[*WebhookHistoryQuery](), whq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (whq *WebhookHistoryQuery) CountX(ctx context.Context) int {
	count, err := whq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (whq *WebhookHistoryQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, whq.ctx, ent.OpQueryExist)
	switch _, err := whq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("generated: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (whq *WebhookHistoryQuery) ExistX(ctx context.Context) bool {
	exist, err := whq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the WebhookHistoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (whq *WebhookHistoryQuery) Clone() *WebhookHistoryQuery {
	if whq == nil {
		return nil
	}
	return &WebhookHistoryQuery{
		config:     whq.config,
		ctx:        whq.ctx.Clone(),
		order:      append([]webhookhistory.OrderOption{}, whq.order...),
		inters:     append([]Interceptor{}, whq.inters...),
		predicates: append([]predicate.WebhookHistory{}, whq.predicates...),
		// clone intermediate query.
		sql:  whq.sql.Clone(),
		path: whq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.WebhookHistory.Query().
//		GroupBy(webhookhistory.FieldHistoryTime).
//		Aggregate(generated.Count()).
//		Scan(ctx, &v)
func (whq *WebhookHistoryQuery) GroupBy(field string, fields ...string) *WebhookHistoryGroupBy {
	whq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &WebhookHistoryGroupBy{build: whq}
	grbuild.flds = &whq.ctx.Fields
	grbuild.label = webhookhistory.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		HistoryTime time.Time `json:"history_time,omitempty"`
//	}
//
//	client.WebhookHistory.Query().
//		Select(webhookhistory.FieldHistoryTime).
//		Scan(ctx, &v)
func (whq *WebhookHistoryQuery) Select(fields ...string) *WebhookHistorySelect {
	whq.ctx.Fields = append(whq.ctx.Fields, fields...)
	sbuild := &WebhookHistorySelect{WebhookHistoryQuery: whq}
	sbuild.label = webhookhistory.Label
	sbuild.flds, sbuild.scan = &whq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a WebhookHistorySelect configured with the given aggregations.
func (whq *WebhookHistoryQuery) Aggregate(fns ...AggregateFunc) *WebhookHistorySelect {
	return whq.Select().Aggregate(fns...)
}

func (whq *WebhookHistoryQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range whq.inters {
		if inter == nil {
			return fmt.Errorf("generated: uninitialized interceptor (forgotten import generated/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, whq); err != nil {
				return err
			}
		}
	}
	for _, f := range whq.ctx.Fields {
		if !webhookhistory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
		}
	}
	if whq.path != nil {
		prev, err := whq.path(ctx)
		if err != nil {
			return err
		}
		whq.sql = prev
	}
	if webhookhistory.Policy == nil {
		return errors.New("generated: uninitialized webhookhistory.Policy (forgotten import generated/runtime?)")
	}
	if err := webhookhistory.Policy.EvalQuery(ctx, whq); err != nil {
		return err
	}
	return nil
}

func (whq *WebhookHistoryQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*WebhookHistory, error) {
	var (
		nodes = []*WebhookHistory{}
		_spec = whq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*WebhookHistory).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &WebhookHistory{config: whq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = whq.schemaConfig.WebhookHistory
	ctx = internal.NewSchemaConfigContext(ctx, whq.schemaConfig)
	if len(whq.modifiers) > 0 {
		_spec.Modifiers = whq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, whq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range whq.loadTotal {
		if err := whq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (whq *WebhookHistoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := whq.querySpec()
	_spec.Node.Schema = whq.schemaConfig.WebhookHistory
	ctx = internal.NewSchemaConfigContext(ctx, whq.schemaConfig)
	if len(whq.modifiers) > 0 {
		_spec.Modifiers = whq.modifiers
	}
	_spec.Node.Columns = whq.ctx.Fields
	if len(whq.ctx.Fields) > 0 {
		_spec.Unique = whq.ctx.Unique != nil && *whq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, whq.driver, _spec)
}

func (whq *WebhookHistoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(webhookhistory.Table, webhookhistory.Columns, sqlgraph.NewFieldSpec(webhookhistory.FieldID, field.TypeString))
	_spec.From = whq.sql
	if unique := whq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if whq.path != nil {
		_spec.Unique = true
	}
	if fields := whq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, webhookhistory.FieldID)
		for i := range fields {
			if fields[i] != webhookhistory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := whq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := whq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := whq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := whq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (whq *WebhookHistoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(whq.driver.Dialect())
	t1 := builder.Table(webhookhistory.Table)
	columns := whq.ctx.Fields
	if len(columns) == 0 {
		columns = webhookhistory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if whq.sql != nil {
		selector = whq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if whq.ctx.Unique != nil && *whq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(whq.schemaConfig.WebhookHistory)
	ctx = internal.NewSchemaConfigContext(ctx, whq.schemaConfig)
	selector.WithContext(ctx)
	for _, p := range whq.predicates {
		p(selector)
	}
	for _, p := range whq.order {
		p(selector)
	}
	if offset := whq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := whq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WebhookHistoryGroupBy is the group-by builder for WebhookHistory entities.
type WebhookHistoryGroupBy struct {
	selector
	build *WebhookHistoryQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (whgb *WebhookHistoryGroupBy) Aggregate(fns ...AggregateFunc) *WebhookHistoryGroupBy {
	whgb.fns = append(whgb.fns, fns...)
	return whgb
}

// Scan applies the selector query and scans the result into the given value.
func (whgb *WebhookHistoryGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, whgb.build.ctx, ent.OpQueryGroupBy)
	if err := whgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebhookHistoryQuery, *WebhookHistoryGroupBy](ctx, whgb.build, whgb, whgb.build.inters, v)
}

func (whgb *WebhookHistoryGroupBy) sqlScan(ctx context.Context, root *WebhookHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(whgb.fns))
	for _, fn := range whgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*whgb.flds)+len(whgb.fns))
		for _, f := range *whgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*whgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := whgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// WebhookHistorySelect is the builder for selecting fields of WebhookHistory entities.
type WebhookHistorySelect struct {
	*WebhookHistoryQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (whs *WebhookHistorySelect) Aggregate(fns ...AggregateFunc) *WebhookHistorySelect {
	whs.fns = append(whs.fns, fns...)
	return whs
}

// Scan applies the selector query and scans the result into the given value.
func (whs *WebhookHistorySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, whs.ctx, ent.OpQuerySelect)
	if err := whs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*WebhookHistoryQuery, *WebhookHistorySelect](ctx, whs.WebhookHistoryQuery, whs, whs.inters, v)
}

func (whs *WebhookHistorySelect) sqlScan(ctx context.Context, root *WebhookHistoryQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(whs.fns))
	for _, fn := range whs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*whs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := whs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
