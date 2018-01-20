package cert

import (
	"errors"
	"time"

	"github.com/satori/go.uuid"
)

// The first return value might be nil in case of an error. The caller has to check!
func NewActivityCertificate(category string, activity *Activity) (*ActivityCertificate, error) {
	// This should be the only exported factory function
	if err := assertActivityNotEmpty(activity); err != nil {
		return nil, err
	}
	activity.category = category
	return newActivityCertificate(activity)
}

func assertActivityNotEmpty(a *Activity) error {
	if a == nil {
		return errors.New("Activity may not be nil")
	}
	emptyTime := time.Time{}
	if a.From == emptyTime || a.To == emptyTime {
		return errors.New("Activity must have a dates 'From' and 'To'")
	}
	return nil
}

func newActivityCertificate(activity *Activity) (*ActivityCertificate, error) {
	uid, err := uuid.NewV4() // TODO remove dependency
	if err != nil {
		return nil, err
	}
	c := &ActivityCertificate{
		id:       ID(uid),
		activity: *activity,
	}
	return c, nil
}

// ID is a globaly unique ID. Perhaps a UUID?
type ID interface{}

type ActivityCertificate struct {
	id       ID       // Globally unique id. Can only be set at creation time.
	activity Activity // Semi final attribute (immutable after beeing set)
	ownerID  ID       // Semi final attribte (immutable after beeing set)
}

// An ActivityCertificate can only be assigned once
func (c *ActivityCertificate) AssignTo(ownerID ID) error {
	if c.ownerID == nil {
		c.ownerID = ownerID
		return nil
	}
	return errors.New("The Certificate is already assigned to an owner")
}

func (c *ActivityCertificate) Activity() Activity {
	return c.activity
}

func (c *ActivityCertificate) OwnerID() ID {
	return c.ownerID
}

type Activity struct {
	From     time.Time
	To       time.Time
	Name     string
	category string // should be set by authorized cert Providers (inside this package)
}

func (a *Activity) Category() string {
	return a.category
}
