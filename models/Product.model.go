package models

import (
	"time"
)

type Product struct {
	ID              string            `bson:"_id,omitempty" json:"id"`
	Name            string            `bson:"name" json:"name"`
	Description     string            `bson:"description" json:"description"`
	Notes           []string          `bson:"notes" json:"notes"`
	Rating          int               `bson:"rating" json:"rating"`
	InStock       int               `bson:"inStock" json:"inStock"`
	Price           float64           `bson:"price" json:"price"` // Price in cents
	Image           []string          `bson:"image" json:"image"`
	Metadata        map[string]string `bson:"metadata,omitempty" json:"metadata,omitempty"`
	StripeProductID string            `bson:"stripe_product_id" json:"stripeProductId"`
	StripePriceID   string            `bson:"stripe_price_id" json:"stripePriceId"`
}



// type Order struct {
// 	ID string `bson:"_id,omitempty" json:"id"`
// 	stripeCustomerId string `bson:"stripe_customer_id" json:"stripeCustomerId"`
// 	Items []OrderItem `bson:"items" json:"products"`
// 	CreatedAt *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
// 	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
// 	Status string `bson:"status" json:"status"`
// }

// type OrderItem struct {
// 	stripeProductId string `bson:"stripe_product_id" json:"stripeProductId"`
// 	stripePriceId string `bson:"stripe_price_id" json:"stripePriceId"`
// 	Quantity int `bson:"quantity" json:"quantity"`
// }

type SaveForLater struct {
	ID string `bson:"_id,omitempty" json:"id"`
	UserID    string `bson:"user_id" json:"userId"`
	Items     []Product `bson:"items" json:"products"`
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}
