package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/internal/controller"
	"github.com/vanthang24803/mini/internal/middleware"
)

func SetupAuthRoutes(r fiber.Router) {
	ctrl := controller.NewAuthController()
	route := r.Group("/auth")

	route.Post("/login", ctrl.Login)
	route.Post("/register", ctrl.Register)
	route.Post("/logout", middleware.HanlderAuthentication(), ctrl.Logout)
}
