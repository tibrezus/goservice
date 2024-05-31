package vproductbus

import (
	"time"

	"github.com/google/uuid"
	"github.com/tiberzus/goservice/business/domain/productbus"
	"github.com/tiberzus/goservice/business/domain/userbus"
)

// Product represents an individual product with extended information.
type Product struct {
	ID          uuid.UUID
	UserID      uuid.UUID
	Name        productbus.Name
	Cost        float64
	Quantity    int
	DateCreated time.Time
	DateUpdated time.Time
	UserName    userbus.Name
}
