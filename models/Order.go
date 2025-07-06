package models

import (
	"time"
)

type Order struct {
	ID              string     `bson:"_id,omitempty" json:"id"`
	UserID          string    `bson:"user_id,omitempty" json:"userId,omitempty"`
	Items           []OrderItem`bson:"items" json:"items"`
	Total           float64   `bson:"total" json:"total"`
	CreatedAt       *time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	StripeSessionID *string    `bson:"stripe_session_id,omitempty" json:"stripeSessionId,omitempty"`
	ShippingAddress *string	 	`bson:"stripe_shipping_address,omitempty" json:"stripeShippingAddress"`
	PaymentStatus   *string    `bson:"payment_status,omitempty" json:"paymentStatus,omitempty"`
	Status          *string    `bson:"status" json:"status"`
}

type OrderItem struct {
	ProductID   string  `bson:"productId" json:"productId"`
	ProductName string  `bson:"productName" json:"productName"`
	Quantity    int64   `bson:"quantity" json:"quantity"`
	UnitPrice   float64 `bson:"unitPrice" json:"unitPrice"`
}
