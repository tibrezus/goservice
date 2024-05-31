package userapp

import (
	"github.com/tiberzus/goservice/business/domain/userbus"
	"github.com/tiberzus/goservice/business/sdk/order"
)

var defaultOrderBy = order.NewBy("user_id", order.ASC)

var orderByFields = map[string]string{
	"user_id": userbus.OrderByID,
	"name":    userbus.OrderByName,
	"email":   userbus.OrderByEmail,
	"roles":   userbus.OrderByRoles,
	"enabled": userbus.OrderByEnabled,
}
