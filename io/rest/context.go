package rest

import (
	"context"

	"github.com/aweisser/ocg-poc/usecase/member"
)

// contextKey should be used as a component specific key type to access values in a context.
type contextKey string

// MemberByIDKey provides access to a XYZService instance
const MemberByIDKey = contextKey("member.ByID")

// GetMemberByID extracts a service instance from the the given Context
// The first return value may be nil, if the service could not be extracted.
// In this case the second return value is false.
func MemberByID(ctx context.Context) (*member.ByID, bool) {
	v, b := ctx.Value(MemberByID).(*member.ByID)
	return v, b
}
