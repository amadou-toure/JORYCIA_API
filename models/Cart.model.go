package models

type CartItem struct {
	Product Product `bson:"product" json:"product"`
	Quantity int `bson:"quantity" json:"quantity"`
}