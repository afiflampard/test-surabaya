package handler

import (
	"context"
	"log"

	"github.com/afif-musyayyidin/hertz-boilerplate/api/service"
	"github.com/afif-musyayyidin/hertz-boilerplate/domain/infra"
	"github.com/afif-musyayyidin/hertz-boilerplate/domain/sensor"
	"github.com/cloudwego/hertz/pkg/app"
)

type AppHandler struct {
	svc *service.Service
}

func NewAppHandler(svc *service.Service) *AppHandler {
	return &AppHandler{svc: svc}
}

// CreateSensorData godoc
// @Summary Create sensor data
// @Description Create a new sensor data entry
// @Tags sensor
// @Accept json
// @Produce json
// @Param data body sensor.SensorDataInput true "Sensor data"
// @Success 200 {object} uuid.UUID "Created sensor data ID"
// @Router /sensor [post]
func (h *AppHandler) CreateSensorData(ctx context.Context, c *app.RequestContext) {
	var input sensor.SensorDataInput
	if err := c.Bind(&input); err != nil {
		infra.JSONError(c, 400, "Bad Request", err)
		return
	}

	id, err := h.svc.CreateSensorData(ctx, &input)
	if err != nil {
		infra.JSONError(c, 500, "Internal Server Error", err)
		return
	}

	// Return the created ID in the response
	infra.JSONSuccess(c, id, "Sensor data created successfully")
}

// CreateDeviceControl godoc
// @Summary Create device control command
// @Description Create a new device control command
// @Tags sensor
// @Accept json
// @Produce json
// @Param command body sensor.DeviceCommandInput true "Device command"
// @Success 200 {object} uuid.UUID "Created command ID"
// @Router /device/control [post]
func (h *AppHandler) CreateDeviceControl(ctx context.Context, c *app.RequestContext) {
	var input sensor.DeviceCommandInput
	if err := c.Bind(&input); err != nil {
		infra.JSONError(c, 400, "Bad Request", err)
		return
	}

	id, err := h.svc.CreateDeviceControl(ctx, &input)
	if err != nil {
		log.Println(err)
		infra.JSONError(c, 500, "Internal Server Error", err)
		return
	}
	infra.JSONSuccess(c, id, "Device control command created successfully")
}
