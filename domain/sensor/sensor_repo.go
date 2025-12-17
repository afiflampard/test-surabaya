package sensor

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type SensorRepo struct {
	db *sqlx.DB
}

func NewSensorRepo(db *sqlx.DB) SensorRepository {
	return &SensorRepo{db: db}
}

// Implement repository methods here
func (r *SensorRepo) CreateSensorData(ctx context.Context, data *SensorDataInput) uuid.UUID {
	createSensor := NewSensorDataInput(*data)
	_, err := r.db.NamedExecContext(ctx, CreateSensorDataQuery, createSensor)
	if err != nil {
		log.Println(err)
		return uuid.Nil
	}
	return createSensor.ID
}

func (r *SensorRepo) CreateDeviceControl(ctx context.Context, control *DeviceCommandInput) uuid.UUID {
	createDeviceCommand := CreateNewDeviceCommand(*control)
	_, err := r.db.NamedExecContext(ctx, CreateDeviceControlQuery, createDeviceCommand)
	if err != nil {
		log.Println(err)
		return uuid.Nil
	}
	return createDeviceCommand.ID
}
