package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"booking-flight-system/ent"
	"booking-flight-system/ent/airport"
	"booking-flight-system/ent/booking"
	"booking-flight-system/ent/flight"
	"booking-flight-system/ent/member"
	graphql1 "booking-flight-system/graphql"
	"context"
	"errors"

	"entgo.io/contrib/entgql"
)

// CreateFlight is the resolver for the create_flight field.
func (r *flightOpsResolver) CreateFlight(ctx context.Context, obj *ent.FlightOps, input ent.CreateFlight) (*ent.Flight, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	if len(input.Name) == 0 {
		return nil, errors.New("name of flight can't be empty")
	}
	if input.LandAt.Compare(input.DepartAt) <= 0 {
		return nil, errors.New("landing time must be greater than take-off time")
	}
	if input.ToID == input.FromID {
		return nil, errors.New("invalid destination")
	}
	plane, err := r.client.Plane.Get(ctx, input.PlaneID)
	if err != nil {
		return nil, err
	}
	return r.client.Flight.Create().SetName(input.Name).SetDepartAt(input.DepartAt).SetLandAt(input.LandAt).SetToAirportID(input.ToID).SetFromAirportID(input.FromID).SetPlaneID(input.PlaneID).SetAvailableBcSlot(int(plane.BusinessClassSlots)).SetAvailableEcSlot(int(plane.EconomyClassSlots)).Save(ctx)
}

// UpdateFlightStatus is the resolver for the update_flight_status field.
func (r *flightOpsResolver) UpdateFlightStatus(ctx context.Context, obj *ent.FlightOps, id int, input *ent.UpdateFlightStatus) (*ent.Flight, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	f, err := r.FindFlightByID(ctx, obj, id)
	if err != nil {
		return nil, err
	}
	fNew, err := r.client.Flight.UpdateOneID(f.ID).SetStatus(input.Status).Save(ctx)
	if err != nil {
		return nil, err
	}
	return fNew, err
}

// SearchFlight is the resolver for the search_flight field.
func (r *flightOpsResolver) SearchFlight(ctx context.Context, obj *ent.FlightOps, input ent.SearchFlight) ([]*ent.Flight, error) {
	fromAirport, err := r.FindAirportByName(ctx, obj, input.FromAirport)
	if err != nil {
		return nil, err
	}
	toAirport, err := r.FindAirportByName(ctx, obj, input.ToAirport)
	if err != nil {
		return nil, err
	}
	flights, err := r.client.Flight.Query().Where(flight.And(
		flight.FromAirportID(fromAirport.ID),
		flight.ToAirportID(toAirport.ID),
		flight.DepartAtGTE(input.DepartAt),
		flight.StatusEQ(flight.StatusScheduled),
	),
	).All(ctx)
	if err != nil {
		return nil, err
	}
	return flights, nil
}

// UpdateFlight is the resolver for the update_flight field.
func (r *flightOpsResolver) UpdateFlight(ctx context.Context, obj *ent.FlightOps, input ent.UpdateFlight) (*ent.Flight, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	if len(*input.Name) == 0 {
		return nil, errors.New("name of flight can't be empty")
	}
	if input.LandAt.Compare(*input.DepartAt) <= 0 {
		return nil, errors.New("landing time must be greater than take-off time")
	}
	flightRes, err := r.FindFlightByID(ctx, obj, input.ID)
	if err != nil {
		return nil, err
	}
	if _, err := r.client.Plane.Get(ctx, *input.PlaneID); err != nil {
		return nil, errors.New("can't find plane")
	}
	if flightRes.Status == flight.StatusCanceled {
		return nil, errors.New("can't update canceled flight")
	}
	flightRes, err = flightRes.Update().SetDepartAt(*input.DepartAt).SetLandAt(*input.LandAt).SetPlaneID(*input.PlaneID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return flightRes, nil
}

// CancelFlight is the resolver for the cancel_flight field.
func (r *flightOpsResolver) CancelFlight(ctx context.Context, obj *ent.FlightOps, id int) (*string, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	message := "cancel flight successfully"
	f, _ := r.FindFlightByID(ctx, obj, id)
	if _, err := f.Update().SetStatus(flight.StatusCanceled).Save(ctx); err != nil {
		return nil, err
	}
	r.client.Booking.Update().Where(booking.FlightID(id)).SetStatus(booking.StatusCancel)
	return &message, nil
}

// FindAirportByName is the resolver for the find_airport_by_name field.
func (r *flightOpsResolver) FindAirportByName(ctx context.Context, obj *ent.FlightOps, name string) (*ent.Airport, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin, member.MemberTypeMember); err != nil {
		return nil, err
	}
	a, err := r.client.Airport.Query().Where(airport.NameEqualFold(name)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// FindFlightByID is the resolver for the find_flight_by_id field.
func (r *flightOpsResolver) FindFlightByID(ctx context.Context, obj *ent.FlightOps, id int) (*ent.Flight, error) {
	f, err := r.client.Flight.Get(ctx, id)
	if err != nil {
		return nil, errors.New("can't find flight")
	}
	return f, nil
}

// ListFlights is the resolver for the list_flights field.
func (r *flightOpsResolver) ListFlights(ctx context.Context, obj *ent.FlightOps, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.FlightOrder) (*ent.FlightConnection, error) {
	_, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin)
	if err != nil {
		return nil, err
	}
	return r.client.Flight.Query().Paginate(ctx, after, first, before, last, ent.WithFlightOrder(orderBy))
}

// Flight is the resolver for the flight field.
func (r *mutationResolver) Flight(ctx context.Context) (*ent.FlightOps, error) {
	return &ent.FlightOps{}, nil
}

// FlightOps returns graphql1.FlightOpsResolver implementation.
func (r *Resolver) FlightOps() graphql1.FlightOpsResolver { return &flightOpsResolver{r} }

type flightOpsResolver struct{ *Resolver }
