// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/flight"
	"booking-flight-system/ent/plane"
	"booking-flight-system/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaneQuery is the builder for querying Plane entities.
type PlaneQuery struct {
	config
	ctx              *QueryContext
	order            []plane.OrderOption
	inters           []Interceptor
	predicates       []predicate.Plane
	withFlights      *FlightQuery
	modifiers        []func(*sql.Selector)
	loadTotal        []func(context.Context, []*Plane) error
	withNamedFlights map[string]*FlightQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PlaneQuery builder.
func (pq *PlaneQuery) Where(ps ...predicate.Plane) *PlaneQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PlaneQuery) Limit(limit int) *PlaneQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PlaneQuery) Offset(offset int) *PlaneQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PlaneQuery) Unique(unique bool) *PlaneQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PlaneQuery) Order(o ...plane.OrderOption) *PlaneQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryFlights chains the current query on the "flights" edge.
func (pq *PlaneQuery) QueryFlights() *FlightQuery {
	query := (&FlightClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(plane.Table, plane.FieldID, selector),
			sqlgraph.To(flight.Table, flight.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, plane.FlightsTable, plane.FlightsColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Plane entity from the query.
// Returns a *NotFoundError when no Plane was found.
func (pq *PlaneQuery) First(ctx context.Context) (*Plane, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{plane.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PlaneQuery) FirstX(ctx context.Context) *Plane {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Plane ID from the query.
// Returns a *NotFoundError when no Plane ID was found.
func (pq *PlaneQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{plane.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PlaneQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Plane entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Plane entity is found.
// Returns a *NotFoundError when no Plane entities are found.
func (pq *PlaneQuery) Only(ctx context.Context) (*Plane, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{plane.Label}
	default:
		return nil, &NotSingularError{plane.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PlaneQuery) OnlyX(ctx context.Context) *Plane {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Plane ID in the query.
// Returns a *NotSingularError when more than one Plane ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PlaneQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{plane.Label}
	default:
		err = &NotSingularError{plane.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PlaneQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Planes.
func (pq *PlaneQuery) All(ctx context.Context) ([]*Plane, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Plane, *PlaneQuery]()
	return withInterceptors[[]*Plane](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PlaneQuery) AllX(ctx context.Context) []*Plane {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Plane IDs.
func (pq *PlaneQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(plane.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PlaneQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PlaneQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PlaneQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PlaneQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PlaneQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PlaneQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PlaneQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PlaneQuery) Clone() *PlaneQuery {
	if pq == nil {
		return nil
	}
	return &PlaneQuery{
		config:      pq.config,
		ctx:         pq.ctx.Clone(),
		order:       append([]plane.OrderOption{}, pq.order...),
		inters:      append([]Interceptor{}, pq.inters...),
		predicates:  append([]predicate.Plane{}, pq.predicates...),
		withFlights: pq.withFlights.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithFlights tells the query-builder to eager-load the nodes that are connected to
// the "flights" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PlaneQuery) WithFlights(opts ...func(*FlightQuery)) *PlaneQuery {
	query := (&FlightClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withFlights = query
	return pq
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
//	client.Plane.Query().
//		GroupBy(plane.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PlaneQuery) GroupBy(field string, fields ...string) *PlaneGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PlaneGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = plane.Label
	grbuild.scan = grbuild.Scan
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
//	client.Plane.Query().
//		Select(plane.FieldCreatedAt).
//		Scan(ctx, &v)
func (pq *PlaneQuery) Select(fields ...string) *PlaneSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PlaneSelect{PlaneQuery: pq}
	sbuild.label = plane.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PlaneSelect configured with the given aggregations.
func (pq *PlaneQuery) Aggregate(fns ...AggregateFunc) *PlaneSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PlaneQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !plane.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PlaneQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Plane, error) {
	var (
		nodes       = []*Plane{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withFlights != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Plane).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Plane{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withFlights; query != nil {
		if err := pq.loadFlights(ctx, query, nodes,
			func(n *Plane) { n.Edges.Flights = []*Flight{} },
			func(n *Plane, e *Flight) { n.Edges.Flights = append(n.Edges.Flights, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range pq.withNamedFlights {
		if err := pq.loadFlights(ctx, query, nodes,
			func(n *Plane) { n.appendNamedFlights(name) },
			func(n *Plane, e *Flight) { n.appendNamedFlights(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range pq.loadTotal {
		if err := pq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PlaneQuery) loadFlights(ctx context.Context, query *FlightQuery, nodes []*Plane, init func(*Plane), assign func(*Plane, *Flight)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Plane)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(flight.FieldPlaneID)
	}
	query.Where(predicate.Flight(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(plane.FlightsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.PlaneID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "plane_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PlaneQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	if len(pq.modifiers) > 0 {
		_spec.Modifiers = pq.modifiers
	}
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PlaneQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(plane.Table, plane.Columns, sqlgraph.NewFieldSpec(plane.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plane.FieldID)
		for i := range fields {
			if fields[i] != plane.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PlaneQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(plane.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = plane.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedFlights tells the query-builder to eager-load the nodes that are connected to the "flights"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (pq *PlaneQuery) WithNamedFlights(name string, opts ...func(*FlightQuery)) *PlaneQuery {
	query := (&FlightClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if pq.withNamedFlights == nil {
		pq.withNamedFlights = make(map[string]*FlightQuery)
	}
	pq.withNamedFlights[name] = query
	return pq
}

// PlaneGroupBy is the group-by builder for Plane entities.
type PlaneGroupBy struct {
	selector
	build *PlaneQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PlaneGroupBy) Aggregate(fns ...AggregateFunc) *PlaneGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PlaneGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaneQuery, *PlaneGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PlaneGroupBy) sqlScan(ctx context.Context, root *PlaneQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PlaneSelect is the builder for selecting fields of Plane entities.
type PlaneSelect struct {
	*PlaneQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PlaneSelect) Aggregate(fns ...AggregateFunc) *PlaneSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PlaneSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PlaneQuery, *PlaneSelect](ctx, ps.PlaneQuery, ps, ps.inters, v)
}

func (ps *PlaneSelect) sqlScan(ctx context.Context, root *PlaneQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
