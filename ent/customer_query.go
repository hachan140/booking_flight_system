// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/booking"
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/member"
	"booking-flight-system/ent/predicate"
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerQuery is the builder for querying Customer entities.
type CustomerQuery struct {
	config
	ctx                 *QueryContext
	order               []customer.OrderOption
	inters              []Interceptor
	predicates          []predicate.Customer
	withHasMember       *MemberQuery
	withHasBooking      *BookingQuery
	modifiers           []func(*sql.Selector)
	loadTotal           []func(context.Context, []*Customer) error
	withNamedHasBooking map[string]*BookingQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the CustomerQuery builder.
func (cq *CustomerQuery) Where(ps ...predicate.Customer) *CustomerQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CustomerQuery) Limit(limit int) *CustomerQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CustomerQuery) Offset(offset int) *CustomerQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CustomerQuery) Unique(unique bool) *CustomerQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CustomerQuery) Order(o ...customer.OrderOption) *CustomerQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// QueryHasMember chains the current query on the "has_member" edge.
func (cq *CustomerQuery) QueryHasMember() *MemberQuery {
	query := (&MemberClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(member.Table, member.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, customer.HasMemberTable, customer.HasMemberColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryHasBooking chains the current query on the "has_booking" edge.
func (cq *CustomerQuery) QueryHasBooking() *BookingQuery {
	query := (&BookingClient{config: cq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(customer.Table, customer.FieldID, selector),
			sqlgraph.To(booking.Table, booking.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, customer.HasBookingTable, customer.HasBookingColumn),
		)
		fromU = sqlgraph.SetNeighbors(cq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Customer entity from the query.
// Returns a *NotFoundError when no Customer was found.
func (cq *CustomerQuery) First(ctx context.Context) (*Customer, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{customer.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CustomerQuery) FirstX(ctx context.Context) *Customer {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Customer ID from the query.
// Returns a *NotFoundError when no Customer ID was found.
func (cq *CustomerQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{customer.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CustomerQuery) FirstIDX(ctx context.Context) int {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Customer entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Customer entity is found.
// Returns a *NotFoundError when no Customer entities are found.
func (cq *CustomerQuery) Only(ctx context.Context) (*Customer, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{customer.Label}
	default:
		return nil, &NotSingularError{customer.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CustomerQuery) OnlyX(ctx context.Context) *Customer {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Customer ID in the query.
// Returns a *NotSingularError when more than one Customer ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CustomerQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{customer.Label}
	default:
		err = &NotSingularError{customer.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CustomerQuery) OnlyIDX(ctx context.Context) int {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Customers.
func (cq *CustomerQuery) All(ctx context.Context) ([]*Customer, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Customer, *CustomerQuery]()
	return withInterceptors[[]*Customer](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CustomerQuery) AllX(ctx context.Context) []*Customer {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Customer IDs.
func (cq *CustomerQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(customer.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CustomerQuery) IDsX(ctx context.Context) []int {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CustomerQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CustomerQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CustomerQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CustomerQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CustomerQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CustomerQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CustomerQuery) Clone() *CustomerQuery {
	if cq == nil {
		return nil
	}
	return &CustomerQuery{
		config:         cq.config,
		ctx:            cq.ctx.Clone(),
		order:          append([]customer.OrderOption{}, cq.order...),
		inters:         append([]Interceptor{}, cq.inters...),
		predicates:     append([]predicate.Customer{}, cq.predicates...),
		withHasMember:  cq.withHasMember.Clone(),
		withHasBooking: cq.withHasBooking.Clone(),
		// clone intermediate query.
		sql:  cq.sql.Clone(),
		path: cq.path,
	}
}

// WithHasMember tells the query-builder to eager-load the nodes that are connected to
// the "has_member" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithHasMember(opts ...func(*MemberQuery)) *CustomerQuery {
	query := (&MemberClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withHasMember = query
	return cq
}

// WithHasBooking tells the query-builder to eager-load the nodes that are connected to
// the "has_booking" edge. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithHasBooking(opts ...func(*BookingQuery)) *CustomerQuery {
	query := (&BookingClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cq.withHasBooking = query
	return cq
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
//	client.Customer.Query().
//		GroupBy(customer.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CustomerQuery) GroupBy(field string, fields ...string) *CustomerGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CustomerGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = customer.Label
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
//	client.Customer.Query().
//		Select(customer.FieldCreatedAt).
//		Scan(ctx, &v)
func (cq *CustomerQuery) Select(fields ...string) *CustomerSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CustomerSelect{CustomerQuery: cq}
	sbuild.label = customer.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CustomerSelect configured with the given aggregations.
func (cq *CustomerQuery) Aggregate(fns ...AggregateFunc) *CustomerSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CustomerQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	for _, f := range cq.ctx.Fields {
		if !customer.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.sql = prev
	}
	return nil
}

func (cq *CustomerQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Customer, error) {
	var (
		nodes       = []*Customer{}
		_spec       = cq.querySpec()
		loadedTypes = [2]bool{
			cq.withHasMember != nil,
			cq.withHasBooking != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Customer).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Customer{config: cq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cq.withHasMember; query != nil {
		if err := cq.loadHasMember(ctx, query, nodes, nil,
			func(n *Customer, e *Member) { n.Edges.HasMember = e }); err != nil {
			return nil, err
		}
	}
	if query := cq.withHasBooking; query != nil {
		if err := cq.loadHasBooking(ctx, query, nodes,
			func(n *Customer) { n.Edges.HasBooking = []*Booking{} },
			func(n *Customer, e *Booking) { n.Edges.HasBooking = append(n.Edges.HasBooking, e) }); err != nil {
			return nil, err
		}
	}
	for name, query := range cq.withNamedHasBooking {
		if err := cq.loadHasBooking(ctx, query, nodes,
			func(n *Customer) { n.appendNamedHasBooking(name) },
			func(n *Customer, e *Booking) { n.appendNamedHasBooking(name, e) }); err != nil {
			return nil, err
		}
	}
	for i := range cq.loadTotal {
		if err := cq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cq *CustomerQuery) loadHasMember(ctx context.Context, query *MemberQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Member)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Customer)
	for i := range nodes {
		fk := nodes[i].MemberID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(member.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "member_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (cq *CustomerQuery) loadHasBooking(ctx context.Context, query *BookingQuery, nodes []*Customer, init func(*Customer), assign func(*Customer, *Booking)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Customer)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(booking.FieldCustomerID)
	}
	query.Where(predicate.Booking(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(customer.HasBookingColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.CustomerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "customer_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (cq *CustomerQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cq.querySpec()
	if len(cq.modifiers) > 0 {
		_spec.Modifiers = cq.modifiers
	}
	_spec.Node.Columns = cq.ctx.Fields
	if len(cq.ctx.Fields) > 0 {
		_spec.Unique = cq.ctx.Unique != nil && *cq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cq.driver, _spec)
}

func (cq *CustomerQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(customer.Table, customer.Columns, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	_spec.From = cq.sql
	if unique := cq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cq.path != nil {
		_spec.Unique = true
	}
	if fields := cq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customer.FieldID)
		for i := range fields {
			if fields[i] != customer.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cq.withHasMember != nil {
			_spec.Node.AddColumnOnce(customer.FieldMemberID)
		}
	}
	if ps := cq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cq *CustomerQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cq.driver.Dialect())
	t1 := builder.Table(customer.Table)
	columns := cq.ctx.Fields
	if len(columns) == 0 {
		columns = customer.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cq.sql != nil {
		selector = cq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cq.ctx.Unique != nil && *cq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cq.predicates {
		p(selector)
	}
	for _, p := range cq.order {
		p(selector)
	}
	if offset := cq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// WithNamedHasBooking tells the query-builder to eager-load the nodes that are connected to the "has_booking"
// edge with the given name. The optional arguments are used to configure the query builder of the edge.
func (cq *CustomerQuery) WithNamedHasBooking(name string, opts ...func(*BookingQuery)) *CustomerQuery {
	query := (&BookingClient{config: cq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	if cq.withNamedHasBooking == nil {
		cq.withNamedHasBooking = make(map[string]*BookingQuery)
	}
	cq.withNamedHasBooking[name] = query
	return cq
}

// CustomerGroupBy is the group-by builder for Customer entities.
type CustomerGroupBy struct {
	selector
	build *CustomerQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CustomerGroupBy) Aggregate(fns ...AggregateFunc) *CustomerGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CustomerGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerQuery, *CustomerGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CustomerGroupBy) sqlScan(ctx context.Context, root *CustomerQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cgb.fns))
	for _, fn := range cgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cgb.flds)+len(cgb.fns))
		for _, f := range *cgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// CustomerSelect is the builder for selecting fields of Customer entities.
type CustomerSelect struct {
	*CustomerQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CustomerSelect) Aggregate(fns ...AggregateFunc) *CustomerSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CustomerSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CustomerQuery, *CustomerSelect](ctx, cs.CustomerQuery, cs, cs.inters, v)
}

func (cs *CustomerSelect) sqlScan(ctx context.Context, root *CustomerQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cs.fns))
	for _, fn := range cs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
