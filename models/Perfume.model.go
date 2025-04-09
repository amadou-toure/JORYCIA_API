package models

type Perfume struct {
	ID       string `bson:"_id,omitempty"`
	Name     string             `bson:"name"`
	Description   string             `bson:"description"`
	Rating int             `bson:"rating"`
	Quantity int          `bson:"quantity"`
	Price int            `bson:"price"`
	Image []string         `bson:"image"`
}

 
   