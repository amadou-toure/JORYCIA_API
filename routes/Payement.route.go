package routes

import (
	"jorycia_api/handlers"
	"jorycia_api/utils"

	"github.com/gofiber/fiber/v2"
)

func PaymentRoutes(router *fiber.App) {
	paymentRoutes := router.Group("/payment")
	paymentRoutes.Post("/ProceedToPayment/", utils.Token.VerifyToken("your-secret-key"), handlers.CreateCheckoutSession)
	//paymentRoutes.Post("/webhook", handlers.HandleWebhook)
}
