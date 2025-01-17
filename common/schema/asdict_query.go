// Code generated by entc, DO NOT EDIT.

package schema

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"math"
	"orginone/common/schema/asdict"
	"orginone/common/schema/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AsDictQuery is the builder for querying AsDict entities.
type AsDictQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.AsDict
	// eager-loading edges.
	withParentx   *AsDictQuery
	withChildrens *AsDictQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the AsDictQuery builder.
func (adq *AsDictQuery) Where(ps ...predicate.AsDict) *AsDictQuery {
	adq.predicates = append(adq.predicates, ps...)
	return adq
}

// Limit adds a limit step to the query.
func (adq *AsDictQuery) Limit(limit int) *AsDictQuery {
	adq.limit = &limit
	return adq
}

// Offset adds an offset step to the query.
func (adq *AsDictQuery) Offset(offset int) *AsDictQuery {
	adq.offset = &offset
	return adq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (adq *AsDictQuery) Unique(unique bool) *AsDictQuery {
	adq.unique = &unique
	return adq
}

// Order adds an order step to the query.
func (adq *AsDictQuery) Order(o ...OrderFunc) *AsDictQuery {
	adq.order = append(adq.order, o...)
	return adq
}

// QueryParentx chains the current query on the "parentx" edge.
func (adq *AsDictQuery) QueryParentx() *AsDictQuery {
	query := &AsDictQuery{config: adq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := adq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := adq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asdict.Table, asdict.FieldID, selector),
			sqlgraph.To(asdict.Table, asdict.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, asdict.ParentxTable, asdict.ParentxColumn),
		)
		fromU = sqlgraph.SetNeighbors(adq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildrens chains the current query on the "childrens" edge.
func (adq *AsDictQuery) QueryChildrens() *AsDictQuery {
	query := &AsDictQuery{config: adq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := adq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := adq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(asdict.Table, asdict.FieldID, selector),
			sqlgraph.To(asdict.Table, asdict.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, asdict.ChildrensTable, asdict.ChildrensColumn),
		)
		fromU = sqlgraph.SetNeighbors(adq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first AsDict entity from the query.
// Returns a *NotFoundError when no AsDict was found.
func (adq *AsDictQuery) First(ctx context.Context) (*AsDict, error) {
	nodes, err := adq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{asdict.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (adq *AsDictQuery) FirstX(ctx context.Context) *AsDict {
	node, err := adq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first AsDict ID from the query.
// Returns a *NotFoundError when no AsDict ID was found.
func (adq *AsDictQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = adq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{asdict.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (adq *AsDictQuery) FirstIDX(ctx context.Context) int64 {
	id, err := adq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single AsDict entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one AsDict entity is found.
// Returns a *NotFoundError when no AsDict entities are found.
func (adq *AsDictQuery) Only(ctx context.Context) (*AsDict, error) {
	nodes, err := adq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{asdict.Label}
	default:
		return nil, &NotSingularError{asdict.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (adq *AsDictQuery) OnlyX(ctx context.Context) *AsDict {
	node, err := adq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only AsDict ID in the query.
// Returns a *NotSingularError when more than one AsDict ID is found.
// Returns a *NotFoundError when no entities are found.
func (adq *AsDictQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = adq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = &NotSingularError{asdict.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (adq *AsDictQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := adq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of AsDicts.
func (adq *AsDictQuery) All(ctx context.Context) ([]*AsDict, error) {
	if err := adq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return adq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (adq *AsDictQuery) AllX(ctx context.Context) []*AsDict {
	nodes, err := adq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of AsDict IDs.
func (adq *AsDictQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := adq.Select(asdict.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (adq *AsDictQuery) IDsX(ctx context.Context) []int64 {
	ids, err := adq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (adq *AsDictQuery) Count(ctx context.Context) (int64, error) {
	if err := adq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return adq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (adq *AsDictQuery) CountX(ctx context.Context) int64 {
	count, err := adq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (adq *AsDictQuery) Exist(ctx context.Context) (bool, error) {
	if err := adq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return adq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (adq *AsDictQuery) ExistX(ctx context.Context) bool {
	exist, err := adq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the AsDictQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (adq *AsDictQuery) Clone() *AsDictQuery {
	if adq == nil {
		return nil
	}
	return &AsDictQuery{
		config:        adq.config,
		limit:         adq.limit,
		offset:        adq.offset,
		order:         append([]OrderFunc{}, adq.order...),
		predicates:    append([]predicate.AsDict{}, adq.predicates...),
		withParentx:   adq.withParentx.Clone(),
		withChildrens: adq.withChildrens.Clone(),
		// clone intermediate query.
		sql:    adq.sql.Clone(),
		path:   adq.path,
		unique: adq.unique,
	}
}

// WithParentx tells the query-builder to eager-load the nodes that are connected to
// the "parentx" edge. The optional arguments are used to configure the query builder of the edge.
func (adq *AsDictQuery) WithParentx(opts ...func(*AsDictQuery)) *AsDictQuery {
	query := &AsDictQuery{config: adq.config}
	for _, opt := range opts {
		opt(query)
	}
	adq.withParentx = query
	return adq
}

// WithChildrens tells the query-builder to eager-load the nodes that are connected to
// the "childrens" edge. The optional arguments are used to configure the query builder of the edge.
func (adq *AsDictQuery) WithChildrens(opts ...func(*AsDictQuery)) *AsDictQuery {
	query := &AsDictQuery{config: adq.config}
	for _, opt := range opts {
		opt(query)
	}
	adq.withChildrens = query
	return adq
}

// ThenParentx tells the query-builder to eager-load the nodes that are connected to
// the "parentx" edge. The optional arguments are used to configure the query builder of the edge.
func (adq *AsDictQuery) ThenParentx(opts ...func(*AsDictQuery)) *AsDictQuery {
	query := &AsDictQuery{config: adq.config}
	for _, opt := range opts {
		opt(query.Where(asdict.IsDeleted(0)))
	}
	adq.withParentx = query
	return adq
}

// ThenChildrens tells the query-builder to eager-load the nodes that are connected to
// the "childrens" edge. The optional arguments are used to configure the query builder of the edge.
func (adq *AsDictQuery) ThenChildrens(opts ...func(*AsDictQuery)) *AsDictQuery {
	query := &AsDictQuery{config: adq.config}
	for _, opt := range opts {
		opt(query.Where(asdict.IsDeleted(0)))
	}
	adq.withChildrens = query
	return adq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ParentID int64 `json:"parentId,string"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.AsDict.Query().
//		GroupBy(asdict.FieldParentID).
//		Aggregate(schema.Count()).
//		Scan(ctx, &v)
//
func (adq *AsDictQuery) GroupBy(field string, fields ...string) *AsDictGroupBy {
	group := &AsDictGroupBy{config: adq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := adq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return adq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ParentID int64 `json:"parentId,string"`
//	}
//
//	client.AsDict.Query().
//		Select(asdict.FieldParentID).
//		Scan(ctx, &v)
//
func (adq *AsDictQuery) Select(fields ...string) *AsDictSelect {
	adq.fields = append(adq.fields, fields...)
	return &AsDictSelect{AsDictQuery: adq}
}

func (adq *AsDictQuery) prepareQuery(ctx context.Context) error {
	for _, f := range adq.fields {
		if !asdict.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("schema: invalid field %q for query", f)}
		}
	}
	if adq.path != nil {
		prev, err := adq.path(ctx)
		if err != nil {
			return err
		}
		adq.sql = prev
	}
	return nil
}

func (adq *AsDictQuery) sqlAll(ctx context.Context) ([]*AsDict, error) {
	var (
		nodes       = []*AsDict{}
		_spec       = adq.querySpec()
		loadedTypes = [2]bool{
			adq.withParentx != nil,
			adq.withChildrens != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &AsDict{config: adq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("schema: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, adq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}

	if query := adq.withParentx; query != nil {
		ids := make([]int64, 0, len(nodes))
		nodeids := make(map[int64][]*AsDict)
		for i := range nodes {
			fk := nodes[i].ParentID
			if _, ok := nodeids[fk]; !ok {
				ids = append(ids, fk)
			}
			nodeids[fk] = append(nodeids[fk], nodes[i])
		}
		query.Where(asdict.IDIn(ids...))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			nodes, ok := nodeids[n.ID]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v`, n.ID)
			}
			for i := range nodes {
				nodes[i].Edges.Parentx = n
			}
		}
	}

	if query := adq.withChildrens; query != nil {
		fks := make([]driver.Value, 0, len(nodes))
		nodeids := make(map[int64]*AsDict)
		for i := range nodes {
			fks = append(fks, nodes[i].ID)
			nodeids[nodes[i].ID] = nodes[i]
			nodes[i].Edges.Childrens = []*AsDict{}
		}
		query.Where(predicate.AsDict(func(s *sql.Selector) {
			s.Where(sql.InValues(asdict.ChildrensColumn, fks...))
		}))
		neighbors, err := query.All(ctx)
		if err != nil {
			return nil, err
		}
		for _, n := range neighbors {
			fk := n.ParentID
			node, ok := nodeids[fk]
			if !ok {
				return nil, fmt.Errorf(`unexpected foreign-key "parent_id" returned %v for node %v`, fk, n.ID)
			}
			node.Edges.Childrens = append(node.Edges.Childrens, n)
		}
	}

	return nodes, nil
}

func (adq *AsDictQuery) sqlCount(ctx context.Context) (int64, error) {
	_spec := adq.querySpec()
	_spec.Node.Columns = adq.fields
	if len(adq.fields) > 0 {
		_spec.Unique = adq.unique != nil && *adq.unique
	}
	c, err := sqlgraph.CountNodes(ctx, adq.driver, _spec)
	return int64(c), err
}

func (adq *AsDictQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := adq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("schema: check existence: %w", err)
	}
	return n > 0, nil
}

func (adq *AsDictQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   asdict.Table,
			Columns: asdict.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: asdict.FieldID,
			},
		},
		From:   adq.sql,
		Unique: true,
	}
	if unique := adq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := adq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, asdict.FieldID)
		for i := range fields {
			if fields[i] != asdict.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := adq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := adq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := adq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := adq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (adq *AsDictQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(adq.driver.Dialect())
	t1 := builder.Table(asdict.Table)
	columns := adq.fields
	if len(columns) == 0 {
		columns = asdict.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if adq.sql != nil {
		selector = adq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if adq.unique != nil && *adq.unique {
		selector.Distinct()
	}
	for _, p := range adq.predicates {
		p(selector)
	}
	for _, p := range adq.order {
		p(selector)
	}
	if offset := adq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := adq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// AsDictGroupBy is the group-by builder for AsDict entities.
type AsDictGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (adgb *AsDictGroupBy) Aggregate(fns ...AggregateFunc) *AsDictGroupBy {
	adgb.fns = append(adgb.fns, fns...)
	return adgb
}

// Scan applies the group-by query and scans the result into the given value.
func (adgb *AsDictGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := adgb.path(ctx)
	if err != nil {
		return err
	}
	adgb.sql = query
	return adgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (adgb *AsDictGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := adgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(adgb.fields) > 1 {
		return nil, errors.New("schema: AsDictGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := adgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (adgb *AsDictGroupBy) StringsX(ctx context.Context) []string {
	v, err := adgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = adgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (adgb *AsDictGroupBy) StringX(ctx context.Context) string {
	v, err := adgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(adgb.fields) > 1 {
		return nil, errors.New("schema: AsDictGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := adgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (adgb *AsDictGroupBy) IntsX(ctx context.Context) []int {
	v, err := adgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = adgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (adgb *AsDictGroupBy) IntX(ctx context.Context) int {
	v, err := adgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(adgb.fields) > 1 {
		return nil, errors.New("schema: AsDictGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := adgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (adgb *AsDictGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := adgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = adgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (adgb *AsDictGroupBy) Float64X(ctx context.Context) float64 {
	v, err := adgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(adgb.fields) > 1 {
		return nil, errors.New("schema: AsDictGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := adgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (adgb *AsDictGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := adgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = adgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (adgb *AsDictGroupBy) BoolX(ctx context.Context) bool {
	v, err := adgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Int64s(ctx context.Context) ([]int64, error) {
	if len(adgb.fields) > 1 {
		return nil, errors.New("schema: AsDictGroupBy.Int64s is not achievable when grouping more than 1 field")
	}
	var v []int64
	if err := adgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (adgb *AsDictGroupBy) Int64sX(ctx context.Context) []int64 {
	v, err := adgb.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (adgb *AsDictGroupBy) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = adgb.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictGroupBy.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (adgb *AsDictGroupBy) Int64X(ctx context.Context) int64 {
	v, err := adgb.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (adgb *AsDictGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range adgb.fields {
		if !asdict.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := adgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := adgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (adgb *AsDictGroupBy) sqlQuery() *sql.Selector {
	selector := adgb.sql.Select()
	aggregation := make([]string, 0, len(adgb.fns))
	for _, fn := range adgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(adgb.fields)+len(adgb.fns))
		for _, f := range adgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(adgb.fields...)...)
}

// AsDictSelect is the builder for selecting fields of AsDict entities.
type AsDictSelect struct {
	*AsDictQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ads *AsDictSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ads.prepareQuery(ctx); err != nil {
		return err
	}
	ads.sql = ads.AsDictQuery.sqlQuery(ctx)
	return ads.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ads *AsDictSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ads.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ads.fields) > 1 {
		return nil, errors.New("schema: AsDictSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ads.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ads *AsDictSelect) StringsX(ctx context.Context) []string {
	v, err := ads.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ads.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ads *AsDictSelect) StringX(ctx context.Context) string {
	v, err := ads.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ads.fields) > 1 {
		return nil, errors.New("schema: AsDictSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ads.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ads *AsDictSelect) IntsX(ctx context.Context) []int {
	v, err := ads.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ads.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ads *AsDictSelect) IntX(ctx context.Context) int {
	v, err := ads.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ads.fields) > 1 {
		return nil, errors.New("schema: AsDictSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ads.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ads *AsDictSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ads.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ads.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ads *AsDictSelect) Float64X(ctx context.Context) float64 {
	v, err := ads.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ads.fields) > 1 {
		return nil, errors.New("schema: AsDictSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ads.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ads *AsDictSelect) BoolsX(ctx context.Context) []bool {
	v, err := ads.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ads.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ads *AsDictSelect) BoolX(ctx context.Context) bool {
	v, err := ads.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64s returns list of int64s from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Int64s(ctx context.Context) ([]int64, error) {
	if len(ads.fields) > 1 {
		return nil, errors.New("schema: AsDictSelect.Int64s is not achievable when selecting more than 1 field")
	}
	var v []int64
	if err := ads.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Int64sX is like Int64s, but panics if an error occurs.
func (ads *AsDictSelect) Int64sX(ctx context.Context) []int64 {
	v, err := ads.Int64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int64 returns a single int64 from a selector. It is only allowed when selecting one field.
func (ads *AsDictSelect) Int64(ctx context.Context) (_ int64, err error) {
	var v []int64
	if v, err = ads.Int64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{asdict.Label}
	default:
		err = fmt.Errorf("schema: AsDictSelect.Int64s returned %d results when one was expected", len(v))
	}
	return
}

// Int64X is like Int64, but panics if an error occurs.
func (ads *AsDictSelect) Int64X(ctx context.Context) int64 {
	v, err := ads.Int64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ads *AsDictSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ads.sql.Query()
	if err := ads.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
