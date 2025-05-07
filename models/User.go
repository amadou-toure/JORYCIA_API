package models

import "time"
	
type User struct {
	ID               string     `bson:"_id,omitempty" json:"id"`
	FirstName        string     `bson:"first_name" json:"firstName"`
	LastName         string     `bson:"last_name" json:"lastName"`
	Email            string     `bson:"email" json:"email"`
	Password         *string    `bson:"password,omitempty" json:"password,omitempty"` // nil si OAuth
	Phone            string     `bson:"phone" json:"phone"`
	Address          string     `bson:"address" json:"address"`
	//Provider         string     `bson:"provider" json:"provider"` // "google" ou "credentials"
	Avatar           *string    `bson:"avatar,omitempty" json:"avatar,omitempty"`
	Role             string     `bson:"role" json:"role"` // "user" ou "admin"
	CreatedAt        *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
	UpdatedAt        *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
	//Verified         *bool      `bson:"verified,omitempty" json:"verified,omitempty"`
	//OrderIDs         []string   `bson:"order_ids,omitempty" json:"orderIDs,omitempty"`
	StripeCustomerID *string    `bson:"stripe_customer_id,omitempty" json:"stripeCustomerId,omitempty"`
}

type OAuthToken struct {
	UserID       string     `bson:"user_id" json:"userId"`
	Provider     string     `bson:"provider" json:"provider"` // ex: "google"
	AccessToken  string     `bson:"access_token" json:"accessToken"`
	RefreshToken *string    `bson:"refresh_token,omitempty" json:"refreshToken,omitempty"`
	ExpiryDate   *time.Time `bson:"expiry_date,omitempty" json:"expiryDate,omitempty"`
	CreatedAt    *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
	UpdatedAt    *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
}
//ajouter les models pour les customers et les editors
// package models

// import "time"
	
// // Customer represents an end customer in the system.
// type Customer struct {
// 	ID               string     `bson:"_id,omitempty" json:"id"`
// 	FirstName        string     `bson:"first_name" json:"firstName"`
// 	LastName         string     `bson:"last_name" json:"lastName"`
// 	Email            string     `bson:"email" json:"email"`
// 	Phone            string     `bson:"phone" json:"phone"`
// 	Address          string     `bson:"address" json:"address"`
// 	Avatar           *string    `bson:"avatar,omitempty" json:"avatar,omitempty"`
// 	CreatedAt        *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
// 	UpdatedAt        *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
// 	StripeCustomerID *string    `bson:"stripe_customer_id,omitempty" json:"stripeCustomerId,omitempty"`
// }

// // Editor represents a content or system editor with elevated permissions.
// type Editor struct {
// 	ID        string     `bson:"_id,omitempty" json:"id"`
// 	FirstName string     `bson:"first_name" json:"firstName"`
// 	LastName  string     `bson:"last_name" json:"lastName"`
// 	Email     string     `bson:"email" json:"email"`
// 	Password  *string    `bson:"password,omitempty" json:"password,omitempty"`
// 	Avatar    *string    `bson:"avatar,omitempty" json:"avatar,omitempty"`
// 	Role      string     `bson:"role" json:"role"`
// 	CreatedAt *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
// 	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
// }

// type OAuthToken struct {
// 	UserID       string     `bson:"user_id" json:"userId"`
// 	Provider     string     `bson:"provider" json:"provider"` // ex: "google"
// 	AccessToken  string     `bson:"access_token" json:"accessToken"`
// 	RefreshToken *string    `bson:"refresh_token,omitempty" json:"refreshToken,omitempty"`
// 	ExpiryDate   *time.Time `bson:"expiry_date,omitempty" json:"expiryDate,omitempty"`
// 	CreatedAt    *time.Time `bson:"created_at,omitempty" json:"createdAt,omitempty"`
// 	UpdatedAt    *time.Time `bson:"updated_at,omitempty" json:"updatedAt,omitempty"`
// }
