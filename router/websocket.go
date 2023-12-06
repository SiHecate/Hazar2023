package router

import (
	"log"
	"time"

	controllers "hazar/controller"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func StartWebSocket(app *fiber.App) {
	app.Use("/websocket", websocket.New(func(c *websocket.Conn) {
		log.Println("WebSocket port open")
		defer log.Println("WebSocket port off")

		for {
			err := controllers.SerialReadHandler(c)
			if err != nil {
				log.Println("SerialReadHandler error:", err)
				break
			}

			// Belirli bir süre bekletme ekleyerek yavaş alım sağla
			time.Sleep(1000 * time.Millisecond)
		}
	}))

	app.Listen(":8080")
}
