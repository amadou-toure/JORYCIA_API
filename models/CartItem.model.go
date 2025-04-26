package models

type CartItem struct {
		ProductID string `json:"product_id"`
		PriceID   string  `bson:"stripe_price_id" json:"stripe_price_id"`
		Quantity  int64  `json:"quantity"`
	}