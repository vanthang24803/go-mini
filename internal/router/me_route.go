package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/internal/controller"
	m "github.com/vanthang24803/mini/internal/middleware"
)

func SetupMeRoutes(r fiber.Router) {
	ctrl := controller.NewMeController()
	route := r.Group("/me")

	route.Post("/", m.HandlerAuthentication(), ctrl.Profile)
	route.Post("/update", m.HandlerAuthentication(), ctrl.UpdateProfile)
	route.Post("/active", m.HandlerAuthentication(), m.HandlerAuthorization([]string{"root"}), ctrl.ActiveAccount)
}
