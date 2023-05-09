// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-sytem/ent/airport"
	"booking-flight-sytem/ent/flight"
	"booking-flight-sytem/ent/predicate"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AirportUpdate is the builder for updating Airport entities.
type AirportUpdate struct {
	config
	hooks    []Hook
	mutation *AirportMutation
}

// Where appends a list predicates to the AirportUpdate builder.
func (au *AirportUpdate) Where(ps ...predicate.Airport) *AirportUpdate {
	au.mutation.Where(ps...)
	return au
}

// SetCreatedAt sets the "created_at" field.
func (au *AirportUpdate) SetCreatedAt(t time.Time) *AirportUpdate {
	au.mutation.SetCreatedAt(t)
	return au
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (au *AirportUpdate) SetNillableCreatedAt(t *time.Time) *AirportUpdate {
	if t != nil {
		au.SetCreatedAt(*t)
	}
	return au
}

// SetUpdatedAt sets the "updated_at" field.
func (au *AirportUpdate) SetUpdatedAt(t time.Time) *AirportUpdate {
	au.mutation.SetUpdatedAt(t)
	return au
}

// SetName sets the "name" field.
func (au *AirportUpdate) SetName(s string) *AirportUpdate {
	au.mutation.SetName(s)
	return au
}

// SetLat sets the "lat" field.
func (au *AirportUpdate) SetLat(f float64) *AirportUpdate {
	au.mutation.ResetLat()
	au.mutation.SetLat(f)
	return au
}

// AddLat adds f to the "lat" field.
func (au *AirportUpdate) AddLat(f float64) *AirportUpdate {
	au.mutation.AddLat(f)
	return au
}

// SetLong sets the "long" field.
func (au *AirportUpdate) SetLong(f float64) *AirportUpdate {
	au.mutation.ResetLong()
	au.mutation.SetLong(f)
	return au
}

// AddLong adds f to the "long" field.
func (au *AirportUpdate) AddLong(f float64) *AirportUpdate {
	au.mutation.AddLong(f)
	return au
}

// AddFromAirportIDIDs adds the "from_airport_id" edge to the Flight entity by IDs.
func (au *AirportUpdate) AddFromAirportIDIDs(ids ...int) *AirportUpdate {
	au.mutation.AddFromAirportIDIDs(ids...)
	return au
}

// AddFromAirportID adds the "from_airport_id" edges to the Flight entity.
func (au *AirportUpdate) AddFromAirportID(f ...*Flight) *AirportUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.AddFromAirportIDIDs(ids...)
}

// AddDestAirportIDIDs adds the "dest_airport_id" edge to the Flight entity by IDs.
func (au *AirportUpdate) AddDestAirportIDIDs(ids ...int) *AirportUpdate {
	au.mutation.AddDestAirportIDIDs(ids...)
	return au
}

// AddDestAirportID adds the "dest_airport_id" edges to the Flight entity.
func (au *AirportUpdate) AddDestAirportID(f ...*Flight) *AirportUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.AddDestAirportIDIDs(ids...)
}

// Mutation returns the AirportMutation object of the builder.
func (au *AirportUpdate) Mutation() *AirportMutation {
	return au.mutation
}

// ClearFromAirportID clears all "from_airport_id" edges to the Flight entity.
func (au *AirportUpdate) ClearFromAirportID() *AirportUpdate {
	au.mutation.ClearFromAirportID()
	return au
}

// RemoveFromAirportIDIDs removes the "from_airport_id" edge to Flight entities by IDs.
func (au *AirportUpdate) RemoveFromAirportIDIDs(ids ...int) *AirportUpdate {
	au.mutation.RemoveFromAirportIDIDs(ids...)
	return au
}

// RemoveFromAirportID removes "from_airport_id" edges to Flight entities.
func (au *AirportUpdate) RemoveFromAirportID(f ...*Flight) *AirportUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.RemoveFromAirportIDIDs(ids...)
}

// ClearDestAirportID clears all "dest_airport_id" edges to the Flight entity.
func (au *AirportUpdate) ClearDestAirportID() *AirportUpdate {
	au.mutation.ClearDestAirportID()
	return au
}

// RemoveDestAirportIDIDs removes the "dest_airport_id" edge to Flight entities by IDs.
func (au *AirportUpdate) RemoveDestAirportIDIDs(ids ...int) *AirportUpdate {
	au.mutation.RemoveDestAirportIDIDs(ids...)
	return au
}

// RemoveDestAirportID removes "dest_airport_id" edges to Flight entities.
func (au *AirportUpdate) RemoveDestAirportID(f ...*Flight) *AirportUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return au.RemoveDestAirportIDIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (au *AirportUpdate) Save(ctx context.Context) (int, error) {
	au.defaults()
	return withHooks(ctx, au.sqlSave, au.mutation, au.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (au *AirportUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AirportUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AirportUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (au *AirportUpdate) defaults() {
	if _, ok := au.mutation.UpdatedAt(); !ok {
		v := airport.UpdateDefaultUpdatedAt()
		au.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (au *AirportUpdate) check() error {
	if v, ok := au.mutation.Name(); ok {
		if err := airport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Airport.name": %w`, err)}
		}
	}
	if v, ok := au.mutation.Lat(); ok {
		if err := airport.LatValidator(v); err != nil {
			return &ValidationError{Name: "lat", err: fmt.Errorf(`ent: validator failed for field "Airport.lat": %w`, err)}
		}
	}
	if v, ok := au.mutation.Long(); ok {
		if err := airport.LongValidator(v); err != nil {
			return &ValidationError{Name: "long", err: fmt.Errorf(`ent: validator failed for field "Airport.long": %w`, err)}
		}
	}
	return nil
}

func (au *AirportUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := au.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeInt))
	if ps := au.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.CreatedAt(); ok {
		_spec.SetField(airport.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.UpdatedAt(); ok {
		_spec.SetField(airport.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := au.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := au.mutation.Lat(); ok {
		_spec.SetField(airport.FieldLat, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedLat(); ok {
		_spec.AddField(airport.FieldLat, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.Long(); ok {
		_spec.SetField(airport.FieldLong, field.TypeFloat64, value)
	}
	if value, ok := au.mutation.AddedLong(); ok {
		_spec.AddField(airport.FieldLong, field.TypeFloat64, value)
	}
	if au.mutation.FromAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.FromAirportIDTable,
			Columns: []string{airport.FromAirportIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedFromAirportIDIDs(); len(nodes) > 0 && !au.mutation.FromAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.FromAirportIDTable,
			Columns: []string{airport.FromAirportIDColumn},
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
	if nodes := au.mutation.FromAirportIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.FromAirportIDTable,
			Columns: []string{airport.FromAirportIDColumn},
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
	if au.mutation.DestAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestAirportIDTable,
			Columns: []string{airport.DestAirportIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.RemovedDestAirportIDIDs(); len(nodes) > 0 && !au.mutation.DestAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestAirportIDTable,
			Columns: []string{airport.DestAirportIDColumn},
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
	if nodes := au.mutation.DestAirportIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestAirportIDTable,
			Columns: []string{airport.DestAirportIDColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	au.mutation.done = true
	return n, nil
}

// AirportUpdateOne is the builder for updating a single Airport entity.
type AirportUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *AirportMutation
}

// SetCreatedAt sets the "created_at" field.
func (auo *AirportUpdateOne) SetCreatedAt(t time.Time) *AirportUpdateOne {
	auo.mutation.SetCreatedAt(t)
	return auo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (auo *AirportUpdateOne) SetNillableCreatedAt(t *time.Time) *AirportUpdateOne {
	if t != nil {
		auo.SetCreatedAt(*t)
	}
	return auo
}

// SetUpdatedAt sets the "updated_at" field.
func (auo *AirportUpdateOne) SetUpdatedAt(t time.Time) *AirportUpdateOne {
	auo.mutation.SetUpdatedAt(t)
	return auo
}

// SetName sets the "name" field.
func (auo *AirportUpdateOne) SetName(s string) *AirportUpdateOne {
	auo.mutation.SetName(s)
	return auo
}

// SetLat sets the "lat" field.
func (auo *AirportUpdateOne) SetLat(f float64) *AirportUpdateOne {
	auo.mutation.ResetLat()
	auo.mutation.SetLat(f)
	return auo
}

// AddLat adds f to the "lat" field.
func (auo *AirportUpdateOne) AddLat(f float64) *AirportUpdateOne {
	auo.mutation.AddLat(f)
	return auo
}

// SetLong sets the "long" field.
func (auo *AirportUpdateOne) SetLong(f float64) *AirportUpdateOne {
	auo.mutation.ResetLong()
	auo.mutation.SetLong(f)
	return auo
}

// AddLong adds f to the "long" field.
func (auo *AirportUpdateOne) AddLong(f float64) *AirportUpdateOne {
	auo.mutation.AddLong(f)
	return auo
}

// AddFromAirportIDIDs adds the "from_airport_id" edge to the Flight entity by IDs.
func (auo *AirportUpdateOne) AddFromAirportIDIDs(ids ...int) *AirportUpdateOne {
	auo.mutation.AddFromAirportIDIDs(ids...)
	return auo
}

// AddFromAirportID adds the "from_airport_id" edges to the Flight entity.
func (auo *AirportUpdateOne) AddFromAirportID(f ...*Flight) *AirportUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.AddFromAirportIDIDs(ids...)
}

// AddDestAirportIDIDs adds the "dest_airport_id" edge to the Flight entity by IDs.
func (auo *AirportUpdateOne) AddDestAirportIDIDs(ids ...int) *AirportUpdateOne {
	auo.mutation.AddDestAirportIDIDs(ids...)
	return auo
}

// AddDestAirportID adds the "dest_airport_id" edges to the Flight entity.
func (auo *AirportUpdateOne) AddDestAirportID(f ...*Flight) *AirportUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.AddDestAirportIDIDs(ids...)
}

// Mutation returns the AirportMutation object of the builder.
func (auo *AirportUpdateOne) Mutation() *AirportMutation {
	return auo.mutation
}

// ClearFromAirportID clears all "from_airport_id" edges to the Flight entity.
func (auo *AirportUpdateOne) ClearFromAirportID() *AirportUpdateOne {
	auo.mutation.ClearFromAirportID()
	return auo
}

// RemoveFromAirportIDIDs removes the "from_airport_id" edge to Flight entities by IDs.
func (auo *AirportUpdateOne) RemoveFromAirportIDIDs(ids ...int) *AirportUpdateOne {
	auo.mutation.RemoveFromAirportIDIDs(ids...)
	return auo
}

// RemoveFromAirportID removes "from_airport_id" edges to Flight entities.
func (auo *AirportUpdateOne) RemoveFromAirportID(f ...*Flight) *AirportUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.RemoveFromAirportIDIDs(ids...)
}

// ClearDestAirportID clears all "dest_airport_id" edges to the Flight entity.
func (auo *AirportUpdateOne) ClearDestAirportID() *AirportUpdateOne {
	auo.mutation.ClearDestAirportID()
	return auo
}

// RemoveDestAirportIDIDs removes the "dest_airport_id" edge to Flight entities by IDs.
func (auo *AirportUpdateOne) RemoveDestAirportIDIDs(ids ...int) *AirportUpdateOne {
	auo.mutation.RemoveDestAirportIDIDs(ids...)
	return auo
}

// RemoveDestAirportID removes "dest_airport_id" edges to Flight entities.
func (auo *AirportUpdateOne) RemoveDestAirportID(f ...*Flight) *AirportUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return auo.RemoveDestAirportIDIDs(ids...)
}

// Where appends a list predicates to the AirportUpdate builder.
func (auo *AirportUpdateOne) Where(ps ...predicate.Airport) *AirportUpdateOne {
	auo.mutation.Where(ps...)
	return auo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (auo *AirportUpdateOne) Select(field string, fields ...string) *AirportUpdateOne {
	auo.fields = append([]string{field}, fields...)
	return auo
}

// Save executes the query and returns the updated Airport entity.
func (auo *AirportUpdateOne) Save(ctx context.Context) (*Airport, error) {
	auo.defaults()
	return withHooks(ctx, auo.sqlSave, auo.mutation, auo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AirportUpdateOne) SaveX(ctx context.Context) *Airport {
	node, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (auo *AirportUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AirportUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auo *AirportUpdateOne) defaults() {
	if _, ok := auo.mutation.UpdatedAt(); !ok {
		v := airport.UpdateDefaultUpdatedAt()
		auo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auo *AirportUpdateOne) check() error {
	if v, ok := auo.mutation.Name(); ok {
		if err := airport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Airport.name": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Lat(); ok {
		if err := airport.LatValidator(v); err != nil {
			return &ValidationError{Name: "lat", err: fmt.Errorf(`ent: validator failed for field "Airport.lat": %w`, err)}
		}
	}
	if v, ok := auo.mutation.Long(); ok {
		if err := airport.LongValidator(v); err != nil {
			return &ValidationError{Name: "long", err: fmt.Errorf(`ent: validator failed for field "Airport.long": %w`, err)}
		}
	}
	return nil
}

func (auo *AirportUpdateOne) sqlSave(ctx context.Context) (_node *Airport, err error) {
	if err := auo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(airport.Table, airport.Columns, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeInt))
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Airport.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := auo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, airport.FieldID)
		for _, f := range fields {
			if !airport.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != airport.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := auo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := auo.mutation.CreatedAt(); ok {
		_spec.SetField(airport.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.UpdatedAt(); ok {
		_spec.SetField(airport.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := auo.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
	}
	if value, ok := auo.mutation.Lat(); ok {
		_spec.SetField(airport.FieldLat, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedLat(); ok {
		_spec.AddField(airport.FieldLat, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.Long(); ok {
		_spec.SetField(airport.FieldLong, field.TypeFloat64, value)
	}
	if value, ok := auo.mutation.AddedLong(); ok {
		_spec.AddField(airport.FieldLong, field.TypeFloat64, value)
	}
	if auo.mutation.FromAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.FromAirportIDTable,
			Columns: []string{airport.FromAirportIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedFromAirportIDIDs(); len(nodes) > 0 && !auo.mutation.FromAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.FromAirportIDTable,
			Columns: []string{airport.FromAirportIDColumn},
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
	if nodes := auo.mutation.FromAirportIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.FromAirportIDTable,
			Columns: []string{airport.FromAirportIDColumn},
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
	if auo.mutation.DestAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestAirportIDTable,
			Columns: []string{airport.DestAirportIDColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.RemovedDestAirportIDIDs(); len(nodes) > 0 && !auo.mutation.DestAirportIDCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestAirportIDTable,
			Columns: []string{airport.DestAirportIDColumn},
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
	if nodes := auo.mutation.DestAirportIDIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.DestAirportIDTable,
			Columns: []string{airport.DestAirportIDColumn},
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
	_node = &Airport{config: auo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{airport.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	auo.mutation.done = true
	return _node, nil
}
