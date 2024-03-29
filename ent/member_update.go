// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/member"
	"booking-flight-system/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MemberUpdate is the builder for updating Member entities.
type MemberUpdate struct {
	config
	hooks    []Hook
	mutation *MemberMutation
}

// Where appends a list predicates to the MemberUpdate builder.
func (mu *MemberUpdate) Where(ps ...predicate.Member) *MemberUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetCreatedAt sets the "created_at" field.
func (mu *MemberUpdate) SetCreatedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableCreatedAt(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// SetUpdatedAt sets the "updated_at" field.
func (mu *MemberUpdate) SetUpdatedAt(t time.Time) *MemberUpdate {
	mu.mutation.SetUpdatedAt(t)
	return mu
}

// SetEmail sets the "email" field.
func (mu *MemberUpdate) SetEmail(s string) *MemberUpdate {
	mu.mutation.SetEmail(s)
	return mu
}

// SetPassword sets the "password" field.
func (mu *MemberUpdate) SetPassword(s string) *MemberUpdate {
	mu.mutation.SetPassword(s)
	return mu
}

// SetPhoneNumber sets the "phone_number" field.
func (mu *MemberUpdate) SetPhoneNumber(s string) *MemberUpdate {
	mu.mutation.SetPhoneNumber(s)
	return mu
}

// SetFullName sets the "full_name" field.
func (mu *MemberUpdate) SetFullName(s string) *MemberUpdate {
	mu.mutation.SetFullName(s)
	return mu
}

// SetDob sets the "dob" field.
func (mu *MemberUpdate) SetDob(t time.Time) *MemberUpdate {
	mu.mutation.SetDob(t)
	return mu
}

// SetNillableDob sets the "dob" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableDob(t *time.Time) *MemberUpdate {
	if t != nil {
		mu.SetDob(*t)
	}
	return mu
}

// ClearDob clears the value of the "dob" field.
func (mu *MemberUpdate) ClearDob() *MemberUpdate {
	mu.mutation.ClearDob()
	return mu
}

// SetCid sets the "cid" field.
func (mu *MemberUpdate) SetCid(s string) *MemberUpdate {
	mu.mutation.SetCid(s)
	return mu
}

// SetMemberType sets the "member_type" field.
func (mu *MemberUpdate) SetMemberType(mt member.MemberType) *MemberUpdate {
	mu.mutation.SetMemberType(mt)
	return mu
}

// SetNillableMemberType sets the "member_type" field if the given value is not nil.
func (mu *MemberUpdate) SetNillableMemberType(mt *member.MemberType) *MemberUpdate {
	if mt != nil {
		mu.SetMemberType(*mt)
	}
	return mu
}

// ClearMemberType clears the value of the "member_type" field.
func (mu *MemberUpdate) ClearMemberType() *MemberUpdate {
	mu.mutation.ClearMemberType()
	return mu
}

// SetHasCustomerID sets the "has_customer" edge to the Customer entity by ID.
func (mu *MemberUpdate) SetHasCustomerID(id int) *MemberUpdate {
	mu.mutation.SetHasCustomerID(id)
	return mu
}

// SetNillableHasCustomerID sets the "has_customer" edge to the Customer entity by ID if the given value is not nil.
func (mu *MemberUpdate) SetNillableHasCustomerID(id *int) *MemberUpdate {
	if id != nil {
		mu = mu.SetHasCustomerID(*id)
	}
	return mu
}

// SetHasCustomer sets the "has_customer" edge to the Customer entity.
func (mu *MemberUpdate) SetHasCustomer(c *Customer) *MemberUpdate {
	return mu.SetHasCustomerID(c.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (mu *MemberUpdate) Mutation() *MemberMutation {
	return mu.mutation
}

// ClearHasCustomer clears the "has_customer" edge to the Customer entity.
func (mu *MemberUpdate) ClearHasCustomer() *MemberUpdate {
	mu.mutation.ClearHasCustomer()
	return mu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MemberUpdate) Save(ctx context.Context) (int, error) {
	mu.defaults()
	return withHooks(ctx, mu.sqlSave, mu.mutation, mu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MemberUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MemberUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MemberUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mu *MemberUpdate) defaults() {
	if _, ok := mu.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		mu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *MemberUpdate) check() error {
	if v, ok := mu.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	if v, ok := mu.mutation.PhoneNumber(); ok {
		if err := member.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Member.phone_number": %w`, err)}
		}
	}
	if v, ok := mu.mutation.FullName(); ok {
		if err := member.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Member.full_name": %w`, err)}
		}
	}
	if v, ok := mu.mutation.MemberType(); ok {
		if err := member.MemberTypeValidator(v); err != nil {
			return &ValidationError{Name: "member_type", err: fmt.Errorf(`ent: validator failed for field "Member.member_type": %w`, err)}
		}
	}
	return nil
}

func (mu *MemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := mu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.SetField(member.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := mu.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := mu.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := mu.mutation.PhoneNumber(); ok {
		_spec.SetField(member.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := mu.mutation.FullName(); ok {
		_spec.SetField(member.FieldFullName, field.TypeString, value)
	}
	if value, ok := mu.mutation.Dob(); ok {
		_spec.SetField(member.FieldDob, field.TypeTime, value)
	}
	if mu.mutation.DobCleared() {
		_spec.ClearField(member.FieldDob, field.TypeTime)
	}
	if value, ok := mu.mutation.Cid(); ok {
		_spec.SetField(member.FieldCid, field.TypeString, value)
	}
	if value, ok := mu.mutation.MemberType(); ok {
		_spec.SetField(member.FieldMemberType, field.TypeEnum, value)
	}
	if mu.mutation.MemberTypeCleared() {
		_spec.ClearField(member.FieldMemberType, field.TypeEnum)
	}
	if mu.mutation.HasCustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.HasCustomerTable,
			Columns: []string{member.HasCustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.HasCustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.HasCustomerTable,
			Columns: []string{member.HasCustomerColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	mu.mutation.done = true
	return n, nil
}

// MemberUpdateOne is the builder for updating a single Member entity.
type MemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MemberMutation
}

// SetCreatedAt sets the "created_at" field.
func (muo *MemberUpdateOne) SetCreatedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableCreatedAt(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// SetUpdatedAt sets the "updated_at" field.
func (muo *MemberUpdateOne) SetUpdatedAt(t time.Time) *MemberUpdateOne {
	muo.mutation.SetUpdatedAt(t)
	return muo
}

// SetEmail sets the "email" field.
func (muo *MemberUpdateOne) SetEmail(s string) *MemberUpdateOne {
	muo.mutation.SetEmail(s)
	return muo
}

// SetPassword sets the "password" field.
func (muo *MemberUpdateOne) SetPassword(s string) *MemberUpdateOne {
	muo.mutation.SetPassword(s)
	return muo
}

// SetPhoneNumber sets the "phone_number" field.
func (muo *MemberUpdateOne) SetPhoneNumber(s string) *MemberUpdateOne {
	muo.mutation.SetPhoneNumber(s)
	return muo
}

// SetFullName sets the "full_name" field.
func (muo *MemberUpdateOne) SetFullName(s string) *MemberUpdateOne {
	muo.mutation.SetFullName(s)
	return muo
}

// SetDob sets the "dob" field.
func (muo *MemberUpdateOne) SetDob(t time.Time) *MemberUpdateOne {
	muo.mutation.SetDob(t)
	return muo
}

// SetNillableDob sets the "dob" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableDob(t *time.Time) *MemberUpdateOne {
	if t != nil {
		muo.SetDob(*t)
	}
	return muo
}

// ClearDob clears the value of the "dob" field.
func (muo *MemberUpdateOne) ClearDob() *MemberUpdateOne {
	muo.mutation.ClearDob()
	return muo
}

// SetCid sets the "cid" field.
func (muo *MemberUpdateOne) SetCid(s string) *MemberUpdateOne {
	muo.mutation.SetCid(s)
	return muo
}

// SetMemberType sets the "member_type" field.
func (muo *MemberUpdateOne) SetMemberType(mt member.MemberType) *MemberUpdateOne {
	muo.mutation.SetMemberType(mt)
	return muo
}

// SetNillableMemberType sets the "member_type" field if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableMemberType(mt *member.MemberType) *MemberUpdateOne {
	if mt != nil {
		muo.SetMemberType(*mt)
	}
	return muo
}

// ClearMemberType clears the value of the "member_type" field.
func (muo *MemberUpdateOne) ClearMemberType() *MemberUpdateOne {
	muo.mutation.ClearMemberType()
	return muo
}

// SetHasCustomerID sets the "has_customer" edge to the Customer entity by ID.
func (muo *MemberUpdateOne) SetHasCustomerID(id int) *MemberUpdateOne {
	muo.mutation.SetHasCustomerID(id)
	return muo
}

// SetNillableHasCustomerID sets the "has_customer" edge to the Customer entity by ID if the given value is not nil.
func (muo *MemberUpdateOne) SetNillableHasCustomerID(id *int) *MemberUpdateOne {
	if id != nil {
		muo = muo.SetHasCustomerID(*id)
	}
	return muo
}

// SetHasCustomer sets the "has_customer" edge to the Customer entity.
func (muo *MemberUpdateOne) SetHasCustomer(c *Customer) *MemberUpdateOne {
	return muo.SetHasCustomerID(c.ID)
}

// Mutation returns the MemberMutation object of the builder.
func (muo *MemberUpdateOne) Mutation() *MemberMutation {
	return muo.mutation
}

// ClearHasCustomer clears the "has_customer" edge to the Customer entity.
func (muo *MemberUpdateOne) ClearHasCustomer() *MemberUpdateOne {
	muo.mutation.ClearHasCustomer()
	return muo
}

// Where appends a list predicates to the MemberUpdate builder.
func (muo *MemberUpdateOne) Where(ps ...predicate.Member) *MemberUpdateOne {
	muo.mutation.Where(ps...)
	return muo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MemberUpdateOne) Select(field string, fields ...string) *MemberUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Member entity.
func (muo *MemberUpdateOne) Save(ctx context.Context) (*Member, error) {
	muo.defaults()
	return withHooks(ctx, muo.sqlSave, muo.mutation, muo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MemberUpdateOne) SaveX(ctx context.Context) *Member {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MemberUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MemberUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (muo *MemberUpdateOne) defaults() {
	if _, ok := muo.mutation.UpdatedAt(); !ok {
		v := member.UpdateDefaultUpdatedAt()
		muo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *MemberUpdateOne) check() error {
	if v, ok := muo.mutation.Email(); ok {
		if err := member.EmailValidator(v); err != nil {
			return &ValidationError{Name: "email", err: fmt.Errorf(`ent: validator failed for field "Member.email": %w`, err)}
		}
	}
	if v, ok := muo.mutation.PhoneNumber(); ok {
		if err := member.PhoneNumberValidator(v); err != nil {
			return &ValidationError{Name: "phone_number", err: fmt.Errorf(`ent: validator failed for field "Member.phone_number": %w`, err)}
		}
	}
	if v, ok := muo.mutation.FullName(); ok {
		if err := member.FullNameValidator(v); err != nil {
			return &ValidationError{Name: "full_name", err: fmt.Errorf(`ent: validator failed for field "Member.full_name": %w`, err)}
		}
	}
	if v, ok := muo.mutation.MemberType(); ok {
		if err := member.MemberTypeValidator(v); err != nil {
			return &ValidationError{Name: "member_type", err: fmt.Errorf(`ent: validator failed for field "Member.member_type": %w`, err)}
		}
	}
	return nil
}

func (muo *MemberUpdateOne) sqlSave(ctx context.Context) (_node *Member, err error) {
	if err := muo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(member.Table, member.Columns, sqlgraph.NewFieldSpec(member.FieldID, field.TypeInt))
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Member.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, member.FieldID)
		for _, f := range fields {
			if !member.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != member.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.SetField(member.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.UpdatedAt(); ok {
		_spec.SetField(member.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := muo.mutation.Email(); ok {
		_spec.SetField(member.FieldEmail, field.TypeString, value)
	}
	if value, ok := muo.mutation.Password(); ok {
		_spec.SetField(member.FieldPassword, field.TypeString, value)
	}
	if value, ok := muo.mutation.PhoneNumber(); ok {
		_spec.SetField(member.FieldPhoneNumber, field.TypeString, value)
	}
	if value, ok := muo.mutation.FullName(); ok {
		_spec.SetField(member.FieldFullName, field.TypeString, value)
	}
	if value, ok := muo.mutation.Dob(); ok {
		_spec.SetField(member.FieldDob, field.TypeTime, value)
	}
	if muo.mutation.DobCleared() {
		_spec.ClearField(member.FieldDob, field.TypeTime)
	}
	if value, ok := muo.mutation.Cid(); ok {
		_spec.SetField(member.FieldCid, field.TypeString, value)
	}
	if value, ok := muo.mutation.MemberType(); ok {
		_spec.SetField(member.FieldMemberType, field.TypeEnum, value)
	}
	if muo.mutation.MemberTypeCleared() {
		_spec.ClearField(member.FieldMemberType, field.TypeEnum)
	}
	if muo.mutation.HasCustomerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.HasCustomerTable,
			Columns: []string{member.HasCustomerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(customer.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.HasCustomerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   member.HasCustomerTable,
			Columns: []string{member.HasCustomerColumn},
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
	_node = &Member{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{member.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	muo.mutation.done = true
	return _node, nil
}
