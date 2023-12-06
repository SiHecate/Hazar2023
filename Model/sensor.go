package model

// Veri yapısı (struct)
type SensorData struct {
	Encoder1 int `json:"encoder1"`
	Encoder2 int `json:"encoder2"`
	Encoder3 int `json:"encoder3"`
	Encoder4 int `json:"encoder4"`
	Ax       int `json:"Ax"`
	Ay       int `json:"Ay"`
	Az       int `json:"Az"`
	Rx       int `json:"Rx"`
	Ry       int `json:"Ry"`
	Rz       int `json:"Rz"`
	Altitude int `json:"Altitude"`
	Temp     int `json:"Temp"`
	Ayak     int `json:"Ayaj"`
}
