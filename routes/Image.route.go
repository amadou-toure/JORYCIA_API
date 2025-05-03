package routes

import (
	"jorycia_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func ImageRoutes(app *fiber.App) {
    // Example of setting up routes for perfumes
    perfumeRoutes := app.Group("/image")
   // perfumeRoutes.Get("/", handlers.GetImage)
    perfumeRoutes.Get("/:fileName", handlers.GetImage)
    
}
