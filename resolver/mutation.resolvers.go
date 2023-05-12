package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31

import (
	"booking-flight-system/ent"
	"booking-flight-system/ent/member"
	graphql1 "booking-flight-system/graphql"
	"booking-flight-system/helper"
	"context"
	"errors"
	"fmt"
)

// SignUp is the resolver for the sign_up field.
func (r *mutationResolver) SignUp(ctx context.Context, input ent.CreateMemberInput) (*ent.Member, error) {
	input.Password = helper.SHA256Hashing(input.Password)
	return r.client.Member.Create().SetInput(input).Save(ctx)
}

// Login is the resolver for the login field.
func (r *mutationResolver) Login(ctx context.Context, email string, password string) (*ent.Token, error) {
	hashPass := helper.SHA256Hashing(password)
	user, err := r.client.Member.Query().Where(member.Email(email), member.Password(hashPass)).Only(ctx)
	if err != nil {
		return nil, err
	}
	token, exp, err := r.jwtService.GenerateToken(*user)
	if err != nil {
		return nil, err
	}
	return &ent.Token{
		Token:     token,
		ExpiredAt: exp,
	}, nil
}

// Self is the resolver for the self field.
func (r *mutationResolver) Self(ctx context.Context) (*ent.Member, error) {
	email, ok := ctx.Value("user_email").(string)
	if !ok {
		return nil, errors.New("not logged in")
	}
	user, err := r.client.Member.Query().Where(member.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// DeleteByID is the resolver for the delete_by_id field.
func (r *mutationResolver) DeleteByID(ctx context.Context, id int) (*string, error) {
	message := "delete success"
	if err := r.client.Member.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}
	return &message, nil
}

// FindMemberByName is the resolver for the find_member_by_name field.
func (r *mutationResolver) FindMemberByName(ctx context.Context, name string) ([]*ent.Member, error) {
	panic(fmt.Errorf("not implemented: FindMemberByName - find_member_by_name"))
}

// ChangePassword is the resolver for the change_password field.
func (r *mutationResolver) ChangePassword(ctx context.Context, oldPassword string, newPassword string) (*string, error) {
	message := "success"
	user, err := r.Self(ctx)
	if err != nil {
		return nil, errors.New("can't change password")
	}
	if helper.SHA256Hashing(oldPassword) != user.Password {
		return nil, errors.New("password doesn't match")
	}
	user, err = user.Update().SetPassword(helper.SHA256Hashing(newPassword)).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &message, nil
}

// UpdateMemberProfile is the resolver for the update_member_profile field.
func (r *mutationResolver) UpdateMemberProfile(ctx context.Context, input *ent.UpdateMemberInput) (*ent.Member, error) {
	user, err := r.Self(ctx)
	if err != nil {
		return nil, errors.New("can't update member")
	}
	user, err = user.Update().
		SetFullName(*input.FullName).
		SetCid(*input.Cid).
		SetPhoneNumber(*input.PhoneNumber).
		SetDob(*input.Dob).Save(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateAirport is the resolver for the create_airport field.
func (r *mutationResolver) CreateAirport(ctx context.Context, input ent.CreateAirportInput) (*ent.Airport, error) {
	return r.client.Airport.Create().SetInput(input).Save(ctx)
}

// UpdateAirport is the resolver for the update_airport field.
func (r *mutationResolver) UpdateAirport(ctx context.Context, id int, input ent.UpdateAirportInput) (*ent.Airport, error) {
	airport, err := r.FindAirportByID(ctx, id)
	if err != nil {
		return nil, errors.New("can't find airport")
	}
	airport, err = airport.Update().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}
	return airport, nil
}

// DeleteAirport is the resolver for the delete_airport field.
func (r *mutationResolver) DeleteAirport(ctx context.Context, id int) (*string, error) {
	message := "success"
	if _, err := r.FindAirportByID(ctx, id); err != nil {
		return nil, errors.New("can't find airport")
	}
	if err := r.client.Airport.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}
	return &message, nil
}

// FindAirportByID is the resolver for the find_airport_by_id field.
func (r *mutationResolver) FindAirportByID(ctx context.Context, id int) (*ent.Airport, error) {
	airport, err := r.client.Airport.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return airport, nil
}

// CreatePlane is the resolver for the create_plane field.
func (r *mutationResolver) CreatePlane(ctx context.Context, input ent.CreatePlaneInput) (*ent.Plane, error) {
	return r.client.Plane.Create().SetInput(input).Save(ctx)
}

// UpdatePlane is the resolver for the update_plane field.
func (r *mutationResolver) UpdatePlane(ctx context.Context, id int, input ent.UpdatePlaneInput) (*ent.Plane, error) {
	return nil, nil
}

// DeletePlane is the resolver for the delete_plane field.
func (r *mutationResolver) DeletePlane(ctx context.Context, id int) (*string, error) {
	message := "delete success!"
	if err := r.client.Plane.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}
	return &message, nil
}

// FindPlaneByID is the resolver for the find_plane_by_id field.
func (r *mutationResolver) FindPlaneByID(ctx context.Context, id int) (*ent.Plane, error) {
	plane, err := r.client.Plane.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return plane, nil
}

// Mutation returns graphql1.MutationResolver implementation.
func (r *Resolver) Mutation() graphql1.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
