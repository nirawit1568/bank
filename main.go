package main

import (
	"bank/route"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// Register user routes
	route.Routes(app)

	// Start the server on port 3000
	app.Listen(":3000")
}
