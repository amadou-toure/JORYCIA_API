package routes

import (
	"jorycia_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func PerfumeRoutes(app *fiber.App) {
    // Example of setting up routes for perfumes
    perfumeRoutes := app.Group("/perfume")

    perfumeRoutes.Get("/", handlers.GetPerfumes)
    perfumeRoutes.Get("/:id", handlers.GetOnePerfume)
    perfumeRoutes.Post("/", handlers.AddPerfume)
    //perfumeRoutes.Put("/:id", handlers.UpdatePerfume)
    perfumeRoutes.Delete("/:id", handlers.DeletePerfume)
}
