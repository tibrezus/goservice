package vproductbus

import (
	"github.com/google/uuid"
	"github.com/tiberzus/goservice/business/domain/productbus"
	"github.com/tiberzus/goservice/business/domain/userbus"
)

// QueryFilter holds the available fields a query can be filtered on.
// We are using pointer semantics because the With API mutates the value.
type QueryFilter struct {
	ID       *uuid.UUID
	Name     *productbus.Name
	Cost     *float64
	Quantity *int
	UserName *userbus.Name
}
