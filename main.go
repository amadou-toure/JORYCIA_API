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
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erreur lors du chargement du fichier .env")
	}

	err = Database.Connect()
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
	routes.PaymentRoutes(app)
	routes.OrderRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString(os.Getenv("API_URL"))
		if err != nil {
			return c.Status(HTTP_CODE.Server_error).SendString("Server failed to respond")
		}
		return nil
	})

	app.Listen(os.Getenv("PORT"))
}
