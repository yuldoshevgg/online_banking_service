package grpc

import (
	config "online_banking_service/config"
	"online_banking_service/genproto/account_service"
	"online_banking_service/genproto/user_service"
	"online_banking_service/grpc/service"
	"online_banking_service/pkg/logger"
	storage "online_banking_service/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func SetUpServer(cfg config.Config, log logger.LoggerI, strg storage.StorageI) (grpcServer *grpc.Server) {
	grpcServer = grpc.NewServer()

	user_service.RegisterUserServiceServer(grpcServer, service.NewUserService(cfg, log, strg))
	account_service.RegisterAccountServiceServer(grpcServer, service.NewAccountService(cfg, log, strg))

	reflection.Register(grpcServer)
	return
}
