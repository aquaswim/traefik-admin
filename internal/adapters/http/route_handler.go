package http

import (
	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/application"
	"traefik-admin-go/internal/domain"
)

// RouteHandler represents the HTTP handler for routes
type RouteHandler struct {
	routeService *application.RouteService
}

// NewRouteHandler creates a new route handler
func NewRouteHandler(routeService *application.RouteService) *RouteHandler {
	return &RouteHandler{
		routeService: routeService,
	}
}

// RegisterRoutes registers all route HTTP routes
func (h *RouteHandler) RegisterRoutes(api fiber.Router) {
	routes := api.Group("/routes")
	routes.Get("/", h.GetAllRoutes)
	routes.Get("/:id", h.GetRouteByID)
	routes.Post("/", h.CreateRoute)
	routes.Put("/:id", h.UpdateRoute)
	routes.Delete("/:id", h.DeleteRoute)
}

// GetAllRoutes handles the get all routes request
func (h *RouteHandler) GetAllRoutes(c *fiber.Ctx) error {
	routes, err := h.routeService.GetAllRoutes()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(routes)
}

// GetRouteByID handles the get route by ID request
func (h *RouteHandler) GetRouteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	route, err := h.routeService.GetRouteByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.JSON(route)
}

// CreateRoute handles the create route request
func (h *RouteHandler) CreateRoute(c *fiber.Ctx) error {
	var route domain.RouteModel
	if err := c.BodyParser(&route); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := h.routeService.CreateRoute(&route); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(route)
}

// UpdateRoute handles the update route request
func (h *RouteHandler) UpdateRoute(c *fiber.Ctx) error {
	id := c.Params("id")
	var route domain.RouteModel
	if err := c.BodyParser(&route); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	// Ensure the ID in the URL matches the ID in the body
	route.ID = id

	if err := h.routeService.UpdateRoute(&route); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(route)
}

// DeleteRoute handles the delete route request
func (h *RouteHandler) DeleteRoute(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := h.routeService.DeleteRoute(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
