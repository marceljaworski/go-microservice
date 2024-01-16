package model

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json:"order_id"`
	CustomerID  uuid.UUID  `json:"customer_id"`
	Products    Products   `json:"products"`
	CreatedAt   *time.Time `json:"created_at"`
	ShippedAt   *time.Time `json:"shipped_at"`
	CompletedAt *time.Time `json:"completed_at"`
}

type Products struct {
	ProductNo uuid.UUID `json:"product_no"`
	Quantity  uint      `json:"quantity"`
	Price     uint      `json:"price"`
}
