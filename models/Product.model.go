package models

type Product struct {
	ID          string   `bson:"_id,omitempty"`
	Name        string   `bson:"name" json:"name"`
	Description string   `bson:"description" json:"description"`
	Notes       []string `bson:"notes"`
	Rating      int      `bson:"rating"`
	Quantity    int      `bson:"quantity"`
	Price       int      `bson:"price"` // Price in cents
	Images      []string `bson:"image" json:"images"`
	Metadata    map[string]string `bson:"metadata,omitempty" json:"metadata,omitempty"` // Optional metadata
	StripeProductID string            `bson:"stripe_product_id" json:"stripe_product_id"`
	StripePriceID   string            `bson:"stripe_price_id" json:"stripe_price_id"`
}

// The new version adds two new fields to the Product struct: StripeProductID and StripePriceID. These fields are used to store the IDs of the Stripe product and price associated with the product.