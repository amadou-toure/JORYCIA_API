package handlers

import (
	"time"

	"jorycia_api/Database"
	"jorycia_api/HTTP_CODE"
	"jorycia_api/models"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// CartItem structure pour la réception du panier côté client
type CartItem struct {
	ProductID string `json:"product_id"`
	Quantity  int64  `json:"quantity"`
}

// CreateOrder crée une nouvelle commande à partir d'un panier
func CreateOrder(c *fiber.Ctx) error {
	userID := c.Query("user_id") // ou récupéré via middleware auth
	if userID == "" {
		return c.Status(HTTP_CODE.Bad_request).SendString("User ID manquant")
	}

	var cart []CartItem
	if err := c.BodyParser(&cart); err != nil {
		return c.Status(HTTP_CODE.Bad_request).SendString("Données de panier invalides")
	}

	var orderItems []models.OrderItem
	var total float64

	for _, item := range cart {
		var product models.Product
		err := Database.Mg.Db.Collection("Product").FindOne(c.Context(), bson.M{"_id": item.ProductID}).Decode(&product)
		if err != nil {
			return c.Status(HTTP_CODE.Not_found).SendString("Produit non trouvé: " + item.ProductID)
		}

		orderItems = append(orderItems, models.OrderItem{
			ProductID:   product.ID,
			ProductName: product.Name,
			Quantity:    item.Quantity,
			UnitPrice:   product.Price,
		})
		total += product.Price * float64(item.Quantity)
	}

	now := time.Now()
	order := models.Order{
		UserID:        userID,
		Items:         orderItems,
		CreatedAt:     &now,
		UpdatedAt:     &now,
		PaymentStatus: stringPtr("pending"),
	}

	_, err := Database.Mg.Db.Collection("Order").InsertOne(c.Context(), order)
	if err != nil {
		return c.Status(HTTP_CODE.Server_error).SendString("Erreur enregistrement commande")
	}

	return c.Status(HTTP_CODE.Created).JSON(order)
}

func stringPtr(s string) *string {
	return &s
}
