package run

import (
	"github.com/gofiber/fiber/v3"
	"github.com/ryszhio/goauth/router"
)

func InitializeApp() error {
	app := fiber.New()

	router.SetupRoutes(app)

	app.Get("/", greetingResponse)
	app.Post("/", greetingResponse)

	app.Listen(":3000")
	return nil
}

func greetingResponse(c fiber.Ctx) error {
	return c.SendString("goauth is running")
}
