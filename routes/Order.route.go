package routes

import (
	"jorycia_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func OrderRoutes(router *fiber.App) {
	orderRoutes := router.Group("/order")
	orderRoutes.Post("/", handlers.CreateOrder)
	orderRoutes.Get("/", handlers.GetOrders)
	orderRoutes.Get("/:id", handlers.GetOneOrder)
	orderRoutes.Put("/:id", handlers.UpdateOrder)
	orderRoutes.Get("/user/:user_id", handlers.GetUsersOrders)
}