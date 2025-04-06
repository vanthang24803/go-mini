package main

import (
	"github.com/gofiber/fiber/v2"
	fiberLogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/vanthang24803/mini/internal/config"
	"github.com/vanthang24803/mini/internal/middleware"
	"github.com/vanthang24803/mini/internal/router"
	"github.com/vanthang24803/mini/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	cfg := config.New()

	logger.Init()
	log := logger.GetLogger()
	defer log.Sync()

	app := fiber.New(fiber.Config{
		AppName:               cfg.AppName,
		DisableStartupMessage: true,
		ErrorHandler:          middleware.ErrorHandler,
	})

	app.Use(fiberLogger.New())
	app.Use(recover.New())

	// Setup routes
	router.SetupRoutes(app)

	// Add NotFound handler
	app.Use(middleware.NotFoundHandler())

	log.Info("Server is starting",
		zap.String("address", "localhost"+cfg.Server.Port),
		zap.String("app_name", cfg.AppName))

	if err := app.Listen(cfg.Server.Port); err != nil {
		log.Fatal("Server failed to start", zap.Error(err))
	}
}
