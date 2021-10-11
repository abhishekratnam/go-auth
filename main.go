package main

import (
	"github.com/abhishekratnam/go-auth/database"
	"github.com/abhishekratnam/go-auth/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	app := fiber.New()
	routes.Setup(app)

	app.Listen(":8000")
}
