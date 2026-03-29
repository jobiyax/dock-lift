package project

import "github.com/gofiber/fiber/v3"

// Définition des routes
func RegisterRoutes(app fiber.Router) {
	app.Post("/projects", createProject)
	app.Get("/projects", getProjects)
}

func createProject(c fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}

	var body Request

	// Parsing du JSON
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"erreur": "Demande invalide",
		})
	}

	project := Create(body.Name)

	return c.JSON(project)
}

// Récupération de la liste
func getProjects(c fiber.Ctx) error {
	return c.JSON(GetAll())
}
