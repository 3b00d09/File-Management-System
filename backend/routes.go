package main

import "github.com/gofiber/fiber/v3"

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello world!")
	})
}
