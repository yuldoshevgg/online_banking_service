package service

import (
	"context"
	"errors"
	"fmt"

	config "online_banking_service/config"
	us "online_banking_service/genproto/user_service"
	"online_banking_service/pkg/logger"
	"online_banking_service/pkg/security"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (u *userService) Login(ctx context.Context, req *us.LoginRequest) (resp *us.LoginResponse, err error) {

	userID, err := u.strg.User().Login(ctx, req)
	if err != nil {
		u.log.Error("!!!Login--->", logger.Error(err))
		e := errors.New("Invalid username or password!")
		return nil, status.Error(codes.InvalidArgument, e.Error())
	}

	m := map[string]interface{}{
		"user_id":        userID,
		"role_id":        config.RoleID,
		"client_type_id": config.ClientTypeID,
	}

	token, err := security.GenerateJWT(m, config.AtExpireInTime, u.cfg.SecretKey)
	if err != nil {
		u.log.Error("!!!GenerateJWTToken--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token = fmt.Sprintf("Bearer %v", token)

	resp = &us.LoginResponse{
		Token:  token,
		UserId: userID,
	}

	return
}

func (u *userService) SignIn(ctx context.Context, req *us.SignInRequest) (resp *us.SignInResponse, err error) {

	result, err := u.strg.User().CheckUserExist(
		ctx,
		&us.CheckUserExistRequest{
			Username: req.GetUsername(),
		},
	)

	if err != nil {
		u.log.Error("!!!Check User Exist", logger.Error(err))
		return nil, err
	}

	if result.GetExist() {
		e := errors.New("User already exist!")
		return nil, e
	}

	user, err := u.strg.User().Create(
		ctx,
		&us.CreateUserRequest{
			Username: req.GetUsername(),
			Password: req.GetPassword(),
		},
	)
	if err != nil {
		u.log.Error("!!!Create User", logger.Error(err))
		return nil, err
	}

	m := map[string]interface{}{
		"user_id":        user.GetUserId(),
		"role_id":        config.RoleID,
		"client_type_id": config.ClientTypeID,
	}

	token, err := security.GenerateJWT(m, config.AtExpireInTime, u.cfg.SecretKey)
	if err != nil {
		u.log.Error("!!!GenerateJWTToken--->", logger.Error(err))
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	token = fmt.Sprintf("Bearer %v", token)
	resp = &us.SignInResponse{
		Token:  token,
		UserId: user.GetUserId(),
	}

	return
}
