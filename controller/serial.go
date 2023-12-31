package controller

import (
	"encoding/json"
	"fmt"
	model "hazar/Model"
	"hazar/core/database"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
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

	responseData := map[string]string{"status": "Connection success!", "Port": connectionRequest.PortName}
	return c.JSON(responseData)
}

func SerialReadHandler(ws *websocket.Conn) error {
	if SerialPort == nil {
		return ws.WriteJSON(map[string]string{"error": "Serial port is not open"})
	}

	serialBuf := make([]byte, 128)
	n, err := SerialPort.Read(serialBuf)
	if err != nil {
		return ws.WriteJSON(map[string]string{"error": err.Error()})
	}

	var data model.SensorData

	fields := strings.Fields(string(serialBuf[:n]))
	for i := 0; i < len(fields); i += 2 {
		if i+1 >= len(fields) {
			break
		}
		switch fields[i] {
		case "encoder1:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Encoder1 = value

		case "encoder2:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Encoder2 = value

		case "encoder3:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Encoder3 = value

		case "encoder4:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Encoder4 = value

		case "Ax:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Ax = value

		case "Ay:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Ay = value

		case "Az:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Az = value

		case "Rx:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Rx = value

		case "Ry:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Ry = value

		case "Rz:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Rz = value

		case "Altitude:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Altitude = value

		case "Temp:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Temp = value

		case "Ayaj:":
			value, err := strconv.Atoi(fields[i+1])
			if err != nil {
				return ws.WriteJSON(map[string]string{"error": err.Error()})
			}
			data.Ayak = value
		}
	}

	data.Time = time.Now()

	// Convert the struct to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ws.WriteJSON(map[string]string{"error": err.Error()})
	}

	// Send JSON data through WebSocket
	responseData := map[string]string{"data": string(jsonData)}

	if err := database.Conn.Create(&data).Error; err != nil {
		return err
	}

	return ws.WriteJSON(responseData)
}

func GetSerialData(c *fiber.Ctx) error {
	// todo: özel databaseler açıp onları tek tek göstermek için fonksiyon eklenecek.
	var existingData []model.SensorData
	if err := database.Conn.Find(&existingData).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Database error"})
	}

	if len(existingData) == 0 {
		return c.Status(404).JSON(fiber.Map{"error": "No products found"})
	}

	return c.JSON(fiber.Map{"message": existingData})
}
