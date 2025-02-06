package main

import (
	loader "github.com/mhthrh/common-lib/pkg/loader/file"
	l "github.com/mhthrh/common-lib/model/logger"
	cError "github.com/mhthrh/common-lib/model/error"
	"go.uber.org/zap"
	"os"
	"os/signal"
)

const (
	configPath = "/customer-service/file/config"
	configName = "config.json"
)

var (
	osInterrupt       chan os.Signal
	listenerInterrupt chan *cError.XError
)

func init() {
	osInterrupt = make(chan os.Signal)
	listenerInterrupt = make(chan *cError.XError)
	signal.Notify(osInterrupt, os.Interrupt)
}
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

	select {
	case <-osInterrupt:
		sugar.Info("OS interrupt signal received")
	case e := <-listenerInterrupt:
		sugar.Infof("customer service listener interrupt signal received, %+v", e)
	}

}
