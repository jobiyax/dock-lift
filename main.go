package main

import "github.com/gofiber/fiber/v3"

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(map[string]string{
			"message": "Dock Lift",
		})
	})

	app.Listen(":9000")
}
