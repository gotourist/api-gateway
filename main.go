package main

import (
	"github.com/iman_task/api-gateway/api"
	configPkg "github.com/iman_task/api-gateway/config"
	"github.com/iman_task/api-gateway/pkg/logger"
	"github.com/iman_task/api-gateway/services"
)

func main() {
	// =========================================================================
	// Configurations loading...
	config := configPkg.Load()

	// =========================================================================
	// Logger
	log := logger.New("debug", "api-gateway")
	defer func() {
		err := logger.Cleanup(log)
		log.Error("failed to cleanup logs", logger.Error(err))
	}()

	// =========================================================================
	// gRPC clients
	serviceManager, err := services.NewServiceManager(&config)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}

	server := api.New(api.Option{
		Conf:           config,
		Logger:         log,
		ServiceManager: serviceManager,
	})

	if err := server.Run(config.HttpPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
