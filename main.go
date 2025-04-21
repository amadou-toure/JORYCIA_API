// main.go
package main

import (
	"log"
	"os"

	"jorycia_api/Database"
	"jorycia_api/HTTP_CODE"
	"jorycia_api/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	err := Database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     10 * 1024 * 1024,
		ServerHeader:  "Fiber",
		AppName:       "jorycia_api v1.0.0",
	})

	// Use CORS middleware with permissive settings
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	routes.ProductRoutes(app)
	routes.UserRoutes(app)
	routes.ImageRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString(os.Getenv("PORT"))
		if err != nil {
			return c.Status(HTTP_CODE.Server_error).SendString("Server failed to respond")
		}
		return nil
	})

	app.Listen(":8080")
}
