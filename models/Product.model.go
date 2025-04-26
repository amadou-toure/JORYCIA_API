package models

type Product struct {
	ID          string   `bson:"_id,omitempty" json:"id"`
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Notes       []string `bson:"notes" json:"notes"`
	Rating      int      `bson:"rating" json:"rating"`
	Quantity    int      `bson:"quantity" json:"quantity"`
	Price       int      `bson:"price" json:"price"` // Price in cents
	Image       []string `bson:"image" json:"image"`
	Metadata    map[string]string `bson:"metadata,omitempty" json:"metadata,omitempty"` // Optional metadata
	StripeProductID string            `bson:"stripe_product_id" json:"stripe_product_id"`
	StripePriceID   string           `bson:"stripe_price_id" json:"stripe_price_id"`
}

// The new version adds two new fields to the Product struct: StripeProductID and StripePriceID. These fields are used to store the IDs of the Stripe product and price associated with the product.