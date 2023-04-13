package service

import (
	"context"
	"online_banking_service/config"
	"online_banking_service/storage"

	us "online_banking_service/genproto/user_service"
	"online_banking_service/pkg/logger"
)

type userService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	us.UnimplementedUserServiceServer
}

func NewUserService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *userService {
	return &userService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (u *userService) Create(ctx context.Context, req *us.CreateUserRequest) (resp *us.CreateUserResponse, err error) {
	resp, err = u.strg.User().Create(ctx, req)
	if err != nil {
		u.log.Error("!!!Create User", logger.Error(err))
		return nil, err
	}

	return
}
