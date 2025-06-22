// main.go
package main

import (
	"fmt"
	"log"
	"os"

	"jorycia_api/Database"
	"jorycia_api/HTTP_CODE"
	"jorycia_api/routes"

	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	dburi:=os.Getenv("MONGO_URI")
	if dburi ==""{
		fmt.Println("vide!!!")
		log.Println("Tentative de connexion à MongoDB avec l'URI :")
	}
	err := Database.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New(fiber.Config{
		Prefork:       false,
		CaseSensitive: true,
		StrictRouting: true,
		BodyLimit:     10 * 1024 * 1024,
		ServerHeader:  "Fiber",
		AppName:       "jorycia_api v1.0.0",
	})

	// Use CORS middleware configured for frontend on localhost:4173, allow all headers (including Authorization), and credentials support
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowHeaders:     "*",
		AllowMethods:     "GET, POST, PUT, DELETE, OPTIONS",
		AllowCredentials: false,
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
	port := os.Getenv("PORT")
	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}
	
	log.Println("Listening on port", port)
	log.Fatal(app.Listen(port))
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Recovered from panic: %v\n", r)
		}
	}()
}
