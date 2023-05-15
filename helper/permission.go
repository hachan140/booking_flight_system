package helper

import (
	"booking-flight-system/ent"
	"booking-flight-system/ent/member"
	"context"
	"errors"
	"fmt"
)

type MemberTypeValidator struct {
	client *ent.Client
}

func NewMemberTypeValidator(client *ent.Client) *MemberTypeValidator {
	return &MemberTypeValidator{client: client}
}

func (m *MemberTypeValidator) OneOf(ctx context.Context, types ...member.MemberType) error {
	email, ok := ctx.Value("user_email").(string)
	if !ok {
		return errors.New("user must be logged in")
	}
	user, err := m.client.Member.Query().Where(member.EmailEQ(email)).Only(ctx)
	if err != nil {
		return err
	}
	for _, t := range types {
		if user.MemberType == t {
			return nil
		}
	}
	return fmt.Errorf("member type must be %v", types)
}
