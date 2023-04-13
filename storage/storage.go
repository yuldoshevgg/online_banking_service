package storage

import (
	"context"

	as "online_banking_service/genproto/account_service"
	us "online_banking_service/genproto/user_service"

	"github.com/golang/protobuf/ptypes/empty"
)

type StorageI interface {
	CloseDB()
	User() UserRepoI
	Account() AccountRepoI
}

type UserRepoI interface {
	Create(ctx context.Context, req *us.CreateUserRequest) (resp *us.CreateUserResponse, err error)
	CheckUserExist(ctx context.Context, req *us.CheckUserExistRequest) (resp *us.CheckUserExistResponse, err error)
	Login(ctx context.Context, req *us.LoginRequest) (resp string, err error)
}

type AccountRepoI interface {
	CreateAccount(ctx context.Context, req *as.CreateAccountRequest) (resp *as.CreateAccountResponse, err error)
	PayForAccount(ctx context.Context, req *as.PayForAccountRequest) (resp *empty.Empty, err error)
	WithdrawFromAccount(ctx context.Context, req *as.WithdrawFromAccountRequest) (resp *empty.Empty, err error)
	GetAccounts(ctx context.Context, req *as.GetAccountsRequest) (resp *as.GetAccountsResponse, err error)
	CheckAccountBalance(ctx context.Context, req *as.CheckAccountBalanceRequest) (resp *as.CheckAccountBalanceResponse, err error)
	TransferBalance(ctx context.Context, req *as.TransferBalanceRequest) (resp *as.TransferBalanceResponse, err error)
	CheckAccountExist(ctx context.Context, req *as.CheckAccountExistRequest) (resp *as.CheckAccountExistResponse, err error)
}
