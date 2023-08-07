package main

import (
	"fmt"

	"csn-backend/internal/application/initiator"
	"csn-backend/internal/config"
	"csn-backend/pkg/logging"

	"go.uber.org/zap"
)

func main() {
	logger, errCreateZap := zap.NewProduction()
	if errCreateZap != nil {
		_ = fmt.Errorf("failed to create logger: %v", errCreateZap)
	}
	defer logger.Sync() //nolint:errcheck

	appLogger := logging.NewLogger(logger, "social_network")

	logger.Info("config initializing")

	cfg := config.GetConfig()
	app, errNewApp := initiator.NewApp(cfg, appLogger)
	if errNewApp != nil {
		appLogger.Fatal("error create app", zap.Error(errNewApp))
		return
	}

	logger.Info("running application")
	app.Run()
}
