package main

import (
	"context"
	"log"
	"net"
	"time"

	grpcserver "github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/infrastructure/grpc"
	"github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/infrastructure/repository"
	"github.com/rafixcs/food-shopping-microservicess/services/user-service/internal/service"
	"github.com/rafixcs/food-shopping-microservicess/shared/env"
	"github.com/rafixcs/food-shopping-microservicess/shared/tracing"
	"google.golang.org/grpc"
)

var GrpcAddr = ":9093"

func main() {
	tracerCfg := tracing.Config{
		ServiceName:    "trip-service",
		Environment:    env.GetString("ENVIRONMENT", "development"),
		JaegerEndpoint: env.GetString("JAEGER_ENDPOINT", "http://jaeger:14268/api/traces"),
	}

	sh, err := tracing.InitTracer(tracerCfg)
	if err != nil {
		log.Fatalf("Failed to initialize the tracer: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	defer sh(ctx)

	repo := repository.NewInmemRepository()
	tokens := service.NewTokenManager("test", 60*time.Second)
	service := service.NewUserService(repo, tokens)

	lis, err := net.Listen("tcp", GrpcAddr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(tracing.WithTracingInterceptors()...)
	grpcserver.NewGrpcHandler(grpcServer, service)

	log.Printf("Starting gRPC server Trip service on port %s", lis.Addr())

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to serve: %v", err)
			cancel()
		}
	}()

	<-ctx.Done()
	grpcServer.GracefulStop()
}
