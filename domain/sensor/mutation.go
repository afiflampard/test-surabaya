package sensor

import (
	"context"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SensorMutation interface {
	CreateSensorData(ctx context.Context, input *SensorDataInput) (uuid.UUID, error)
	CreateDeviceControl(ctx context.Context, input *DeviceCommandInput) (uuid.UUID, error)
	Publish(topic string, payload string) error
	IsConnected() bool
}

type sensorMutation struct {
	repo       SensorRepository
	db         *sqlx.DB
	mqtt       mqtt.Client
	validation SensorValidation
}

func NewSensorMutation(repo SensorRepository, db *sqlx.DB, mqttClient mqtt.Client) SensorMutation {
	return &sensorMutation{repo: repo, db: db, mqtt: mqttClient, validation: NewSensorValidation()}
}

func (m *sensorMutation) CreateSensorData(ctx context.Context, input *SensorDataInput) (uuid.UUID, error) {
	// Validate input
	if err := m.validation.ValidateSensorData(input); err != nil {
		return uuid.Nil, err
	}

	// Save to database
	id := m.repo.CreateSensorData(ctx, input)
	if id == uuid.Nil {
		return uuid.Nil, fmt.Errorf("failed to create sensor data")
	}

	// Publish to MQTT
	message := fmt.Sprintf("Temperature: %.1f, Humidity: %.1f", input.Temperature, input.Humidity)
	if err := m.Publish("sensor/data/"+input.DeviceID, message); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (m *sensorMutation) CreateDeviceControl(ctx context.Context, input *DeviceCommandInput) (uuid.UUID, error) {
	// Validate input
	if err := m.validation.ValidateDeviceCommand(input); err != nil {
		return uuid.Nil, err
	}

	// Save to database
	controlID := m.repo.CreateDeviceControl(ctx, input)
	if controlID == uuid.Nil {
		return uuid.Nil, fmt.Errorf("failed to create device control")
	}

	topic := fmt.Sprintf("greenhouse/control/%s", input.DeviceID)
	payload := fmt.Sprintf(`{"command":"%s"}`, input.Command)

	// Publish to MQTT
	if err := m.Publish(topic, payload); err != nil {
		return uuid.Nil, err
	}
	return controlID, nil
}

func (m *sensorMutation) Publish(topic string, payload string) error {
	token := m.mqtt.Publish(topic, 1, false, payload)
	token.Wait()
	return token.Error()
}

func (m *sensorMutation) IsConnected() bool {
	return m.mqtt != nil && m.mqtt.IsConnected()
}
