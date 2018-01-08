package member

import (
	"fmt"

	"github.com/aweisser/ocg-poc/domain/ocg"
)

// ByID is a stateless service that retrieves single members by ID.
type ByID struct {
}

// Get retrieves a Member bei its ID.
func (b *ByID) Get(id interface{}) *ocg.Member {
	return &ocg.Member{
		Name: fmt.Sprintf("johndoe%v", id),
	}
}
