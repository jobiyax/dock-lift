package main

import (
	"dock-lift/internal/project"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	// Route racine
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Dock Lift API 🚀")
	})

	// Groupement et routes API
	api := app.Group("/api")

	project.RegisterRoutes(api)

	// Démarrage port 9000
	app.Listen(":9000")
}
