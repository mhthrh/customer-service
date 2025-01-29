package main

import (
	loader "github.com/mhthrh/common-lib/config"
	l "github.com/mhthrh/common-lib/config/logger"
	"go.uber.org/zap"
)

const (
	configPath = "/customer-service/file/config"
	configName = "config.json"
)

func main() {
	logger := zap.New(l.LogConfig())
	defer func() {
		_ = logger.Sync()
	}()

	sugar := logger.Sugar()
	sugar.Info("Loading config...")
	config := loader.NewFile(configName, configPath)
	if err := config.Initialize(); err != nil {
		sugar.Fatal("Failed to initialize config", zap.Error(err))
	}
}
