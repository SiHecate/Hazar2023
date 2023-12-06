package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/hedhyw/Go-Serial-Detector/pkg/v1/serialdet"
	"github.com/tarm/serial"
)

var SerialPort *serial.Port

func ConnectionHandler(c *fiber.Ctx) error {
	var connectionRequest struct {
		PortName string `json:"port_name"`
		BaudRate int    `json:"baud_rate"`
	}

	if err := c.BodyParser(&connectionRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Invalid request data: " + err.Error(),
		})
	}

	serialConfig := &serial.Config{
		Name: connectionRequest.PortName,
		Baud: connectionRequest.BaudRate,
	}

	fmt.Println("Connection Request:", connectionRequest)

	s, err := serial.OpenPort(serialConfig)
	if err != nil {
		fmt.Println("Error opening serial port:", err)
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	SerialPort = s

	responseData := map[string]string{"status": "Connection success!"}
	return c.JSON(responseData)
}

func ConnectedPorts(c *fiber.Ctx) error {
	if list, err := serialdet.List(); err == nil {
		fmt.Println("Connected serial ports:")
		for _, p := range list {
			fmt.Println(p)
		}
		return c.SendString("List of connected serial ports printed to console.")
	} else {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
}

func SerialReadHandler(c *fiber.Ctx) error {
	if SerialPort == nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Serial port is not open")
	}

	buf := make([]byte, 128)
	n, err := SerialPort.Read(buf)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	responseData := map[string]string{"data": string(buf[:n])}
	return c.JSON(responseData)
}
