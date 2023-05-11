// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/flight"
	"booking-flight-system/ent/plane"
	"booking-flight-system/ent/predicate"
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PlaneUpdate is the builder for updating Plane entities.
type PlaneUpdate struct {
	config
	hooks    []Hook
	mutation *PlaneMutation
}

// Where appends a list predicates to the PlaneUpdate builder.
func (pu *PlaneUpdate) Where(ps ...predicate.Plane) *PlaneUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetName sets the "name" field.
func (pu *PlaneUpdate) SetName(s string) *PlaneUpdate {
	pu.mutation.SetName(s)
	return pu
}

// SetEconomyClassSlots sets the "economy_class_slots" field.
func (pu *PlaneUpdate) SetEconomyClassSlots(i int64) *PlaneUpdate {
	pu.mutation.ResetEconomyClassSlots()
	pu.mutation.SetEconomyClassSlots(i)
	return pu
}

// AddEconomyClassSlots adds i to the "economy_class_slots" field.
func (pu *PlaneUpdate) AddEconomyClassSlots(i int64) *PlaneUpdate {
	pu.mutation.AddEconomyClassSlots(i)
	return pu
}

// SetBusinessClassSlots sets the "business_class_slots" field.
func (pu *PlaneUpdate) SetBusinessClassSlots(i int64) *PlaneUpdate {
	pu.mutation.ResetBusinessClassSlots()
	pu.mutation.SetBusinessClassSlots(i)
	return pu
}

// AddBusinessClassSlots adds i to the "business_class_slots" field.
func (pu *PlaneUpdate) AddBusinessClassSlots(i int64) *PlaneUpdate {
	pu.mutation.AddBusinessClassSlots(i)
	return pu
}

// SetStatus sets the "status" field.
func (pu *PlaneUpdate) SetStatus(pl plane.Status) *PlaneUpdate {
	pu.mutation.SetStatus(pl)
	return pu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (pu *PlaneUpdate) SetNillableStatus(pl *plane.Status) *PlaneUpdate {
	if pl != nil {
		pu.SetStatus(*pl)
	}
	return pu
}

// AddFlightIDs adds the "flights" edge to the Flight entity by IDs.
func (pu *PlaneUpdate) AddFlightIDs(ids ...int) *PlaneUpdate {
	pu.mutation.AddFlightIDs(ids...)
	return pu
}

// AddFlights adds the "flights" edges to the Flight entity.
func (pu *PlaneUpdate) AddFlights(f ...*Flight) *PlaneUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.AddFlightIDs(ids...)
}

// Mutation returns the PlaneMutation object of the builder.
func (pu *PlaneUpdate) Mutation() *PlaneMutation {
	return pu.mutation
}

// ClearFlights clears all "flights" edges to the Flight entity.
func (pu *PlaneUpdate) ClearFlights() *PlaneUpdate {
	pu.mutation.ClearFlights()
	return pu
}

// RemoveFlightIDs removes the "flights" edge to Flight entities by IDs.
func (pu *PlaneUpdate) RemoveFlightIDs(ids ...int) *PlaneUpdate {
	pu.mutation.RemoveFlightIDs(ids...)
	return pu
}

// RemoveFlights removes "flights" edges to Flight entities.
func (pu *PlaneUpdate) RemoveFlights(f ...*Flight) *PlaneUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return pu.RemoveFlightIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PlaneUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pu.sqlSave, pu.mutation, pu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PlaneUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PlaneUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PlaneUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *PlaneUpdate) check() error {
	if v, ok := pu.mutation.Name(); ok {
		if err := plane.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Plane.name": %w`, err)}
		}
	}
	if v, ok := pu.mutation.Status(); ok {
		if err := plane.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Plane.status": %w`, err)}
		}
	}
	return nil
}

func (pu *PlaneUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := pu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(plane.Table, plane.Columns, sqlgraph.NewFieldSpec(plane.FieldID, field.TypeInt))
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Name(); ok {
		_spec.SetField(plane.FieldName, field.TypeString, value)
	}
	if value, ok := pu.mutation.EconomyClassSlots(); ok {
		_spec.SetField(plane.FieldEconomyClassSlots, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedEconomyClassSlots(); ok {
		_spec.AddField(plane.FieldEconomyClassSlots, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.BusinessClassSlots(); ok {
		_spec.SetField(plane.FieldBusinessClassSlots, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.AddedBusinessClassSlots(); ok {
		_spec.AddField(plane.FieldBusinessClassSlots, field.TypeInt64, value)
	}
	if value, ok := pu.mutation.Status(); ok {
		_spec.SetField(plane.FieldStatus, field.TypeEnum, value)
	}
	if pu.mutation.FlightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plane.FlightsTable,
			Columns: []string{plane.FlightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.RemovedFlightsIDs(); len(nodes) > 0 && !pu.mutation.FlightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plane.FlightsTable,
			Columns: []string{plane.FlightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.FlightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plane.FlightsTable,
			Columns: []string{plane.FlightsColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plane.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pu.mutation.done = true
	return n, nil
}

// PlaneUpdateOne is the builder for updating a single Plane entity.
type PlaneUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PlaneMutation
}

// SetName sets the "name" field.
func (puo *PlaneUpdateOne) SetName(s string) *PlaneUpdateOne {
	puo.mutation.SetName(s)
	return puo
}

// SetEconomyClassSlots sets the "economy_class_slots" field.
func (puo *PlaneUpdateOne) SetEconomyClassSlots(i int64) *PlaneUpdateOne {
	puo.mutation.ResetEconomyClassSlots()
	puo.mutation.SetEconomyClassSlots(i)
	return puo
}

// AddEconomyClassSlots adds i to the "economy_class_slots" field.
func (puo *PlaneUpdateOne) AddEconomyClassSlots(i int64) *PlaneUpdateOne {
	puo.mutation.AddEconomyClassSlots(i)
	return puo
}

// SetBusinessClassSlots sets the "business_class_slots" field.
func (puo *PlaneUpdateOne) SetBusinessClassSlots(i int64) *PlaneUpdateOne {
	puo.mutation.ResetBusinessClassSlots()
	puo.mutation.SetBusinessClassSlots(i)
	return puo
}

// AddBusinessClassSlots adds i to the "business_class_slots" field.
func (puo *PlaneUpdateOne) AddBusinessClassSlots(i int64) *PlaneUpdateOne {
	puo.mutation.AddBusinessClassSlots(i)
	return puo
}

// SetStatus sets the "status" field.
func (puo *PlaneUpdateOne) SetStatus(pl plane.Status) *PlaneUpdateOne {
	puo.mutation.SetStatus(pl)
	return puo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (puo *PlaneUpdateOne) SetNillableStatus(pl *plane.Status) *PlaneUpdateOne {
	if pl != nil {
		puo.SetStatus(*pl)
	}
	return puo
}

// AddFlightIDs adds the "flights" edge to the Flight entity by IDs.
func (puo *PlaneUpdateOne) AddFlightIDs(ids ...int) *PlaneUpdateOne {
	puo.mutation.AddFlightIDs(ids...)
	return puo
}

// AddFlights adds the "flights" edges to the Flight entity.
func (puo *PlaneUpdateOne) AddFlights(f ...*Flight) *PlaneUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.AddFlightIDs(ids...)
}

// Mutation returns the PlaneMutation object of the builder.
func (puo *PlaneUpdateOne) Mutation() *PlaneMutation {
	return puo.mutation
}

// ClearFlights clears all "flights" edges to the Flight entity.
func (puo *PlaneUpdateOne) ClearFlights() *PlaneUpdateOne {
	puo.mutation.ClearFlights()
	return puo
}

// RemoveFlightIDs removes the "flights" edge to Flight entities by IDs.
func (puo *PlaneUpdateOne) RemoveFlightIDs(ids ...int) *PlaneUpdateOne {
	puo.mutation.RemoveFlightIDs(ids...)
	return puo
}

// RemoveFlights removes "flights" edges to Flight entities.
func (puo *PlaneUpdateOne) RemoveFlights(f ...*Flight) *PlaneUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return puo.RemoveFlightIDs(ids...)
}

// Where appends a list predicates to the PlaneUpdate builder.
func (puo *PlaneUpdateOne) Where(ps ...predicate.Plane) *PlaneUpdateOne {
	puo.mutation.Where(ps...)
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PlaneUpdateOne) Select(field string, fields ...string) *PlaneUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Plane entity.
func (puo *PlaneUpdateOne) Save(ctx context.Context) (*Plane, error) {
	return withHooks(ctx, puo.sqlSave, puo.mutation, puo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PlaneUpdateOne) SaveX(ctx context.Context) *Plane {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PlaneUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PlaneUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *PlaneUpdateOne) check() error {
	if v, ok := puo.mutation.Name(); ok {
		if err := plane.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Plane.name": %w`, err)}
		}
	}
	if v, ok := puo.mutation.Status(); ok {
		if err := plane.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Plane.status": %w`, err)}
		}
	}
	return nil
}

func (puo *PlaneUpdateOne) sqlSave(ctx context.Context) (_node *Plane, err error) {
	if err := puo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(plane.Table, plane.Columns, sqlgraph.NewFieldSpec(plane.FieldID, field.TypeInt))
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Plane.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, plane.FieldID)
		for _, f := range fields {
			if !plane.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != plane.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Name(); ok {
		_spec.SetField(plane.FieldName, field.TypeString, value)
	}
	if value, ok := puo.mutation.EconomyClassSlots(); ok {
		_spec.SetField(plane.FieldEconomyClassSlots, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedEconomyClassSlots(); ok {
		_spec.AddField(plane.FieldEconomyClassSlots, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.BusinessClassSlots(); ok {
		_spec.SetField(plane.FieldBusinessClassSlots, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.AddedBusinessClassSlots(); ok {
		_spec.AddField(plane.FieldBusinessClassSlots, field.TypeInt64, value)
	}
	if value, ok := puo.mutation.Status(); ok {
		_spec.SetField(plane.FieldStatus, field.TypeEnum, value)
	}
	if puo.mutation.FlightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plane.FlightsTable,
			Columns: []string{plane.FlightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.RemovedFlightsIDs(); len(nodes) > 0 && !puo.mutation.FlightsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plane.FlightsTable,
			Columns: []string{plane.FlightsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.FlightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   plane.FlightsTable,
			Columns: []string{plane.FlightsColumn},
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
	_node = &Plane{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{plane.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	puo.mutation.done = true
	return _node, nil
}
