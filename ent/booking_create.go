// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/booking"
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/flight"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// BookingCreate is the builder for creating a Booking entity.
type BookingCreate struct {
	config
	mutation *BookingMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (bc *BookingCreate) SetCreatedAt(t time.Time) *BookingCreate {
	bc.mutation.SetCreatedAt(t)
	return bc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (bc *BookingCreate) SetNillableCreatedAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetCreatedAt(*t)
	}
	return bc
}

// SetUpdatedAt sets the "updated_at" field.
func (bc *BookingCreate) SetUpdatedAt(t time.Time) *BookingCreate {
	bc.mutation.SetUpdatedAt(t)
	return bc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (bc *BookingCreate) SetNillableUpdatedAt(t *time.Time) *BookingCreate {
	if t != nil {
		bc.SetUpdatedAt(*t)
	}
	return bc
}

// SetCode sets the "code" field.
func (bc *BookingCreate) SetCode(s string) *BookingCreate {
	bc.mutation.SetCode(s)
	return bc
}

// SetStatus sets the "status" field.
func (bc *BookingCreate) SetStatus(b booking.Status) *BookingCreate {
	bc.mutation.SetStatus(b)
	return bc
}

// SetSeatType sets the "seat_type" field.
func (bc *BookingCreate) SetSeatType(bt booking.SeatType) *BookingCreate {
	bc.mutation.SetSeatType(bt)
	return bc
}

// SetCustomerID sets the "customer_id" field.
func (bc *BookingCreate) SetCustomerID(i int) *BookingCreate {
	bc.mutation.SetCustomerID(i)
	return bc
}

// SetNillableCustomerID sets the "customer_id" field if the given value is not nil.
func (bc *BookingCreate) SetNillableCustomerID(i *int) *BookingCreate {
	if i != nil {
		bc.SetCustomerID(*i)
	}
	return bc
}

// SetFlightID sets the "flight_id" field.
func (bc *BookingCreate) SetFlightID(i int) *BookingCreate {
	bc.mutation.SetFlightID(i)
	return bc
}

// SetNillableFlightID sets the "flight_id" field if the given value is not nil.
func (bc *BookingCreate) SetNillableFlightID(i *int) *BookingCreate {
	if i != nil {
		bc.SetFlightID(*i)
	}
	return bc
}

// SetHasFlightID sets the "has_flight" edge to the Flight entity by ID.
func (bc *BookingCreate) SetHasFlightID(id int) *BookingCreate {
	bc.mutation.SetHasFlightID(id)
	return bc
}

// SetNillableHasFlightID sets the "has_flight" edge to the Flight entity by ID if the given value is not nil.
func (bc *BookingCreate) SetNillableHasFlightID(id *int) *BookingCreate {
	if id != nil {
		bc = bc.SetHasFlightID(*id)
	}
	return bc
}

// SetHasFlight sets the "has_flight" edge to the Flight entity.
func (bc *BookingCreate) SetHasFlight(f *Flight) *BookingCreate {
	return bc.SetHasFlightID(f.ID)
}

// SetHasCustomerID sets the "has_customer" edge to the Customer entity by ID.
func (bc *BookingCreate) SetHasCustomerID(id int) *BookingCreate {
	bc.mutation.SetHasCustomerID(id)
	return bc
}

// SetNillableHasCustomerID sets the "has_customer" edge to the Customer entity by ID if the given value is not nil.
func (bc *BookingCreate) SetNillableHasCustomerID(id *int) *BookingCreate {
	if id != nil {
		bc = bc.SetHasCustomerID(*id)
	}
	return bc
}

// SetHasCustomer sets the "has_customer" edge to the Customer entity.
func (bc *BookingCreate) SetHasCustomer(c *Customer) *BookingCreate {
	return bc.SetHasCustomerID(c.ID)
}

// Mutation returns the BookingMutation object of the builder.
func (bc *BookingCreate) Mutation() *BookingMutation {
	return bc.mutation
}

// Save creates the Booking in the database.
func (bc *BookingCreate) Save(ctx context.Context) (*Booking, error) {
	bc.defaults()
	return withHooks(ctx, bc.sqlSave, bc.mutation, bc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BookingCreate) SaveX(ctx context.Context) *Booking {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bc *BookingCreate) Exec(ctx context.Context) error {
	_, err := bc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bc *BookingCreate) ExecX(ctx context.Context) {
	if err := bc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (bc *BookingCreate) defaults() {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		v := booking.DefaultCreatedAt()
		bc.mutation.SetCreatedAt(v)
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		v := booking.DefaultUpdatedAt()
		bc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (bc *BookingCreate) check() error {
	if _, ok := bc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Booking.created_at"`)}
	}
	if _, ok := bc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Booking.updated_at"`)}
	}
	if _, ok := bc.mutation.Code(); !ok {
		return &ValidationError{Name: "code", err: errors.New(`ent: missing required field "Booking.code"`)}
	}
	if v, ok := bc.mutation.Code(); ok {
		if err := booking.CodeValidator(v); err != nil {
			return &ValidationError{Name: "code", err: fmt.Errorf(`ent: validator failed for field "Booking.code": %w`, err)}
		}
	}
	if _, ok := bc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Booking.status"`)}
	}
	if v, ok := bc.mutation.Status(); ok {
		if err := booking.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Booking.status": %w`, err)}
		}
	}
	if _, ok := bc.mutation.SeatType(); !ok {
		return &ValidationError{Name: "seat_type", err: errors.New(`ent: missing required field "Booking.seat_type"`)}
	}
	if v, ok := bc.mutation.SeatType(); ok {
		if err := booking.SeatTypeValidator(v); err != nil {
			return &ValidationError{Name: "seat_type", err: fmt.Errorf(`ent: validator failed for field "Booking.seat_type": %w`, err)}
		}
	}
	return nil
}

func (bc *BookingCreate) sqlSave(ctx context.Context) (*Booking, error) {
	if err := bc.check(); err != nil {
		return nil, err
	}
	_node, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	bc.mutation.id = &_node.ID
	bc.mutation.done = true
	return _node, nil
}

func (bc *BookingCreate) createSpec() (*Booking, *sqlgraph.CreateSpec) {
	var (
		_node = &Booking{config: bc.config}
		_spec = sqlgraph.NewCreateSpec(booking.Table, sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt))
	)
	if value, ok := bc.mutation.CreatedAt(); ok {
		_spec.SetField(booking.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := bc.mutation.UpdatedAt(); ok {
		_spec.SetField(booking.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := bc.mutation.Code(); ok {
		_spec.SetField(booking.FieldCode, field.TypeString, value)
		_node.Code = value
	}
	if value, ok := bc.mutation.Status(); ok {
		_spec.SetField(booking.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := bc.mutation.SeatType(); ok {
		_spec.SetField(booking.FieldSeatType, field.TypeEnum, value)
		_node.SeatType = value
	}
	if nodes := bc.mutation.HasFlightIDs(); len(nodes) > 0 {
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
		_node.FlightID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.HasCustomerIDs(); len(nodes) > 0 {
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
		_node.CustomerID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// BookingCreateBulk is the builder for creating many Booking entities in bulk.
type BookingCreateBulk struct {
	config
	builders []*BookingCreate
}

// Save creates the Booking entities in the database.
func (bcb *BookingCreateBulk) Save(ctx context.Context) ([]*Booking, error) {
	specs := make([]*sqlgraph.CreateSpec, len(bcb.builders))
	nodes := make([]*Booking, len(bcb.builders))
	mutators := make([]Mutator, len(bcb.builders))
	for i := range bcb.builders {
		func(i int, root context.Context) {
			builder := bcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*BookingMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, bcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, bcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, bcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (bcb *BookingCreateBulk) SaveX(ctx context.Context) []*Booking {
	v, err := bcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (bcb *BookingCreateBulk) Exec(ctx context.Context) error {
	_, err := bcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bcb *BookingCreateBulk) ExecX(ctx context.Context) {
	if err := bcb.Exec(ctx); err != nil {
		panic(err)
	}
}
