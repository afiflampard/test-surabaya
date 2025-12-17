package router

import (
	"context"
	"time"

	"github.com/afif-musyayyidin/hertz-boilerplate/api/handler"
	"github.com/afif-musyayyidin/hertz-boilerplate/api/service"
	"github.com/afif-musyayyidin/hertz-boilerplate/domain/infra"
	"github.com/afif-musyayyidin/hertz-boilerplate/domain/sensor"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(ctx context.Context, h *server.Hertz, db *sqlx.DB, mqttClient mqtt.Client) {

	sensorRepo := sensor.NewSensorRepo(db)

	svc := service.NewService(ctx, db, mqttClient, sensorRepo)
	handler := handler.NewAppHandler(svc)

	sensor := h.Group("/sensor")
	{
		sensor.POST("/create-sensor", handler.CreateSensorData)
		sensor.POST("/create-control", handler.CreateDeviceControl)
	}
	health := h.Group("/health")
	{
		health.GET("", func(c context.Context, ctx *app.RequestContext) {

			response := map[string]interface{}{
				"status":    "ok",
				"service":   "sensor-core",
				"timestamp": time.Now().UTC().Format(time.RFC3339),
				"checks":    map[string]interface{}{},
			}

			httpStatus := 200
			checks := response["checks"].(map[string]interface{})

			// ✅ MQTT Check
			if !mqttClient.IsConnected() {
				checks["mqtt"] = map[string]interface{}{
					"status": "down",
					"error":  "MQTT connection lost",
				}
				response["status"] = "degraded"
				httpStatus = 503
			} else {
				checks["mqtt"] = map[string]interface{}{
					"status": "up",
				}
			}

			// ✅ PostgreSQL Check
			if err := infra.CheckPostgresConnection(db); err != nil {
				checks["database"] = map[string]interface{}{
					"status": "down",
					"error":  err.Error(),
				}
				response["status"] = "degraded"
				httpStatus = 503
			} else {
				checks["database"] = map[string]interface{}{
					"status": "up",
				}
			}

			ctx.JSON(httpStatus, response)
		})
	}
}
