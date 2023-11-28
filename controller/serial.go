package controller

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.bug.st/serial"
)

var SerialPort *serial.Port

// func ConnectionHandler(c *fiber.Ctx) error {
// 	var connectionRequest struct {
// 		PortName string `json:"port_name"`
// 		BaudRate int    `json:"baud_rate"`
// 	}

// 	if err := c.BodyParser(&connectionRequest); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error":   true,
// 			"message": "Invalid request data: " + err.Error(),
// 		})
// 	}

// 	serialConfig := &serial.Config{
// 		Name: connectionRequest.PortName,
// 		Baud: connectionRequest.BaudRate,
// 	}

// 	fmt.Println("Connection Request:", connectionRequest)

// 	s, err := serial.OpenPort(serialConfig)
// 	if err != nil {
// 		fmt.Println("Error opening serial port:", err)
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	SerialPort = s

// 	responseData := map[string]string{"status": "Connection success!"}
// 	return c.JSON(responseData)
// }

func SerialFindHandler(c *fiber.Ctx) error {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if len(ports) == 0 {
		return c.Status(fiber.StatusOK).SendString("No serial ports found!")
	}

	portsJSON, err := json.Marshal(ports)
	if err != nil {
		log.Fatal(err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.Status(fiber.StatusOK).Send(portsJSON)
}

// func SerialReadHandler(c *fiber.Ctx) error {
// 	buf := make([]byte, 128)
// 	n, err := SerialPort.Read(buf)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
// 	}

// 	responseData := map[string]string{"data": string(buf[:n])}
// 	return c.JSON(responseData)
// }
