package models

import (
	"time"
)

type Order struct {
	ID              string     `bson:"_id,omitempty" json:"id"`
	UserID          string     `bson:"user_id" json:"userId"`
	Items           []OrderItem`bson:"items" json:"items"`
	CreatedAt       *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
	UpdatedAt       *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	StripeSessionID *string    `bson:"stripe_session_id,omitempty" json:"stripeSessionId,omitempty"`
	PaymentStatus   *string    `bson:"payment_status,omitempty" json:"paymentStatus,omitempty"`
}

type OrderItem struct {
	ProductID   string  `bson:"product_id" json:"product_id"`
	ProductName string  `bson:"product_name" json:"product_name"`
	Quantity    int64   `bson:"quantity" json:"quantity"`
	UnitPrice   float64 `bson:"unit_price" json:"unit_price"`
}
