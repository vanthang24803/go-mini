package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/internal/controller"
	"github.com/vanthang24803/mini/internal/middleware"
)

func SetupMeRoutes(r fiber.Router) {
	ctrl := controller.NewMeController()
	route := r.Group("/me")

	route.Post("/", middleware.HanlderAuthentication(), ctrl.Profile)
}
