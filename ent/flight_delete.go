// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/flight"
	"booking-flight-system/ent/predicate"
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FlightDelete is the builder for deleting a Flight entity.
type FlightDelete struct {
	config
	hooks    []Hook
	mutation *FlightMutation
}

// Where appends a list predicates to the FlightDelete builder.
func (fd *FlightDelete) Where(ps ...predicate.Flight) *FlightDelete {
	fd.mutation.Where(ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FlightDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, fd.sqlExec, fd.mutation, fd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FlightDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FlightDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(flight.Table, sqlgraph.NewFieldSpec(flight.FieldID, field.TypeInt))
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fd.mutation.done = true
	return affected, err
}

// FlightDeleteOne is the builder for deleting a single Flight entity.
type FlightDeleteOne struct {
	fd *FlightDelete
}

// Where appends a list predicates to the FlightDelete builder.
func (fdo *FlightDeleteOne) Where(ps ...predicate.Flight) *FlightDeleteOne {
	fdo.fd.mutation.Where(ps...)
	return fdo
}

// Exec executes the deletion query.
func (fdo *FlightDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{flight.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FlightDeleteOne) ExecX(ctx context.Context) {
	if err := fdo.Exec(ctx); err != nil {
		panic(err)
	}
}
