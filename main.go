package main

import (
	"context"
	"strconv"

	"github.com/afif-musyayyidin/hertz-boilerplate/api/router"
	"github.com/afif-musyayyidin/hertz-boilerplate/config"
	_ "github.com/afif-musyayyidin/hertz-boilerplate/docs"
	"github.com/afif-musyayyidin/hertz-boilerplate/domain/infra"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzSwagger "github.com/hertz-contrib/swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title           Kumparan API
// @version         1.0
// @description     This is the API documentation for Kumparan Project
// @host            localhost:8080
// @BasePath        /
func main() {
	cfg := config.LoadConfig()
	db := infra.InitPostgres(cfg)
	// dbReplica := infra.InitPostgresReplica(cfg)
	ctx := context.Background()
	// es := infra.ConnectElasticsearch(cfg)
	mqtt := infra.InitMQTT(cfg.MQTTBroker)
	port := ":" + strconv.Itoa(cfg.Port)
	h := server.Default(server.WithHostPorts(port))
	router.SetupRouter(ctx, h, db, mqtt)
	h.GET("/swagger/*any", hertzSwagger.WrapHandler(swaggerFiles.Handler))
	h.Spin()
}
