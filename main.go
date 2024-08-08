package main

import (
	"log"

	"github.com/Ameya117/fiber-api/database"
	"github.com/Ameya117/fiber-api/routes"
	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome to the API!")
}

func setupRoutes(app *fiber.App) {
	// welcome route
	app.Get("/api", welcome)

	//user endpoints
	app.Post("/api/users", routes.CreateUser)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello!!")
	})

	app.Get("/api", welcome)

	// Start the server on port 3000
	log.Fatal(app.Listen(":3000"))
}
