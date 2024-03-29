// Code generated by ent, DO NOT EDIT.

package ent

import (
	"booking-flight-system/ent/airport"
	"booking-flight-system/ent/booking"
	"booking-flight-system/ent/customer"
	"booking-flight-system/ent/flight"
	"booking-flight-system/ent/member"
	"booking-flight-system/ent/plane"
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *AirportQuery) CollectFields(ctx context.Context, satisfies ...string) (*AirportQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return a, nil
	}
	if err := a.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return a, nil
}

func (a *AirportQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(airport.Columns))
		selectedFields = []string{airport.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "fromFlight":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FlightClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedFromFlight(alias, func(wq *FlightQuery) {
				*wq = *query
			})
		case "toFlight":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FlightClient{config: a.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			a.WithNamedToFlight(alias, func(wq *FlightQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[airport.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, airport.FieldCreatedAt)
				fieldSeen[airport.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[airport.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, airport.FieldUpdatedAt)
				fieldSeen[airport.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[airport.FieldName]; !ok {
				selectedFields = append(selectedFields, airport.FieldName)
				fieldSeen[airport.FieldName] = struct{}{}
			}
		case "lat":
			if _, ok := fieldSeen[airport.FieldLat]; !ok {
				selectedFields = append(selectedFields, airport.FieldLat)
				fieldSeen[airport.FieldLat] = struct{}{}
			}
		case "long":
			if _, ok := fieldSeen[airport.FieldLong]; !ok {
				selectedFields = append(selectedFields, airport.FieldLong)
				fieldSeen[airport.FieldLong] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		a.Select(selectedFields...)
	}
	return nil
}

type airportPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []AirportPaginateOption
}

func newAirportPaginateArgs(rv map[string]any) *airportPaginateArgs {
	args := &airportPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &AirportOrder{Field: &AirportOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithAirportOrder(order))
			}
		case *AirportOrder:
			if v != nil {
				args.opts = append(args.opts, WithAirportOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*AirportWhereInput); ok {
		args.opts = append(args.opts, WithAirportFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (b *BookingQuery) CollectFields(ctx context.Context, satisfies ...string) (*BookingQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return b, nil
	}
	if err := b.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *BookingQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(booking.Columns))
		selectedFields = []string{booking.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "hasFlight":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FlightClient{config: b.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			b.withHasFlight = query
			if _, ok := fieldSeen[booking.FieldFlightID]; !ok {
				selectedFields = append(selectedFields, booking.FieldFlightID)
				fieldSeen[booking.FieldFlightID] = struct{}{}
			}
		case "hasCustomer":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CustomerClient{config: b.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			b.withHasCustomer = query
			if _, ok := fieldSeen[booking.FieldCustomerID]; !ok {
				selectedFields = append(selectedFields, booking.FieldCustomerID)
				fieldSeen[booking.FieldCustomerID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[booking.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, booking.FieldCreatedAt)
				fieldSeen[booking.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[booking.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, booking.FieldUpdatedAt)
				fieldSeen[booking.FieldUpdatedAt] = struct{}{}
			}
		case "code":
			if _, ok := fieldSeen[booking.FieldCode]; !ok {
				selectedFields = append(selectedFields, booking.FieldCode)
				fieldSeen[booking.FieldCode] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[booking.FieldStatus]; !ok {
				selectedFields = append(selectedFields, booking.FieldStatus)
				fieldSeen[booking.FieldStatus] = struct{}{}
			}
		case "seatType":
			if _, ok := fieldSeen[booking.FieldSeatType]; !ok {
				selectedFields = append(selectedFields, booking.FieldSeatType)
				fieldSeen[booking.FieldSeatType] = struct{}{}
			}
		case "isRound":
			if _, ok := fieldSeen[booking.FieldIsRound]; !ok {
				selectedFields = append(selectedFields, booking.FieldIsRound)
				fieldSeen[booking.FieldIsRound] = struct{}{}
			}
		case "customerID":
			if _, ok := fieldSeen[booking.FieldCustomerID]; !ok {
				selectedFields = append(selectedFields, booking.FieldCustomerID)
				fieldSeen[booking.FieldCustomerID] = struct{}{}
			}
		case "flightID":
			if _, ok := fieldSeen[booking.FieldFlightID]; !ok {
				selectedFields = append(selectedFields, booking.FieldFlightID)
				fieldSeen[booking.FieldFlightID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		b.Select(selectedFields...)
	}
	return nil
}

type bookingPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []BookingPaginateOption
}

func newBookingPaginateArgs(rv map[string]any) *bookingPaginateArgs {
	args := &bookingPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &BookingOrder{Field: &BookingOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithBookingOrder(order))
			}
		case *BookingOrder:
			if v != nil {
				args.opts = append(args.opts, WithBookingOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*BookingWhereInput); ok {
		args.opts = append(args.opts, WithBookingFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (c *CustomerQuery) CollectFields(ctx context.Context, satisfies ...string) (*CustomerQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return c, nil
	}
	if err := c.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *CustomerQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(customer.Columns))
		selectedFields = []string{customer.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "hasMember":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&MemberClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.withHasMember = query
			if _, ok := fieldSeen[customer.FieldMemberID]; !ok {
				selectedFields = append(selectedFields, customer.FieldMemberID)
				fieldSeen[customer.FieldMemberID] = struct{}{}
			}
		case "hasBooking":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&BookingClient{config: c.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			c.WithNamedHasBooking(alias, func(wq *BookingQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[customer.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, customer.FieldCreatedAt)
				fieldSeen[customer.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[customer.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, customer.FieldUpdatedAt)
				fieldSeen[customer.FieldUpdatedAt] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[customer.FieldEmail]; !ok {
				selectedFields = append(selectedFields, customer.FieldEmail)
				fieldSeen[customer.FieldEmail] = struct{}{}
			}
		case "phoneNumber":
			if _, ok := fieldSeen[customer.FieldPhoneNumber]; !ok {
				selectedFields = append(selectedFields, customer.FieldPhoneNumber)
				fieldSeen[customer.FieldPhoneNumber] = struct{}{}
			}
		case "fullName":
			if _, ok := fieldSeen[customer.FieldFullName]; !ok {
				selectedFields = append(selectedFields, customer.FieldFullName)
				fieldSeen[customer.FieldFullName] = struct{}{}
			}
		case "dob":
			if _, ok := fieldSeen[customer.FieldDob]; !ok {
				selectedFields = append(selectedFields, customer.FieldDob)
				fieldSeen[customer.FieldDob] = struct{}{}
			}
		case "cid":
			if _, ok := fieldSeen[customer.FieldCid]; !ok {
				selectedFields = append(selectedFields, customer.FieldCid)
				fieldSeen[customer.FieldCid] = struct{}{}
			}
		case "memberID":
			if _, ok := fieldSeen[customer.FieldMemberID]; !ok {
				selectedFields = append(selectedFields, customer.FieldMemberID)
				fieldSeen[customer.FieldMemberID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		c.Select(selectedFields...)
	}
	return nil
}

type customerPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CustomerPaginateOption
}

func newCustomerPaginateArgs(rv map[string]any) *customerPaginateArgs {
	args := &customerPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &CustomerOrder{Field: &CustomerOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithCustomerOrder(order))
			}
		case *CustomerOrder:
			if v != nil {
				args.opts = append(args.opts, WithCustomerOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*CustomerWhereInput); ok {
		args.opts = append(args.opts, WithCustomerFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FlightQuery) CollectFields(ctx context.Context, satisfies ...string) (*FlightQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return f, nil
	}
	if err := f.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FlightQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(flight.Columns))
		selectedFields = []string{flight.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "hasPlane":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&PlaneClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			f.withHasPlane = query
			if _, ok := fieldSeen[flight.FieldPlaneID]; !ok {
				selectedFields = append(selectedFields, flight.FieldPlaneID)
				fieldSeen[flight.FieldPlaneID] = struct{}{}
			}
		case "hasBooking":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&BookingClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			f.WithNamedHasBooking(alias, func(wq *BookingQuery) {
				*wq = *query
			})
		case "fromAirport":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AirportClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			f.withFromAirport = query
			if _, ok := fieldSeen[flight.FieldFromAirportID]; !ok {
				selectedFields = append(selectedFields, flight.FieldFromAirportID)
				fieldSeen[flight.FieldFromAirportID] = struct{}{}
			}
		case "toAirport":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&AirportClient{config: f.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			f.withToAirport = query
			if _, ok := fieldSeen[flight.FieldToAirportID]; !ok {
				selectedFields = append(selectedFields, flight.FieldToAirportID)
				fieldSeen[flight.FieldToAirportID] = struct{}{}
			}
		case "createdAt":
			if _, ok := fieldSeen[flight.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, flight.FieldCreatedAt)
				fieldSeen[flight.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[flight.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, flight.FieldUpdatedAt)
				fieldSeen[flight.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[flight.FieldName]; !ok {
				selectedFields = append(selectedFields, flight.FieldName)
				fieldSeen[flight.FieldName] = struct{}{}
			}
		case "departAt":
			if _, ok := fieldSeen[flight.FieldDepartAt]; !ok {
				selectedFields = append(selectedFields, flight.FieldDepartAt)
				fieldSeen[flight.FieldDepartAt] = struct{}{}
			}
		case "landAt":
			if _, ok := fieldSeen[flight.FieldLandAt]; !ok {
				selectedFields = append(selectedFields, flight.FieldLandAt)
				fieldSeen[flight.FieldLandAt] = struct{}{}
			}
		case "availableEcSlot":
			if _, ok := fieldSeen[flight.FieldAvailableEcSlot]; !ok {
				selectedFields = append(selectedFields, flight.FieldAvailableEcSlot)
				fieldSeen[flight.FieldAvailableEcSlot] = struct{}{}
			}
		case "availableBcSlot":
			if _, ok := fieldSeen[flight.FieldAvailableBcSlot]; !ok {
				selectedFields = append(selectedFields, flight.FieldAvailableBcSlot)
				fieldSeen[flight.FieldAvailableBcSlot] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[flight.FieldStatus]; !ok {
				selectedFields = append(selectedFields, flight.FieldStatus)
				fieldSeen[flight.FieldStatus] = struct{}{}
			}
		case "planeID":
			if _, ok := fieldSeen[flight.FieldPlaneID]; !ok {
				selectedFields = append(selectedFields, flight.FieldPlaneID)
				fieldSeen[flight.FieldPlaneID] = struct{}{}
			}
		case "fromAirportID":
			if _, ok := fieldSeen[flight.FieldFromAirportID]; !ok {
				selectedFields = append(selectedFields, flight.FieldFromAirportID)
				fieldSeen[flight.FieldFromAirportID] = struct{}{}
			}
		case "toAirportID":
			if _, ok := fieldSeen[flight.FieldToAirportID]; !ok {
				selectedFields = append(selectedFields, flight.FieldToAirportID)
				fieldSeen[flight.FieldToAirportID] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		f.Select(selectedFields...)
	}
	return nil
}

type flightPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []FlightPaginateOption
}

func newFlightPaginateArgs(rv map[string]any) *flightPaginateArgs {
	args := &flightPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &FlightOrder{Field: &FlightOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithFlightOrder(order))
			}
		case *FlightOrder:
			if v != nil {
				args.opts = append(args.opts, WithFlightOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*FlightWhereInput); ok {
		args.opts = append(args.opts, WithFlightFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (m *MemberQuery) CollectFields(ctx context.Context, satisfies ...string) (*MemberQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return m, nil
	}
	if err := m.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return m, nil
}

func (m *MemberQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(member.Columns))
		selectedFields = []string{member.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "hasCustomer":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&CustomerClient{config: m.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			m.withHasCustomer = query
		case "createdAt":
			if _, ok := fieldSeen[member.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, member.FieldCreatedAt)
				fieldSeen[member.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[member.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, member.FieldUpdatedAt)
				fieldSeen[member.FieldUpdatedAt] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[member.FieldEmail]; !ok {
				selectedFields = append(selectedFields, member.FieldEmail)
				fieldSeen[member.FieldEmail] = struct{}{}
			}
		case "phoneNumber":
			if _, ok := fieldSeen[member.FieldPhoneNumber]; !ok {
				selectedFields = append(selectedFields, member.FieldPhoneNumber)
				fieldSeen[member.FieldPhoneNumber] = struct{}{}
			}
		case "fullName":
			if _, ok := fieldSeen[member.FieldFullName]; !ok {
				selectedFields = append(selectedFields, member.FieldFullName)
				fieldSeen[member.FieldFullName] = struct{}{}
			}
		case "dob":
			if _, ok := fieldSeen[member.FieldDob]; !ok {
				selectedFields = append(selectedFields, member.FieldDob)
				fieldSeen[member.FieldDob] = struct{}{}
			}
		case "cid":
			if _, ok := fieldSeen[member.FieldCid]; !ok {
				selectedFields = append(selectedFields, member.FieldCid)
				fieldSeen[member.FieldCid] = struct{}{}
			}
		case "memberType":
			if _, ok := fieldSeen[member.FieldMemberType]; !ok {
				selectedFields = append(selectedFields, member.FieldMemberType)
				fieldSeen[member.FieldMemberType] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		m.Select(selectedFields...)
	}
	return nil
}

type memberPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []MemberPaginateOption
}

func newMemberPaginateArgs(rv map[string]any) *memberPaginateArgs {
	args := &memberPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &MemberOrder{Field: &MemberOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithMemberOrder(order))
			}
		case *MemberOrder:
			if v != nil {
				args.opts = append(args.opts, WithMemberOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*MemberWhereInput); ok {
		args.opts = append(args.opts, WithMemberFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pl *PlaneQuery) CollectFields(ctx context.Context, satisfies ...string) (*PlaneQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pl, nil
	}
	if err := pl.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pl, nil
}

func (pl *PlaneQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(plane.Columns))
		selectedFields = []string{plane.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "flights":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FlightClient{config: pl.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			pl.WithNamedFlights(alias, func(wq *FlightQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[plane.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, plane.FieldCreatedAt)
				fieldSeen[plane.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[plane.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, plane.FieldUpdatedAt)
				fieldSeen[plane.FieldUpdatedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[plane.FieldName]; !ok {
				selectedFields = append(selectedFields, plane.FieldName)
				fieldSeen[plane.FieldName] = struct{}{}
			}
		case "economyClassSlots":
			if _, ok := fieldSeen[plane.FieldEconomyClassSlots]; !ok {
				selectedFields = append(selectedFields, plane.FieldEconomyClassSlots)
				fieldSeen[plane.FieldEconomyClassSlots] = struct{}{}
			}
		case "businessClassSlots":
			if _, ok := fieldSeen[plane.FieldBusinessClassSlots]; !ok {
				selectedFields = append(selectedFields, plane.FieldBusinessClassSlots)
				fieldSeen[plane.FieldBusinessClassSlots] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[plane.FieldStatus]; !ok {
				selectedFields = append(selectedFields, plane.FieldStatus)
				fieldSeen[plane.FieldStatus] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pl.Select(selectedFields...)
	}
	return nil
}

type planePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PlanePaginateOption
}

func newPlanePaginateArgs(rv map[string]any) *planePaginateArgs {
	args := &planePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &PlaneOrder{Field: &PlaneOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPlaneOrder(order))
			}
		case *PlaneOrder:
			if v != nil {
				args.opts = append(args.opts, WithPlaneOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PlaneWhereInput); ok {
		args.opts = append(args.opts, WithPlaneFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
