package router

import (
	"hazar/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Router(app *fiber.App) {
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hazar rocket team says Hello world")
	})

	serial := app.Group("/serial")
	serial.Get("ab123", func(c *fiber.Ctx) error {
		return c.SendString("Serial Communication denemesi")
	})
	// serial.Post("connection", controller.ConnectionHandler)
	// serial.Get("read", controller.SerialReadHandler)
	serial.Get("ports", controller.SerialFindHandler)

}
