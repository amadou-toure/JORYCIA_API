package models

import (
	"time"
)

type Order struct {
	ID              string        `bson:"_id,omitempty" json:"id"`
	UserID          string        `bson:"user_id,omitempty" json:"userId,omitempty"`
	Items           []OrderItem   `bson:"items" json:"items"`
	Total           float64       `bson:"total" json:"total"`
	CreatedAt       *time.Time    `bson:"created_at" json:"createdAt"`
	UpdatedAt       *time.Time    `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	StripeSessionID *string       `bson:"stripe_session_id,omitempty" json:"stripeSessionId,omitempty"`
	ShippingAddress StripeAddress `bson:"stripe_shipping_address" json:"shippingAddress"`
	PaymentStatus   *string       `bson:"payment_status" json:"paymentStatus"`
	Status          *string       `bson:"status" json:"status"`
}

type OrderItem struct {
	ProductID   string  `bson:"productId" json:"productId"`
	ProductName string  `bson:"productName" json:"productName"`
	Quantity    int64   `bson:"quantity" json:"quantity"`
	UnitPrice   float64 `bson:"unitPrice" json:"unitPrice"`
}

type StripeAddress struct {
	City       *string `bson:"city" json:"city"`
	Country    *string `bson:"country" json:"country"`
	Line1      *string `bson:"line1" json:"line1"`
	Line2      string `bson:"line2,omitempty" json:"line2,omitempty"`
	PostalCode string `bson:"postal_code" json:"postal_code"`
	State      string `bson:"state" json:"state"`
}
