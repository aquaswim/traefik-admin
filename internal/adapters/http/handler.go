package http

import (
	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/application"
)

// Handler represents the HTTP handler
type Handler struct {
	serviceHandler       *ServiceHandler
	routeHandler         *RouteHandler
	traefikConfigHandler *TraefikConfigHandler
}

// NewHandler creates a new HTTP handler
func NewHandler(
	serviceService *application.ServiceService,
	routeService *application.RouteService,
	traefikConfigService *application.TraefikConfigService,
) *Handler {
	return &Handler{
		serviceHandler:       NewServiceHandler(serviceService),
		routeHandler:         NewRouteHandler(routeService),
		traefikConfigHandler: NewTraefikConfigHandler(traefikConfigService),
	}
}

// RegisterRoutes registers all HTTP routes
func (h *Handler) RegisterRoutes(app *fiber.App) {
	api := app.Group("/api")

	// Register service routes
	h.serviceHandler.RegisterRoutes(api)

	// Register route routes
	h.routeHandler.RegisterRoutes(api)

	// Register traefik config routes
	h.traefikConfigHandler.RegisterRoutes(api)
}
