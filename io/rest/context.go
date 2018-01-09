package rest

import (
	"context"

	"github.com/aweisser/ocg-poc/usecase/member"
)

// contextKey should be used as a component specific key type to access values in a context.
type contextKey string

// MemberProfileKey provides access to a member.Profile instance
const MemberProfileKey = contextKey("member.Profile")

// GetMemberProfileInteraction extracts a service instance from the the given Context
// The first return value may be nil, if the service could not be extracted.
// In this case the second return value is false.
func GetMemberProfileInteraction(ctx context.Context) (*member.ProfileInteraction, bool) {
	v, b := ctx.Value(MemberProfileKey).(*member.ProfileInteraction)
	return v, b
}
