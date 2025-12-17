package sensor

import (
	"context"

	"github.com/google/uuid"
)

type SensorRepository interface {
	CreateSensorData(ctx context.Context, data *SensorDataInput) uuid.UUID
	CreateDeviceControl(ctx context.Context, control *DeviceCommandInput) uuid.UUID
}
