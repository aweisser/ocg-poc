package cert

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

// NewCertificate creates a Certificate and ensures that the activity properly set empty
func NewCertificate(category string, activity *Activity) (*Certificate, error) {
	if err := validate(activity); err != nil {
		return nil, err
	}
	uid, err := uuid.NewV4() // TODO remove dependency to uuid lib
	if err != nil {
		return nil, err
	}
	activity.category = category
	c := &Certificate{
		id:       ID(uid),
		activity: *activity,
	}
	return c, nil
}

// ID is a globally unique ID. Perhaps a UUID?
type ID interface{}

// A Certificate certifies an Activity and belongs to an owner (identified by any id)
type Certificate struct {
	id       ID       // Globally unique id. Can only be set at creation time.
	activity Activity // Semi final attribute (immutable after beeing set)
	ownerID  ID       // Semi final attribte (immutable after beeing set)
}

// AssignTo assigns the Certificate to an owner. This can only be executed ONCE!
// So a Certificate can only be assigned once.
func (c *Certificate) AssignTo(ownerID ID) error {
	if c.ownerID == nil {
		c.ownerID = ownerID
		return nil
	}
	return errors.New("The Certificate is already assigned to an owner")
}

// Activity returns a copy of the certified activity,
// so that the private attribute can't be manipluated from outside the package.
func (c *Certificate) Activity() Activity {
	return c.activity
}

// OwnerID returns a copy of the owner ID,
// so that the private attribute can't be manipluated from outside the package.
func (c *Certificate) OwnerID() ID {
	return c.ownerID
}

// Activity that can be included in a Certificate.
type Activity struct {
	From     time.Time
	To       time.Time
	Name     string
	category string // should be set by authorized cert Providers (inside this package)
}

// Category of this Activity
func (a *Activity) Category() string {
	return a.category
}

var emptyTime = time.Time{}

func validate(a *Activity) error {
	if a == nil {
		return errors.New("Activity may not be nil")
	}
	if a.From == emptyTime || a.To == emptyTime {
		return errors.New("Activity must have a dates 'From' and 'To'")
	}
	return nil
}
