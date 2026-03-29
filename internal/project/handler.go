package project

import "github.com/gofiber/fiber/v3"

// Injection de dépendance du service
type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(app fiber.Router) {
	app.Post("/projects", h.createProject)
	app.Get("/projects", h.getProjects)
}

func (h *Handler) createProject(c fiber.Ctx) error {
	type Request struct {
		Name string `json:"name"`
	}

	var body Request

	// Vérification du corps de la requête
	if err := c.Bind().Body(&body); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"erreur": "Requête invalide",
		})
	}

	// Appel à la logique métier
	project, err := h.service.Create(body.Name)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"erreur": "Impossible de créer le projet",
		})
	}

	return c.JSON(project)
}

func (h *Handler) getProjects(c fiber.Ctx) error {
	return c.JSON(h.service.GetAll())
}
