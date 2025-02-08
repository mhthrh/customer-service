package main

import (
	"context"
	control "customer-service/control"
	"fmt"
	cError "github.com/mhthrh/GoNest/model/error"
	loader "github.com/mhthrh/GoNest/pkg/loader/file"
	l "github.com/mhthrh/GoNest/pkg/logger"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

const (
	configPath = "/customer-service/config"
	configName = "config.json"
)

var (
	osInterrupt       chan os.Signal
	listenerInterrupt chan *cError.XError
)

func init() {
	osInterrupt = make(chan os.Signal)
	listenerInterrupt = make(chan *cError.XError)
	signal.Notify(osInterrupt, os.Interrupt, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGHUP)
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
	go control.FillDbPool(context.Background(), *config)
	fmt.Println()
	select {
	case <-osInterrupt:
		sugar.Info("OS interrupt signal received")
	case e := <-listenerInterrupt:
		sugar.Infof("customer service listener interrupt signal received, %+v", e)
	}

}
