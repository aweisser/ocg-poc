package dummy

import "github.com/aweisser/ocg-poc/domain/ocg"
import "fmt"

// Repo provides access to dummmy data
type Repo struct {
}

// GetMemberByID retrieves a dummy member by its
func (r *Repo) GetMemberByID(id interface{}) *ocg.Member {
	return &ocg.Member{
		Name: fmt.Sprintf("Johndoe%v", id),
	}
}
