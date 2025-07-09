package web

import (
	"embed"
	"github.com/aquaswim/govite"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"html/template"
)

//go:embed dist/*
var dist embed.FS

//go:embed index.gohtml
var index string

func RegisterRoutes(app *fiber.App) {
	// vite glue
	vite := govite.NewWithFS(&govite.Config{
		ViteOutputPath: "./dist",
		AssetBaseUrl:   "/assets",
		IsReact:        true,
	}, dist)

	// index.html template
	templ, err := template.New("index").Parse(index)
	if err != nil {
		panic(err)
	}

	app.Get("/", func(c *fiber.Ctx) error {
		viteAsset := vite.MustGetBuilder().AddAsset("src/main.jsx")

		c.Set("Content-Type", "text/html; charset=utf-8")
		return templ.ExecuteTemplate(c.Context(), "index", fiber.Map{
			"title":   "Traefik Admin",
			"style":   template.HTML(viteAsset.CreateStyleTags()),
			"script":  template.HTML(viteAsset.CreateScriptTags()),
			"preload": template.HTML(viteAsset.CreatePreloadTags()),
		})
	})

	app.Use(adaptor.HTTPHandler(vite.FileServer()))
}
