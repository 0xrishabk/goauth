package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ryszhio/goauth/handler"
)

func SetupRoutes(app *fiber.App) {
	// Authentication Routes
	auth := app.Group("/api/auth")
	auth.Post("/login", handler.Login)
	auth.Post("/register", handler.Register)
}
