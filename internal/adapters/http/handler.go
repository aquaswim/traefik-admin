package http

import (
	"github.com/gofiber/fiber/v2"
	"traefik-admin-go/internal/application"
)

// Handler represents the HTTP handler
type Handler struct {
	appService *application.Service
}

// NewHandler creates a new HTTP handler
func NewHandler(appService *application.Service) *Handler {
	return &Handler{
		appService: appService,
	}
}

// RegisterRoutes registers all HTTP routes
func (h *Handler) RegisterRoutes(app *fiber.App) {
	// Register the hello world route
	app.Get("/", h.HelloWorld)
}

// HelloWorld handles the hello world request
func (h *Handler) HelloWorld(c *fiber.Ctx) error {
	message := h.appService.GetHelloMessage()
	return c.SendString(message)
}
