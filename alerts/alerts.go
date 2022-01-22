package alerts

import (
	"errors"
	"sync"
)

// Alerts represent all alerts to users
type Alerts struct {
	mtx         sync.Mutex
	usersAlerts map[string][]Alert
}

// ErrUserNotFound is returned when listinf alerts of a non existing user
var ErrUserNotFound = errors.New("User not found")

// NewAlerts create the structure from the alerts and the user it belongs to
func NewAlerts(as map[string][]Alert) Alerts {
	return Alerts{usersAlerts: as}
}

// List alerts that belongs to an user
func (as *Alerts) List(uname string) ([]Alert, error) {
	as.mtx.Lock()
	defer as.mtx.Unlock()

	alerts, ok := as.usersAlerts[uname]

	if !ok {
		return make([]Alert, 0), ErrUserNotFound
	}

	return alerts, nil
}
