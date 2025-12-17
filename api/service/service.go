package service

import (
	"context"

	"github.com/afif-musyayyidin/hertz-boilerplate/domain/sensor"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type Service struct {
	db         *sqlx.DB
	mqtt       mqtt.Client
	sensorRepo sensor.SensorRepository
}

func NewService(ctx context.Context, db *sqlx.DB, mqttClient mqtt.Client, sensorRepo sensor.SensorRepository) *Service {
	return &Service{
		db:         db,
		mqtt:       mqttClient,
		sensorRepo: sensorRepo,
	}
}

func (s *Service) CreateSensorData(ctx context.Context, input *sensor.SensorDataInput) (uuid.UUID, error) {
	mutation := sensor.NewSensorMutation(s.sensorRepo, s.db, s.mqtt)
	id, err := mutation.CreateSensorData(ctx, input)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (s *Service) CreateDeviceControl(ctx context.Context, input *sensor.DeviceCommandInput) (uuid.UUID, error) {
	mutation := sensor.NewSensorMutation(s.sensorRepo, s.db, s.mqtt)
	id, err := mutation.CreateDeviceControl(ctx, input)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}
