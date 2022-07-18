// Code generated by ent, DO NOT EDIT.

package ent

import (
	"backend/pkg/ent/predicate"
	"backend/pkg/ent/url"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// URLQuery is the builder for querying Url entities.
type URLQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Url
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the URLQuery builder.
func (uq *URLQuery) Where(ps ...predicate.Url) *URLQuery {
	uq.predicates = append(uq.predicates, ps...)
	return uq
}

// Limit adds a limit step to the query.
func (uq *URLQuery) Limit(limit int) *URLQuery {
	uq.limit = &limit
	return uq
}

// Offset adds an offset step to the query.
func (uq *URLQuery) Offset(offset int) *URLQuery {
	uq.offset = &offset
	return uq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uq *URLQuery) Unique(unique bool) *URLQuery {
	uq.unique = &unique
	return uq
}

// Order adds an order step to the query.
func (uq *URLQuery) Order(o ...OrderFunc) *URLQuery {
	uq.order = append(uq.order, o...)
	return uq
}

// First returns the first Url entity from the query.
// Returns a *NotFoundError when no Url was found.
func (uq *URLQuery) First(ctx context.Context) (*Url, error) {
	nodes, err := uq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{url.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uq *URLQuery) FirstX(ctx context.Context) *Url {
	node, err := uq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Url ID from the query.
// Returns a *NotFoundError when no Url ID was found.
func (uq *URLQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{url.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uq *URLQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := uq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Url entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Url entity is found.
// Returns a *NotFoundError when no Url entities are found.
func (uq *URLQuery) Only(ctx context.Context) (*Url, error) {
	nodes, err := uq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{url.Label}
	default:
		return nil, &NotSingularError{url.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uq *URLQuery) OnlyX(ctx context.Context) *Url {
	node, err := uq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Url ID in the query.
// Returns a *NotSingularError when more than one Url ID is found.
// Returns a *NotFoundError when no entities are found.
func (uq *URLQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = uq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{url.Label}
	default:
		err = &NotSingularError{url.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uq *URLQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := uq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Urls.
func (uq *URLQuery) All(ctx context.Context) ([]*Url, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uq *URLQuery) AllX(ctx context.Context) []*Url {
	nodes, err := uq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Url IDs.
func (uq *URLQuery) IDs(ctx context.Context) ([]uuid.UUID, error) {
	var ids []uuid.UUID
	if err := uq.Select(url.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uq *URLQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := uq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uq *URLQuery) Count(ctx context.Context) (int, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uq *URLQuery) CountX(ctx context.Context) int {
	count, err := uq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uq *URLQuery) Exist(ctx context.Context) (bool, error) {
	if err := uq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uq *URLQuery) ExistX(ctx context.Context) bool {
	exist, err := uq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the URLQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uq *URLQuery) Clone() *URLQuery {
	if uq == nil {
		return nil
	}
	return &URLQuery{
		config:     uq.config,
		limit:      uq.limit,
		offset:     uq.offset,
		order:      append([]OrderFunc{}, uq.order...),
		predicates: append([]predicate.Url{}, uq.predicates...),
		// clone intermediate query.
		sql:    uq.sql.Clone(),
		path:   uq.path,
		unique: uq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ShortID string `json:"short_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.URL.Query().
//		GroupBy(url.FieldShortID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (uq *URLQuery) GroupBy(field string, fields ...string) *URLGroupBy {
	grbuild := &URLGroupBy{config: uq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uq.sqlQuery(ctx), nil
	}
	grbuild.label = url.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ShortID string `json:"short_id,omitempty"`
//	}
//
//	client.URL.Query().
//		Select(url.FieldShortID).
//		Scan(ctx, &v)
//
func (uq *URLQuery) Select(fields ...string) *URLSelect {
	uq.fields = append(uq.fields, fields...)
	selbuild := &URLSelect{URLQuery: uq}
	selbuild.label = url.Label
	selbuild.flds, selbuild.scan = &uq.fields, selbuild.Scan
	return selbuild
}

func (uq *URLQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uq.fields {
		if !url.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uq.path != nil {
		prev, err := uq.path(ctx)
		if err != nil {
			return err
		}
		uq.sql = prev
	}
	return nil
}

func (uq *URLQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Url, error) {
	var (
		nodes = []*Url{}
		_spec = uq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Url).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Url{config: uq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uq *URLQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uq.querySpec()
	_spec.Node.Columns = uq.fields
	if len(uq.fields) > 0 {
		_spec.Unique = uq.unique != nil && *uq.unique
	}
	return sqlgraph.CountNodes(ctx, uq.driver, _spec)
}

func (uq *URLQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (uq *URLQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   url.Table,
			Columns: url.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: url.FieldID,
			},
		},
		From:   uq.sql,
		Unique: true,
	}
	if unique := uq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, url.FieldID)
		for i := range fields {
			if fields[i] != url.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uq *URLQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uq.driver.Dialect())
	t1 := builder.Table(url.Table)
	columns := uq.fields
	if len(columns) == 0 {
		columns = url.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uq.sql != nil {
		selector = uq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uq.unique != nil && *uq.unique {
		selector.Distinct()
	}
	for _, p := range uq.predicates {
		p(selector)
	}
	for _, p := range uq.order {
		p(selector)
	}
	if offset := uq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// URLGroupBy is the group-by builder for Url entities.
type URLGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ugb *URLGroupBy) Aggregate(fns ...AggregateFunc) *URLGroupBy {
	ugb.fns = append(ugb.fns, fns...)
	return ugb
}

// Scan applies the group-by query and scans the result into the given value.
func (ugb *URLGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ugb.path(ctx)
	if err != nil {
		return err
	}
	ugb.sql = query
	return ugb.sqlScan(ctx, v)
}

func (ugb *URLGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ugb.fields {
		if !url.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ugb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ugb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ugb *URLGroupBy) sqlQuery() *sql.Selector {
	selector := ugb.sql.Select()
	aggregation := make([]string, 0, len(ugb.fns))
	for _, fn := range ugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ugb.fields)+len(ugb.fns))
		for _, f := range ugb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ugb.fields...)...)
}

// URLSelect is the builder for selecting fields of URL entities.
type URLSelect struct {
	*URLQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (us *URLSelect) Scan(ctx context.Context, v interface{}) error {
	if err := us.prepareQuery(ctx); err != nil {
		return err
	}
	us.sql = us.URLQuery.sqlQuery(ctx)
	return us.sqlScan(ctx, v)
}

func (us *URLSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := us.sql.Query()
	if err := us.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
