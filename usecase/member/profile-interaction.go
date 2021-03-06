package member

import (
	"github.com/aweisser/ocg-poc/domain/ocg"
)

// ProfileInteraction is a stateless service that retrieves single members by ID.
type ProfileInteraction struct {
	MemberRepo memberRepo
}

type memberRepo interface {
	GetMemberByID(id interface{}) *ocg.Member
}

// LoadByID retrieves a Member bei its ID.
func (pi *ProfileInteraction) LoadByID(id interface{}) (*ocg.Member, bool) {
	m := pi.MemberRepo.GetMemberByID(id)
	found := true
	if m == nil {
		found = false
	}
	return m, found
}
