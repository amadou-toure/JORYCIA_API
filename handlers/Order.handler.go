package handlers

import (
	"fmt"
	"jorycia_api/Database"
	"jorycia_api/HTTP_CODE"
	"jorycia_api/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// CartItem structure pour la réception du panier côté client

// CreateOrder crée une nouvelle commande à partir d'un panier
func CreateOrder(c *fiber.Ctx) error {
	var Order models.Order
	fmt.Println(Order.ShippingAddress)
	fmt.Println( Order.ShippingAddress)
	err := c.BodyParser(&Order)
	if err != nil {
		fmt.Println(Order)
		return c.Status(HTTP_CODE.Bad_request).SendString("Erreur de parsing de la requête")
	}

	now := time.Now().UTC()
	Order.CreatedAt = &now
	fmt.Println(Order)
	_, err = Database.Mg.Db.Collection("Order").InsertOne(c.Context(), Order)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString("Erreur insertion de la commande: "+err.Error())
	}

	err= SendMail("support@jorycia.ca","amadoumojatoure@outlook.fr","new order","Hello, une nouvelle commande a ete passee!")
	if err !=nil{
		return c.Status(HTTP_CODE.Server_error).SendString("Impossible d'envoyer l'email: "+err.Error())
	}
	return c.Status(HTTP_CODE.Created).JSON(Order)
}



func GetOneOrder(c *fiber.Ctx) error {
	

	orderID := c.Params("id")
	
	if orderID == "" {
		return c.Status(HTTP_CODE.Bad_request).SendString("Order ID manquant")
	}
	id:= c.Params("id")
	ID,err:= primitive.ObjectIDFromHex(id)
	if err != nil{
		return c.Status(HTTP_CODE.Bad_request).SendString("unvalid id")
	}
	filter:= bson.D{{Key: "_id",Value: ID}}
	product:=new(models.Order)
	query:= Database.Mg.Db.Collection("Order").FindOne(c.Context(),filter)
	if query.Err() != nil{
		if query.Err() == mongo.ErrNoDocuments{
		return c.Status(HTTP_CODE.Not_found).SendString("Orders not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(query.Err().Error())
	}
	err=query.Decode(product)
	if err != nil{
		if err == mongo.ErrNoDocuments{
			return c.Status(HTTP_CODE.Not_found).SendString("Orders not found 2")
		}
	}
	return c.Status(HTTP_CODE.Ok).JSON(product)
}
func GetOrders(c* fiber.Ctx) error {
	filter:=bson.D{{}}
	var orders []models.Order
	 result,err := Database.Mg.Db.Collection("Order").Find(c.Context(),filter)
	 if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(HTTP_CODE.Not_found).SendString("items not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 err = result.All(c.Context(),&orders)
	 if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	 }
	 return c.Status(HTTP_CODE.Ok).JSON(orders)
}
func GetUsersOrders(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	if userID == "" {
		return c.Status(HTTP_CODE.Bad_request).SendString("User ID manquant")
	}

	filter := bson.D{{Key:"user_id", Value: userID}}
	var orders []models.Order
	result, err := Database.Mg.Db.Collection("Order").Find(c.Context(), filter)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(HTTP_CODE.Not_found).SendString("items not found")
		}
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	err = result.All(c.Context(), &orders)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	return c.Status(HTTP_CODE.Ok).JSON(orders)
}
func UpdateOrder(c *fiber.Ctx) error {
	ID := c.Params("id")
	orderID,err:= primitive.ObjectIDFromHex(ID)
	if err != nil{
		return c.Status(HTTP_CODE.Bad_request).SendString("invalid id")
	}

	var order models.Order
	err = c.BodyParser(&order)
	if err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString("Erreur de parsing de la requête")
	}
	filter := bson.D{{Key:"_id", Value:orderID}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "status", Value: order.Status},
		}},
	}
	result, err := Database.Mg.Db.Collection("Order").UpdateOne(c.Context(), filter, update)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString(err.Error())
	}
	if result.ModifiedCount == 0 {
		return c.Status(HTTP_CODE.Not_found).SendString("Order not found")
	}
	return c.Status(HTTP_CODE.Ok).SendString("Order updated")
}

