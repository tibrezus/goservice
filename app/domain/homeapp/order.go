package homeapp

import (
	"github.com/tiberzus/goservice/business/domain/homebus"
	"github.com/tiberzus/goservice/business/sdk/order"
)

var defaultOrderBy = order.NewBy("home_id", order.ASC)

var orderByFields = map[string]string{
	"home_id": homebus.OrderByID,
	"type":    homebus.OrderByType,
	"user_id": homebus.OrderByUserID,
}
