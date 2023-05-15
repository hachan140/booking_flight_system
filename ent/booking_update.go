// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/booking"
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/flight"
	"booking-flight-system/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookingUpdate is the builder for updating Booking entities.
type BookingUpdate struct {
	config
	hooks    []Hook
	mutation *BookingMutation
}

// Where appends a list predicates to the BookingUpdate builder.
func (bu *BookingUpdate) Where(ps ...predicate.Booking) *BookingUpdate {
	bu.mutation.Where(ps...)
	return bu
}

// SetCreatedAt sets the "created_at" field.
func (bu *BookingUpdate) SetCreatedAt(t time.Time) *BookingUpdate {
	bu.mutation.SetCreatedAt(t)
	return bu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableCreatedAt(t *time.Time) *BookingUpdate {
	if t != nil {
		bu.SetCreatedAt(*t)
	}
	return bu
}

// SetUpdatedAt sets the "updated_at" field.
func (bu *BookingUpdate) SetUpdatedAt(t time.Time) *BookingUpdate {
	bu.mutation.SetUpdatedAt(t)
	return bu
}

// SetCode sets the "code" field.
func (bu *BookingUpdate) SetCode(s string) *BookingUpdate {
	bu.mutation.SetCode(s)
	return bu
}

// SetStatus sets the "status" field.
func (bu *BookingUpdate) SetStatus(s string) *BookingUpdate {
	bu.mutation.SetStatus(s)
	return bu
}

// SetCustomerID sets the "customer_id" field.
func (bu *BookingUpdate) SetCustomerID(i int) *BookingUpdate {
	bu.mutation.SetCustomerID(i)
	return bu
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableCustomerID(i *int) *BookingUpdate {
	if i != nil {
		bu.SetCustomerID(*i)
	}
	return bu
}

// ClearCustomerID clears the value of the "customer_id" field.
func (bu *BookingUpdate) ClearCustomerID() *BookingUpdate {
	bu.mutation.ClearCustomerID()
	return bu
}

// SetFlightID sets the "flight_id" field.
func (bu *BookingUpdate) SetFlightID(i int) *BookingUpdate {
	bu.mutation.SetFlightID(i)
	return bu
}

// SetNillableFlightID sets the "flight_id" field if the given value is not nil.
func (bu *BookingUpdate) SetNillableFlightID(i *int) *BookingUpdate {
	if i != nil {
		bu.SetFlightID(*i)
	}
	return bu
}

// ClearFlightID clears the value of the "flight_id" field.
func (bu *BookingUpdate) ClearFlightID() *BookingUpdate {
	bu.mutation.ClearFlightID()
	return bu
}

// SetHasFlightID sets the "has_flight" edge to the Flight entity by ID.
func (bu *BookingUpdate) SetHasFlightID(id int) *BookingUpdate {
	bu.mutation.SetHasFlightID(id)
	return bu
}

// SetNillableHasFlightID sets the "has_flight" edge to the Flight entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableHasFlightID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetHasFlightID(*id)
	}
	return bu
}

// SetHasFlight sets the "has_flight" edge to the Flight entity.
func (bu *BookingUpdate) SetHasFlight(f *Flight) *BookingUpdate {
	return bu.SetHasFlightID(f.ID)
}

// SetHasCustomerID sets the "has_customer" edge to the Customer entity by ID.
func (bu *BookingUpdate) SetHasCustomerID(id int) *BookingUpdate {
	bu.mutation.SetHasCustomerID(id)
	return bu
}

// SetNillableHasCustomerID sets the "has_customer" edge to the Customer entity by ID if the given value is not nil.
func (bu *BookingUpdate) SetNillableHasCustomerID(id *int) *BookingUpdate {
	if id != nil {
		bu = bu.SetHasCustomerID(*id)
	}
	return bu
}

// SetHasCustomer sets the "has_customer" edge to the Customer entity.
func (bu *BookingUpdate) SetHasCustomer(c *Customer) *BookingUpdate {
	return bu.SetHasCustomerID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bu *BookingUpdate) Mutation() *BookingMutation {
	return bu.mutation
}

// ClearHasFlight clears the "has_flight" edge to the Flight entity.
func (bu *BookingUpdate) ClearHasFlight() *BookingUpdate {
	bu.mutation.ClearHasFlight()
	return bu
}

// ClearHasCustomer clears the "has_customer" edge to the Customer entity.
func (bu *BookingUpdate) ClearHasCustomer() *BookingUpdate {
	bu.mutation.ClearHasCustomer()
	return bu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (bu *BookingUpdate) Save(ctx context.Context) (int, error) {
	bu.defaults()
	return withHooks(ctx, bu.sqlSave, bu.mutation, bu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BookingUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BookingUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BookingUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bu *BookingUpdate) defaults() {
	if _, ok := bu.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		bu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bu *BookingUpdate) check() error {
	if v, ok := bu.mutation.Code(); ok {
		if err := booking.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Booking.code": %w`, err)}
		}
	}
	return nil
}

func (bu *BookingUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := bu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	if ps := bu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.CreatedAt(); ok {
		_spec.SetField(booking.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := bu.mutation.Code(); ok {
		_spec.SetField(booking.FieldCode, field.TypeString, value)
	}
	if value, ok := bu.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeString, value)
	}
	if bu.mutation.HasFlightCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasFlightTable,
			Columns: []string{booking.HasFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.HasFlightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasFlightTable,
			Columns: []string{booking.HasFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.HasCustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasCustomerTable,
			Columns: []string{booking.HasCustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.HasCustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasCustomerTable,
			Columns: []string{booking.HasCustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	bu.mutation.done = true
	return n, nil
}

// BookingUpdateOne is the builder for updating a single Booking entity.
type BookingUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BookingMutation
}

// SetCreatedAt sets the "created_at" field.
func (buo *BookingUpdateOne) SetCreatedAt(t time.Time) *BookingUpdateOne {
	buo.mutation.SetCreatedAt(t)
	return buo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableCreatedAt(t *time.Time) *BookingUpdateOne {
	if t != nil {
		buo.SetCreatedAt(*t)
	}
	return buo
}

// SetUpdatedAt sets the "updated_at" field.
func (buo *BookingUpdateOne) SetUpdatedAt(t time.Time) *BookingUpdateOne {
	buo.mutation.SetUpdatedAt(t)
	return buo
}

// SetCode sets the "code" field.
func (buo *BookingUpdateOne) SetCode(s string) *BookingUpdateOne {
	buo.mutation.SetCode(s)
	return buo
}

// SetStatus sets the "status" field.
func (buo *BookingUpdateOne) SetStatus(s string) *BookingUpdateOne {
	buo.mutation.SetStatus(s)
	return buo
}

// SetCustomerID sets the "customer_id" field.
func (buo *BookingUpdateOne) SetCustomerID(i int) *BookingUpdateOne {
	buo.mutation.SetCustomerID(i)
	return buo
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableCustomerID(i *int) *BookingUpdateOne {
	if i != nil {
		buo.SetCustomerID(*i)
	}
	return buo
}

// ClearCustomerID clears the value of the "customer_id" field.
func (buo *BookingUpdateOne) ClearCustomerID() *BookingUpdateOne {
	buo.mutation.ClearCustomerID()
	return buo
}

// SetFlightID sets the "flight_id" field.
func (buo *BookingUpdateOne) SetFlightID(i int) *BookingUpdateOne {
	buo.mutation.SetFlightID(i)
	return buo
}

// SetNillableFlightID sets the "flight_id" field if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableFlightID(i *int) *BookingUpdateOne {
	if i != nil {
		buo.SetFlightID(*i)
	}
	return buo
}

// ClearFlightID clears the value of the "flight_id" field.
func (buo *BookingUpdateOne) ClearFlightID() *BookingUpdateOne {
	buo.mutation.ClearFlightID()
	return buo
}

// SetHasFlightID sets the "has_flight" edge to the Flight entity by ID.
func (buo *BookingUpdateOne) SetHasFlightID(id int) *BookingUpdateOne {
	buo.mutation.SetHasFlightID(id)
	return buo
}

// SetNillableHasFlightID sets the "has_flight" edge to the Flight entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableHasFlightID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetHasFlightID(*id)
	}
	return buo
}

// SetHasFlight sets the "has_flight" edge to the Flight entity.
func (buo *BookingUpdateOne) SetHasFlight(f *Flight) *BookingUpdateOne {
	return buo.SetHasFlightID(f.ID)
}

// SetHasCustomerID sets the "has_customer" edge to the Customer entity by ID.
func (buo *BookingUpdateOne) SetHasCustomerID(id int) *BookingUpdateOne {
	buo.mutation.SetHasCustomerID(id)
	return buo
}

// SetNillableHasCustomerID sets the "has_customer" edge to the Customer entity by ID if the given value is not nil.
func (buo *BookingUpdateOne) SetNillableHasCustomerID(id *int) *BookingUpdateOne {
	if id != nil {
		buo = buo.SetHasCustomerID(*id)
	}
	return buo
}

// SetHasCustomer sets the "has_customer" edge to the Customer entity.
func (buo *BookingUpdateOne) SetHasCustomer(c *Customer) *BookingUpdateOne {
	return buo.SetHasCustomerID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (buo *BookingUpdateOne) Mutation() *BookingMutation {
	return buo.mutation
}

// ClearHasFlight clears the "has_flight" edge to the Flight entity.
func (buo *BookingUpdateOne) ClearHasFlight() *BookingUpdateOne {
	buo.mutation.ClearHasFlight()
	return buo
}

// ClearHasCustomer clears the "has_customer" edge to the Customer entity.
func (buo *BookingUpdateOne) ClearHasCustomer() *BookingUpdateOne {
	buo.mutation.ClearHasCustomer()
	return buo
}

// Where appends a list predicates to the BookingUpdate builder.
func (buo *BookingUpdateOne) Where(ps ...predicate.Booking) *BookingUpdateOne {
	buo.mutation.Where(ps...)
	return buo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (buo *BookingUpdateOne) Select(field string, fields ...string) *BookingUpdateOne {
	buo.fields = append([]string{field}, fields...)
	return buo
}

// Save executes the query and returns the updated Booking entity.
func (buo *BookingUpdateOne) Save(ctx context.Context) (*Booking, error) {
	buo.defaults()
	return withHooks(ctx, buo.sqlSave, buo.mutation, buo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BookingUpdateOne) SaveX(ctx context.Context) *Booking {
	node, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (buo *BookingUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BookingUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (buo *BookingUpdateOne) defaults() {
	if _, ok := buo.mutation.UpdatedAt(); !ok {
		v := booking.UpdateDefaultUpdatedAt()
		buo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (buo *BookingUpdateOne) check() error {
	if v, ok := buo.mutation.Code(); ok {
		if err := booking.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Booking.code": %w`, err)}
		}
	}
	return nil
}

func (buo *BookingUpdateOne) sqlSave(ctx context.Context) (_node *Booking, err error) {
	if err := buo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(booking.Table, booking.Columns, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Booking.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := buo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, booking.FieldID)
		for _, f := range fields {
			if !booking.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != booking.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := buo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := buo.mutation.CreatedAt(); ok {
		_spec.SetField(booking.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := buo.mutation.Code(); ok {
		_spec.SetField(booking.FieldCode, field.TypeString, value)
	}
	if value, ok := buo.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeString, value)
	}
	if buo.mutation.HasFlightCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasFlightTable,
			Columns: []string{booking.HasFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.HasFlightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasFlightTable,
			Columns: []string{booking.HasFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.HasCustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasCustomerTable,
			Columns: []string{booking.HasCustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.HasCustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   booking.HasCustomerTable,
			Columns: []string{booking.HasCustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Booking{config: buo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{booking.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	buo.mutation.done = true
	return _node, nil
}
