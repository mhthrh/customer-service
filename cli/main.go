package main

import (
	"context"
	"customer-service/control"
	"customer-service/pkg/grpcApi"
	"fmt"
	"github.com/mhthrh/GoNest/model/customer/grpc/customer"
	cError "github.com/mhthrh/GoNest/model/error"
	loader "github.com/mhthrh/GoNest/pkg/loader/file"
	l "github.com/mhthrh/GoNest/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
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
	internalInterrupt chan *cError.XError
)

func init() {
	osInterrupt = make(chan os.Signal)
	internalInterrupt = make(chan *cError.XError)
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
	sugar.Info("customer.v1 service config loaded successfully")

	go control.Run(ctx, *config, internalInterrupt)

	lis, e := net.Listen("tcp", fmt.Sprintf("%s:%d", config.Grpc.Ip, config.Grpc.Port))
	if e != nil {
		log.Fatalf("Error listening on %s. error %v", config.Grpc.Ip, e)
	}
	rpcServer := grpc.NewServer()
	defer rpcServer.GracefulStop()

	customer.RegisterCustomerServiceServer(
		rpcServer, &grpcApi.Customer{
			UnimplementedCustomerServiceServer: customer.UnimplementedCustomerServiceServer{},
		},
	)

	go func() {
		reflection.Register(rpcServer)
		if e = rpcServer.Serve(lis); e != nil {
			log.Fatalf("failed to serve: %v \n", e)
		}
	}()
	sugar.Info("gRPC server started on port ", config.Grpc.Port)

	sugar.Info("service listening for any interrupt signals...")
	select {
	case <-osInterrupt:
		sugar.Info("OS interrupt signal received")
	case e := <-internalInterrupt:
		sugar.Errorf("customer.v1 service listener interrupt signal received, %+v", e)
	}

	sugar.Info("stopping customer.v1 service...")
	cancel()

	<-internalInterrupt
}
