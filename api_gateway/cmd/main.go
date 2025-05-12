package main

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/api_gateway/config"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/app"
)

func main() {
	ctx := context.Background()
	// TODO: add telemetry here when the topic of logging will be covered

	// Parse config
	cfg, err := config.New()
	if err != nil {
		logger.GetLogger().Println("failed to parse config: %v", err)

		return
	}

	application, err := app.New(ctx, cfg)
	if err != nil {
		logger.GetLogger().Println("failed to setup application:", err)

		return
	}

	err = application.Run()
	if err != nil {
		logger.GetLogger().Println("failed to run application: ", err)

		return
	}
}
