// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/booking"
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/member"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// CustomerCreate is the builder for creating a Customer entity.
type CustomerCreate struct {
	config
	mutation *CustomerMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (cc *CustomerCreate) SetCreatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableCreatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetCreatedAt(*t)
	}
	return cc
}

// SetUpdatedAt sets the "updated_at" field.
func (cc *CustomerCreate) SetUpdatedAt(t time.Time) *CustomerCreate {
	cc.mutation.SetUpdatedAt(t)
	return cc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableUpdatedAt(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetUpdatedAt(*t)
	}
	return cc
}

// SetEmail sets the "email" field.
func (cc *CustomerCreate) SetEmail(s string) *CustomerCreate {
	cc.mutation.SetEmail(s)
	return cc
}

// SetPhoneNumber sets the "phone_number" field.
func (cc *CustomerCreate) SetPhoneNumber(s string) *CustomerCreate {
	cc.mutation.SetPhoneNumber(s)
	return cc
}

// SetFullName sets the "full_name" field.
func (cc *CustomerCreate) SetFullName(s string) *CustomerCreate {
	cc.mutation.SetFullName(s)
	return cc
}

// SetDob sets the "dob" field.
func (cc *CustomerCreate) SetDob(t time.Time) *CustomerCreate {
	cc.mutation.SetDob(t)
	return cc
}

// SetNillableDob sets the "dob" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableDob(t *time.Time) *CustomerCreate {
	if t != nil {
		cc.SetDob(*t)
	}
	return cc
}

// SetCid sets the "cid" field.
func (cc *CustomerCreate) SetCid(s string) *CustomerCreate {
	cc.mutation.SetCid(s)
	return cc
}

// SetMemberID sets the "member_id" field.
func (cc *CustomerCreate) SetMemberID(i int) *CustomerCreate {
	cc.mutation.SetMemberID(i)
	return cc
}

// SetNillableMemberID sets the "member_id" field if the given value is not nil.
func (cc *CustomerCreate) SetNillableMemberID(i *int) *CustomerCreate {
	if i != nil {
		cc.SetMemberID(*i)
	}
	return cc
}

// SetHasMemberID sets the "has_member" edge to the Member entity by ID.
func (cc *CustomerCreate) SetHasMemberID(id int) *CustomerCreate {
	cc.mutation.SetHasMemberID(id)
	return cc
}

// SetNillableHasMemberID sets the "has_member" edge to the Member entity by ID if the given value is not nil.
func (cc *CustomerCreate) SetNillableHasMemberID(id *int) *CustomerCreate {
	if id != nil {
		cc = cc.SetHasMemberID(*id)
	}
	return cc
}

// SetHasMember sets the "has_member" edge to the Member entity.
func (cc *CustomerCreate) SetHasMember(m *Member) *CustomerCreate {
	return cc.SetHasMemberID(m.ID)
}

// AddHasBookingIDs adds the "has_booking" edge to the Booking entity by IDs.
func (cc *CustomerCreate) AddHasBookingIDs(ids ...int) *CustomerCreate {
	cc.mutation.AddHasBookingIDs(ids...)
	return cc
}

// AddHasBooking adds the "has_booking" edges to the Booking entity.
func (cc *CustomerCreate) AddHasBooking(b ...*Booking) *CustomerCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return cc.AddHasBookingIDs(ids...)
}

// Mutation returns the CustomerMutation object of the builder.
func (cc *CustomerCreate) Mutation() *CustomerMutation {
	return cc.mutation
}

// Save creates the Customer in the database.
func (cc *CustomerCreate) Save(ctx context.Context) (*Customer, error) {
	cc.defaults()
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CustomerCreate) SaveX(ctx context.Context) *Customer {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CustomerCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CustomerCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cc *CustomerCreate) defaults() {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		v := customer.DefaultCreatedAt()
		cc.mutation.SetCreatedAt(v)
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		v := customer.DefaultUpdatedAt()
		cc.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CustomerCreate) check() error {
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Customer.created_at"`)}
	}
	if _, ok := cc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Customer.updated_at"`)}
	}
	if _, ok := cc.mutation.Email(); !ok {
		return &ValidationError{Name: "email", err: errors.New(`ent: missing required field "Customer.email"`)}
	}
	if v, ok := cc.mutation.Email(); ok {
		if err := customer.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Customer.email": %w`, err)}
		}
	}
	if _, ok := cc.mutation.PhoneNumber(); !ok {
		return &ValidationError{Name: "phone_number", err: errors.New(`ent: missing required field "Customer.phone_number"`)}
	}
	if v, ok := cc.mutation.PhoneNumber(); ok {
		if err := customer.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Customer.phone_number": %w`, err)}
		}
	}
	if _, ok := cc.mutation.FullName(); !ok {
		return &ValidationError{Name: "full_name", err: errors.New(`ent: missing required field "Customer.full_name"`)}
	}
	if v, ok := cc.mutation.FullName(); ok {
		if err := customer.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Customer.full_name": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Cid(); !ok {
		return &ValidationError{Name: "cid", err: errors.New(`ent: missing required field "Customer.cid"`)}
	}
	if v, ok := cc.mutation.Cid(); ok {
		if err := customer.CidValidator(v); err != nil {
			return &ValidationError{Name: "cid", err: fmt.Errorf(`ent: validator failed for field "Customer.cid": %w`, err)}
		}
	}
	return nil
}

func (cc *CustomerCreate) sqlSave(ctx context.Context) (*Customer, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CustomerCreate) createSpec() (*Customer, *sqlgraph.CreateSpec) {
	var (
		_node = &Customer{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(customer.Table, sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt))
	)
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(customer.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := cc.mutation.UpdatedAt(); ok {
		_spec.SetField(customer.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := cc.mutation.Email(); ok {
		_spec.SetField(customer.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := cc.mutation.PhoneNumber(); ok {
		_spec.SetField(customer.FieldPhoneNumber, field.TypeString, value)
		_node.PhoneNumber = value
	}
	if value, ok := cc.mutation.FullName(); ok {
		_spec.SetField(customer.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := cc.mutation.Dob(); ok {
		_spec.SetField(customer.FieldDob, field.TypeTime, value)
		_node.Dob = value
	}
	if value, ok := cc.mutation.Cid(); ok {
		_spec.SetField(customer.FieldCid, field.TypeString, value)
		_node.Cid = value
	}
	if nodes := cc.mutation.HasMemberIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   customer.HasMemberTable,
			Columns: []string{customer.HasMemberColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.MemberID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.HasBookingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   customer.HasBookingTable,
			Columns: []string{customer.HasBookingColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(booking.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CustomerCreateBulk is the builder for creating many Customer entities in bulk.
type CustomerCreateBulk struct {
	config
	builders []*CustomerCreate
}

// Save creates the Customer entities in the database.
func (ccb *CustomerCreateBulk) Save(ctx context.Context) ([]*Customer, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Customer, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CustomerMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CustomerCreateBulk) SaveX(ctx context.Context) []*Customer {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CustomerCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CustomerCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
