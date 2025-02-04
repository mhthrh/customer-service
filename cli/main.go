package main

import (
	loader "github.com/mhthrh/common-lib/config/loader/file"
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

	config, err := loader.New(configPath, configName).Initialize()
	if err != nil {
		sugar.Fatal(err)
	}
	sugar.Info("customer service config loaded successfully")
	sugar.Info(config)
}
