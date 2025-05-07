package routes

import (
	"jorycia_api/handlers"
	"jorycia_api/utils"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(app *fiber.App) {
    // Example of setting up routes for products
    productRoutes := app.Group("/product")

    productRoutes.Get("/", handlers.GetProducts)
    productRoutes.Get("/:id", handlers.GetOneProduct)
    productRoutes.Post("/", utils.Token.VerifyToken("your-secret-key"), handlers.AddProduct)
    //productRoutes.Put("/:id", handlers.UpdateProduct)
    productRoutes.Delete("/:id", handlers.DeleteProduct)
}
