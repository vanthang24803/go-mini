// router/setup.go
package router

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	// Root route
	app.Get("/", rootHandler)

	// API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	SetupAuthRoutes(v1)
	SetupHealthRoutes(v1)

}

func rootHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
