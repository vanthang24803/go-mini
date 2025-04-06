package router

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/pkg/logger"
	"github.com/vanthang24803/mini/pkg/util"
)

func SetupRoutes(app *fiber.App) {
	log := logger.GetLogger()

	// Root route
	app.Get("/", func(c *fiber.Ctx) error {
		log.Info("Handling root request")
		return c.JSON(util.SuccessResponse("Welcome to Mini App!", fiber.Map{
			"message": "Hello, World!",
		}))
	})

	// API routes
	api := app.Group("/api")
	v1 := api.Group("/v1")

	v1.Get("/health", func(c *fiber.Ctx) error {
		log.Info("Health check requested")
		return c.JSON(util.SuccessResponse("Service is healthy", fiber.Map{
			"status": "ok",
			"uptime": time.Since(time.Now()).String(),
		}))
	})

	v1.Get("/jwt", func(c *fiber.Ctx) error {
		accessToken, refreshToken, err := util.GenerateJWT(123, "maynguyen")

		if err != nil {
			log.Error("error generate jwt")
		}

		log.Info("JWT requested")
		return c.JSON(util.SuccessResponse("JWT generated", fiber.Map{
			util.ACCESS_TOKEN:  accessToken,
			util.REFRESH_TOKEN: refreshToken,
		}))
	})
}
