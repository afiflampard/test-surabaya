package sensor

import (
	"time"

	"github.com/google/uuid"
)

type SensorData struct {
	ID          uuid.UUID `db:"id" json:"id"`
	DeviceID    string    `db:"device_id" json:"device_id"`
	Temperature float64   `db:"temperature" json:"temperature"`
	Humidity    float64   `db:"humidity" json:"humidity"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
}

type SensorDataInput struct {
	DeviceID    string  `json:"device_id"`
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func (s *SensorData) TableName() string {
	return "sensor_data"
}

func NewSensorDataInput(input SensorDataInput) *SensorData {
	return &SensorData{
		ID:          uuid.New(),
		DeviceID:    input.DeviceID,
		Temperature: input.Temperature,
		Humidity:    input.Humidity,
		CreatedAt:   time.Now(),
	}
}
