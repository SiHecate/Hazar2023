package main

import (
	// "hazar/core/database"
	"hazar/core/database"
	"hazar/router"

	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Hazar say hello!")
	database.Connect()
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	router.Router(app)
	router.StartWebSocket(app)
	app.Listen(":8080")
}
