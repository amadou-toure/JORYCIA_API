package models

import "time"
	
type User struct {
	ID        string    `bson:"_id,omitempty" json:"id"`
	FirstName string    `bson:"first_name" json:"firstName"`
	LastName  string    `bson:"last_name" json:"lastName"`
	Email     string    `bson:"email" json:"email"`
	Password  string    `bson:"password,omitempty" json:"-"` // hashé, jamais renvoyé côté client
	Phone     string    `bson:"phone,omitempty" json:"phone,omitempty"`
	Address   Address   `bson:"address" json:"address"`       // struct séparée ci-dessous
	Orders    []string  `bson:"orders,omitempty" json:"orders,omitempty"` // IDs des commandes
	CreatedAt time.Time `bson:"created_at" json:"createdAt"`
	UpdatedAt time.Time `bson:"updated_at" json:"updatedAt"`
}	