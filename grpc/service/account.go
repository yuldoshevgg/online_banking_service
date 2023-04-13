package service

import (
	"context"
	"errors"
	"online_banking_service/config"
	"online_banking_service/storage"

	as "online_banking_service/genproto/account_service"
	"online_banking_service/pkg/logger"

	"github.com/golang/protobuf/ptypes/empty"
)

type accountService struct {
	cfg  config.Config
	log  logger.LoggerI
	strg storage.StorageI
	as.UnimplementedAccountServiceServer
}

func NewAccountService(cfg config.Config, log logger.LoggerI, strg storage.StorageI) *accountService {
	return &accountService{
		cfg:  cfg,
		log:  log,
		strg: strg,
	}
}

func (c *accountService) CreateAccount(ctx context.Context, req *as.CreateAccountRequest) (resp *as.CreateAccountResponse, err error) {

	result, err := c.strg.Account().CheckAccountExist(
		ctx,
		&as.CheckAccountExistRequest{
			AccountNumber: req.GetAccountNumber(),
		},
	)

	if err != nil {
		c.log.Error("!!!Check account exist", logger.Error(err))
		return nil, err
	}

	if result.GetExist() {
		e := errors.New("Account already exist!")
		return nil, e
	}

	resp, err = c.strg.Account().CreateAccount(ctx, req)
	if err != nil {
		c.log.Error("!!!Create account", logger.Error(err))
		return nil, err
	}

	return
}

func (c *accountService) PayForAccount(ctx context.Context, req *as.PayForAccountRequest) (resp *empty.Empty, err error) {
	resp, err = c.strg.Account().PayForAccount(ctx, req)
	if err != nil {
		c.log.Error("!!!Pay for account", logger.Error(err))
		return nil, err
	}

	return
}

func (c *accountService) WithdrawFromAccount(ctx context.Context, req *as.WithdrawFromAccountRequest) (resp *empty.Empty, err error) {
	result, err := c.strg.Account().CheckAccountBalance(
		ctx,
		&as.CheckAccountBalanceRequest{
			AccountNumber: req.GetAccountNumber(),
			Balance:       req.GetBalance(),
		},
	)

	if err != nil {
		c.log.Error("!!!Check Account balance", logger.Error(err))
		return nil, err
	}

	if !result.GetExist() {
		e := errors.New("You don't have enough money to withdraw from the Account")
		return nil, e
	}

	resp, err = c.strg.Account().WithdrawFromAccount(ctx, req)
	if err != nil {
		c.log.Error("!!!Withdraw from Account", logger.Error(err))
		return nil, err
	}

	return
}

func (c *accountService) GetAccounts(ctx context.Context, req *as.GetAccountsRequest) (resp *as.GetAccountsResponse, err error) {
	resp, err = c.strg.Account().GetAccounts(ctx, req)
	if err != nil {
		c.log.Error("!!!Get Accounts", logger.Error(err))
		return nil, err
	}

	return
}

func (c *accountService) TransferBalance(ctx context.Context, req *as.TransferBalanceRequest) (resp *as.TransferBalanceResponse, err error) {
	result, err := c.strg.Account().CheckAccountBalance(
		ctx,
		&as.CheckAccountBalanceRequest{
			AccountNumber: req.GetSender(),
			Balance:       req.GetBalance(),
		},
	)

	if err != nil {
		c.log.Error("!!!Check Account balance", logger.Error(err))
		return nil, err
	}

	if !result.GetExist() {
		e := errors.New("You don't have enough money to transfer from the Account")
		return nil, e
	}

	resp, err = c.strg.Account().TransferBalance(ctx, req)
	if err != nil {
		c.log.Error("!!!Transfer Balance", logger.Error(err))
		return nil, err
	}

	return
}
