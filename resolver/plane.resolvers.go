package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"booking-flight-system/ent"
	"booking-flight-system/ent/member"
	graphql1 "booking-flight-system/graphql"
	"context"
	"errors"

	"entgo.io/contrib/entgql"
)

// Plane is the resolver for the plane field.
func (r *mutationResolver) Plane(ctx context.Context) (*ent.PlaneOps, error) {
	return &ent.PlaneOps{}, nil
}

// CreatePlane is the resolver for the create_plane field.
func (r *planeOpsResolver) CreatePlane(ctx context.Context, obj *ent.PlaneOps, input ent.CreatePlaneInput) (*ent.Plane, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	if len(input.Name) == 0 || *input.BusinessClassSlots <= 0 || *input.EconomyClassSlots <= 0 {
		return nil, errors.New("invalid plane input")
	}
	return r.client.Plane.Create().SetInput(input).Save(ctx)
}

// UpdatePlane is the resolver for the update_plane field.
func (r *planeOpsResolver) UpdatePlane(ctx context.Context, obj *ent.PlaneOps, id int, input ent.UpdatePlane) (*ent.Plane, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	if len(*input.Name) == 0 || *input.BusinessClassSlots <= 0 || *input.EconomyClassSlots <= 0 {
		return nil, errors.New("invalid plane input")
	}
	plane, err := r.client.Plane.UpdateOneID(id).SetName(*input.Name).SetStatus(*input.Status).SetBusinessClassSlots(int64(*input.BusinessClassSlots)).SetEconomyClassSlots(int64(*input.EconomyClassSlots)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return plane, nil
}

// DeletePlane is the resolver for the delete_plane field.
func (r *planeOpsResolver) DeletePlane(ctx context.Context, obj *ent.PlaneOps, id int) (*string, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	message := "delete success!"
	if err := r.client.Plane.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}
	return &message, nil
}

// FindPlaneByID is the resolver for the find_plane_by_id field.
func (r *planeOpsResolver) FindPlaneByID(ctx context.Context, obj *ent.PlaneOps, id int) (*ent.Plane, error) {
	if _, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin); err != nil {
		return nil, err
	}
	plane, err := r.client.Plane.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return plane, nil
}

// ListPlanes is the resolver for the list_planes field.
func (r *planeOpsResolver) ListPlanes(ctx context.Context, obj *ent.PlaneOps, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.PlaneOrder) (*ent.PlaneConnection, error) {
	_, err := r.memberTypeValidator.OneOf(ctx, member.MemberTypeAdmin)
	if err != nil {
		return nil, err
	}
	return r.client.Plane.Query().Paginate(ctx, after, first, before, last, ent.WithPlaneOrder(orderBy))
}

// PlaneOps returns graphql1.PlaneOpsResolver implementation.
func (r *Resolver) PlaneOps() graphql1.PlaneOpsResolver { return &planeOpsResolver{r} }

type planeOpsResolver struct{ *Resolver }
