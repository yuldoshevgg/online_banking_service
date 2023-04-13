package client

import (
	config "online_banking_service/config"
	"online_banking_service/genproto/account_service"
	"online_banking_service/genproto/user_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceManagerI interface {
	UserService() user_service.UserServiceClient
	AccountService() account_service.AccountServiceClient
}

type grpcClients struct {
	userService    user_service.UserServiceClient
	accountService account_service.AccountServiceClient
}

func NewGrpcClients(cfg config.Config) (ServiceManagerI, error) {

	connService, err := grpc.Dial(
		cfg.GRPCHost+cfg.GRPCPort,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)

	if err != nil {
		return nil, err
	}

	return &grpcClients{
		userService:    user_service.NewUserServiceClient(connService),
		accountService: account_service.NewAccountServiceClient(connService),
	}, nil
}

func (g *grpcClients) UserService() user_service.UserServiceClient {
	return g.userService
}

func (g *grpcClients) AccountService() account_service.AccountServiceClient {
	return g.accountService
}
