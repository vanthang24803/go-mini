// router/auth.go
package router

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/vanthang24803/mini/internal/controller"
)

type authHandler struct {
	validate   *validator.Validate
	controller *controller.AuthController
}

func newAuthHandler() *authHandler {
	return &authHandler{
		validate:   validator.New(),
		controller: controller.NewAuthController(),
	}
}
func SetupAuthRoutes(r fiber.Router) {
	authHandler := newAuthHandler()
	auth := r.Group("/auth")

	auth.Post("/login", authHandler.controller.Login)
	auth.Post("/register", authHandler.controller.Register)
}
