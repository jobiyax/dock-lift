package main

import (
	"dock-lift/internal/project"

	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Dock Lift API 🚀")
	})

	api := app.Group("/api")

	// Initialisation service et handler
	service := project.NewService()
	handler := project.NewHandler(service)

	handler.RegisterRoutes(api)

	// Lancement du serveur
	app.Listen(":9000")
}
