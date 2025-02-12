package main

import (
	"context"
	"customer-service/control"
	"fmt"
	cError "github.com/mhthrh/GoNest/model/error"
	loader "github.com/mhthrh/GoNest/pkg/loader/file"
	l "github.com/mhthrh/GoNest/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
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
	ctx, cancel := context.WithCancel(context.Background())
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

	go control.Run(ctx, *config, listenerInterrupt)

	lis, e := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Grpc.Ip, config.Grpc.Port))
	if e != nil {
		log.Fatalf("Error listening on %s. error %v", config.Grpc.Ip, e)
	}
	rpcServer := grpc.NewServer()
	defer rpcServer.GracefulStop()

	go func() {
		if e = rpcServer.Serve(lis); e != nil {
			log.Fatalf("failed to serve: %v \n", e)
		}
	}()
	select {
	case <-osInterrupt:
		sugar.Info("OS interrupt signal received")
	case e := <-listenerInterrupt:
		sugar.Infof("customer service listener interrupt signal received, %+v", e)
	}

	sugar.Info("stopping customer service...")
	cancel()
	<-time.After(2 * time.Second)
}
