package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
)

func welcome(c fiber.Ctx) error {
	return c.SendString("Welcome to the API!")
}

func main() {
    app := fiber.New()

    app.Get("/", func(c fiber.Ctx) error {
        // Send a string response to the client
        return c.SendString("Hello, World 👋!")
    })

	app.Get("/api", welcome)




    // Start the server on port 3000
    log.Fatal(app.Listen(":3000"))
}