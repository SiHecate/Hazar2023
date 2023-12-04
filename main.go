package main

import (
	// "hazar/core/database"
	"hazar/router"

	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Zenithar say hello!")
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// database.Connect()
	app := fiber.New()
	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	router.Router(app)

	app.Listen(":8080")
}
