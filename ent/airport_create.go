// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/airport"
	"booking-flight-system/ent/flight"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// AirportCreate is the builder for creating a Airport entity.
type AirportCreate struct {
	config
	mutation *AirportMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ac *AirportCreate) SetCreatedAt(t time.Time) *AirportCreate {
	ac.mutation.SetCreatedAt(t)
	return ac
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ac *AirportCreate) SetNillableCreatedAt(t *time.Time) *AirportCreate {
	if t != nil {
		ac.SetCreatedAt(*t)
	}
	return ac
}

// SetUpdatedAt sets the "updated_at" field.
func (ac *AirportCreate) SetUpdatedAt(t time.Time) *AirportCreate {
	ac.mutation.SetUpdatedAt(t)
	return ac
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ac *AirportCreate) SetNillableUpdatedAt(t *time.Time) *AirportCreate {
	if t != nil {
		ac.SetUpdatedAt(*t)
	}
	return ac
}

// SetName sets the "name" field.
func (ac *AirportCreate) SetName(s string) *AirportCreate {
	ac.mutation.SetName(s)
	return ac
}

// SetLat sets the "lat" field.
func (ac *AirportCreate) SetLat(f float64) *AirportCreate {
	ac.mutation.SetLat(f)
	return ac
}

// SetLong sets the "long" field.
func (ac *AirportCreate) SetLong(f float64) *AirportCreate {
	ac.mutation.SetLong(f)
	return ac
}

// AddHasFlightIDs adds the "has_flight" edge to the Flight entity by IDs.
func (ac *AirportCreate) AddHasFlightIDs(ids ...int) *AirportCreate {
	ac.mutation.AddHasFlightIDs(ids...)
	return ac
}

// AddHasFlight adds the "has_flight" edges to the Flight entity.
func (ac *AirportCreate) AddHasFlight(f ...*Flight) *AirportCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ac.AddHasFlightIDs(ids...)
}

// Mutation returns the AirportMutation object of the builder.
func (ac *AirportCreate) Mutation() *AirportMutation {
	return ac.mutation
}

// Save creates the Airport in the database.
func (ac *AirportCreate) Save(ctx context.Context) (*Airport, error) {
	ac.defaults()
	return withHooks(ctx, ac.sqlSave, ac.mutation, ac.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ac *AirportCreate) SaveX(ctx context.Context) *Airport {
	v, err := ac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ac *AirportCreate) Exec(ctx context.Context) error {
	_, err := ac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ac *AirportCreate) ExecX(ctx context.Context) {
	if err := ac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ac *AirportCreate) defaults() {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		v := airport.DefaultCreatedAt()
		ac.mutation.SetCreatedAt(v)
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		v := airport.DefaultUpdatedAt()
		ac.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ac *AirportCreate) check() error {
	if _, ok := ac.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Airport.created_at"`)}
	}
	if _, ok := ac.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`ent: missing required field "Airport.updated_at"`)}
	}
	if _, ok := ac.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Airport.name"`)}
	}
	if v, ok := ac.mutation.Name(); ok {
		if err := airport.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Airport.name": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Lat(); !ok {
		return &ValidationError{Name: "lat", err: errors.New(`ent: missing required field "Airport.lat"`)}
	}
	if v, ok := ac.mutation.Lat(); ok {
		if err := airport.LatValidator(v); err != nil {
			return &ValidationError{Name: "lat", err: fmt.Errorf(`ent: validator failed for field "Airport.lat": %w`, err)}
		}
	}
	if _, ok := ac.mutation.Long(); !ok {
		return &ValidationError{Name: "long", err: errors.New(`ent: missing required field "Airport.long"`)}
	}
	if v, ok := ac.mutation.Long(); ok {
		if err := airport.LongValidator(v); err != nil {
			return &ValidationError{Name: "long", err: fmt.Errorf(`ent: validator failed for field "Airport.long": %w`, err)}
		}
	}
	return nil
}

func (ac *AirportCreate) sqlSave(ctx context.Context) (*Airport, error) {
	if err := ac.check(); err != nil {
		return nil, err
	}
	_node, _spec := ac.createSpec()
	if err := sqlgraph.CreateNode(ctx, ac.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	ac.mutation.id = &_node.ID
	ac.mutation.done = true
	return _node, nil
}

func (ac *AirportCreate) createSpec() (*Airport, *sqlgraph.CreateSpec) {
	var (
		_node = &Airport{config: ac.config}
		_spec = sqlgraph.NewCreateSpec(airport.Table, sqlgraph.NewFieldSpec(airport.FieldID, field.TypeInt))
	)
	if value, ok := ac.mutation.CreatedAt(); ok {
		_spec.SetField(airport.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ac.mutation.UpdatedAt(); ok {
		_spec.SetField(airport.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ac.mutation.Name(); ok {
		_spec.SetField(airport.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := ac.mutation.Lat(); ok {
		_spec.SetField(airport.FieldLat, field.TypeFloat64, value)
		_node.Lat = &value
	}
	if value, ok := ac.mutation.Long(); ok {
		_spec.SetField(airport.FieldLong, field.TypeFloat64, value)
		_node.Long = &value
	}
	if nodes := ac.mutation.HasFlightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   airport.HasFlightTable,
			Columns: []string{airport.HasFlightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// AirportCreateBulk is the builder for creating many Airport entities in bulk.
type AirportCreateBulk struct {
	config
	builders []*AirportCreate
}

// Save creates the Airport entities in the database.
func (acb *AirportCreateBulk) Save(ctx context.Context) ([]*Airport, error) {
	specs := make([]*sqlgraph.CreateSpec, len(acb.builders))
	nodes := make([]*Airport, len(acb.builders))
	mutators := make([]Mutator, len(acb.builders))
	for i := range acb.builders {
		func(i int, root context.Context) {
			builder := acb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AirportMutation)
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
					_, err = mutators[i+1].Mutate(root, acb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, acb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, acb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (acb *AirportCreateBulk) SaveX(ctx context.Context) []*Airport {
	v, err := acb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (acb *AirportCreateBulk) Exec(ctx context.Context) error {
	_, err := acb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (acb *AirportCreateBulk) ExecX(ctx context.Context) {
	if err := acb.Exec(ctx); err != nil {
		panic(err)
	}
}
