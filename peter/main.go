package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"peter/handler"
)

func routes(app *fiber.App) {
	handler.StaffRoutes(app)
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to staff management")
	})
	routes(app)
	err := app.Listen(":8080")
	if err != nil {
		log.Fatal(err)
		return
	}
}
