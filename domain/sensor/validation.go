package sensor

import "fmt"

type SensorValidation interface {
	ValidateSensorData(input *SensorDataInput) error
	ValidateDeviceCommand(input *DeviceCommandInput) error
}

type sensorValidation struct{}

func NewSensorValidation() SensorValidation {
	return &sensorValidation{}
}

func (s *sensorValidation) ValidateSensorData(input *SensorDataInput) error {
	if input.DeviceID == "" {
		return fmt.Errorf("device_id is required")
	}
	if input.Temperature < -273.15 {
		return fmt.Errorf("temperature cannot be below absolute zero")
	}
	if input.Humidity < 0 || input.Humidity > 100 {
		return fmt.Errorf("humidity must be between 0 and 100")
	}
	return nil
}

func (s *sensorValidation) ValidateDeviceCommand(input *DeviceCommandInput) error {
	if input.DeviceID == "" {
		return fmt.Errorf("device_id is required")
	}
	if input.Command == "" {
		return fmt.Errorf("command is required")
	}
	return nil
}
