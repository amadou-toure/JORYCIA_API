package models

type Address struct {
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	ZipCode string `bson:"zip_code" json:"zipCode"`
	Country string `bson:"country" json:"country"`
}