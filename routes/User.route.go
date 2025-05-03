package routes

import (
	"jorycia_api/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(app *fiber.App) {
    // Example of setting up routes for perfumes
    UserRoutes := app.Group("/user")

    UserRoutes.Get("/", handlers.GetUsers)
    UserRoutes.Get("/:id", handlers.GetOneUser)
	UserRoutes.Post("/", handlers.CreateUser)
    //UserRoutes.Put("/:id", handlers.UpdatePerfume)
    UserRoutes.Delete("/:id", handlers.DeleteUser)
}
