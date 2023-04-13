package main

import (
	"context"
	"net"

	api "online_banking_service/api"
	"online_banking_service/api/handlers"
	config "online_banking_service/config"
	grpc "online_banking_service/grpc"
	"online_banking_service/grpc/client"
	"online_banking_service/pkg/logger"
	"online_banking_service/storage/postgres"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	gin.SetMode(gin.ReleaseMode)
	log := logger.NewLogger(cfg.ServiceName, logger.LevelInfo)

	defer func() {
		if err := logger.Cleanup(log); err != nil {
			log.Error("Failed to cleanup logger", logger.Error(err))
		}
	}()

	pgStore, err := postgres.NewPostgres(context.Background(), cfg)
	if err != nil {
		log.Panic("postgres.NewPostgres", logger.Error(err))
	}

	defer pgStore.CloseDB()

	grpcServer := grpc.SetUpServer(cfg, log, pgStore)

	go func() {

		lis, err := net.Listen("tcp", cfg.GRPCPort)
		if err != nil {
			log.Panic("net.Listen", logger.Error(err))
		}

		log.Info("GRPC: Server being started...", logger.String("port", cfg.GRPCPort))

		if err := grpcServer.Serve(lis); err != nil {
			log.Panic("grpcServer.Serve", logger.Error(err))
		}
	}()

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(h, cfg)

	if r.Run(cfg.HTTPPort) != nil {
		panic(err)
	}
}
