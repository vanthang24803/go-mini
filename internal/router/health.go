package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

func SetupHealthRoutes(r fiber.Router) {
	r.Get("/health", healthHandler)
}

func healthHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "ok",
		"uptime": time.Since(time.Now()).String(),
	})
}
