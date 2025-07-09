package http

import (
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
	"traefik-admin-go/internal/application"
)

// TraefikConfigHandler represents the HTTP handler for traefikConfig
type TraefikConfigHandler struct {
	traefikConfigService *application.TraefikConfigService
}

// NewTraefikConfigHandler creates a new service handler
func NewTraefikConfigHandler(
	traefikConfigService *application.TraefikConfigService,
) *TraefikConfigHandler {
	return &TraefikConfigHandler{
		traefikConfigService: traefikConfigService,
	}
}

// RegisterRoutes registers all service HTTP routes
func (h *TraefikConfigHandler) RegisterRoutes(api fiber.Router) {
	r := api.Group("/traefik-config")
	r.Get("/json", h.renderJson)
	r.Get("/yaml", h.renderYaml)
}

func (h *TraefikConfigHandler) renderJson(ctx *fiber.Ctx) error {
	cfg, err := h.traefikConfigService.GetConfig()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(cfg)
}

func (h *TraefikConfigHandler) renderYaml(ctx *fiber.Ctx) error {
	cfg, err := h.traefikConfigService.GetConfig()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	yamlCfg, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return ctx.Send(yamlCfg)
}
