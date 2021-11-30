package main

import (
	"log"

	"github.com/abiiranathan/goclinic/database"
	"github.com/abiiranathan/goclinic/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load(".env")
	database.ConnectToDatabase()
}

func main() {
	app := fiber.New(fiber.Config{
		ServerHeader:  "Eclinic HMS",
		StrictRouting: true,
	})

	routes.SetupRoutes(app)
	log.Fatal(app.Listen(":8080"))
}
