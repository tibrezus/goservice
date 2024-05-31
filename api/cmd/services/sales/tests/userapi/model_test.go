package user_test

import (
	"time"

	"github.com/tiberzus/goservice/app/domain/userapp"
	"github.com/tiberzus/goservice/business/domain/userbus"
)

func toAppUser(bus userbus.User) userapp.User {
	roles := make([]string, len(bus.Roles))
	for i, role := range bus.Roles {
		roles[i] = role.String()
	}

	return userapp.User{
		ID:           bus.ID.String(),
		Name:         bus.Name.String(),
		Email:        bus.Email.Address,
		Roles:        roles,
		PasswordHash: nil, // This field is not marshalled.
		Department:   bus.Department,
		Enabled:      bus.Enabled,
		DateCreated:  bus.DateCreated.Format(time.RFC3339),
		DateUpdated:  bus.DateUpdated.Format(time.RFC3339),
	}
}

func toAppUsers(users []userbus.User) []userapp.User {
	items := make([]userapp.User, len(users))
	for i, usr := range users {
		items[i] = toAppUser(usr)
	}

	return items
}

func toAppUserPtr(bus userbus.User) *userapp.User {
	appUsr := toAppUser(bus)
	return &appUsr
}
