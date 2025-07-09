package http

import (
	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/application"
)

// Handler represents the HTTP handler
type Handler struct {
	serviceHandler *ServiceHandler
	routeHandler   *RouteHandler
}

// NewHandler creates a new HTTP handler
func NewHandler(serviceService *application.ServiceService, routeService *application.RouteService) *Handler {
	return &Handler{
		serviceHandler: NewServiceHandler(serviceService),
		routeHandler:   NewRouteHandler(routeService),
	}
}

// RegisterRoutes registers all HTTP routes
func (h *Handler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Register service routes
	h.serviceHandler.RegisterRoutes(api)

	// Register route routes
	h.routeHandler.RegisterRoutes(api)
}
