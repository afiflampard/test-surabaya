package sensor

import (
	"time"

	"github.com/google/uuid"
)

type DeviceCommand struct {
	ID        uuid.UUID `db:"id" json:"id"`
	DeviceID  string    `db:"device_id" json:"device_id"`
	Command   string    `db:"command" json:"command"`
	Status    string    `db:"status" json:"status"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}

type DeviceCommandInput struct {
	DeviceID string `json:"device_id"`
	Command  string `json:"command"`
	Status   string `json:"status"`
}

func (d *DeviceCommand) TableName() string {
	return "device_commands"
}

func CreateNewDeviceCommand(input DeviceCommandInput) DeviceCommand {
	return DeviceCommand{
		ID:       uuid.New(),
		DeviceID: input.DeviceID,
		Command:  input.Command,
		Status:   input.Status,
	}
}
