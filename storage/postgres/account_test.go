package postgres

import (
	"context"
	"testing"

	as "online_banking_service/genproto/account_service"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func createAccount(t *testing.T) *as.CreateAccountResponse {
	var account = NewAccountRepo(db)
	user := createUser(t)

	res, err := account.CreateAccount(
		context.Background(),
		&as.CreateAccountRequest{
			AccountNumber: gofakeit.CreditCard().Number,
			UserId:        user.GetUserId(),
		},
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return res
}

func Test_CreateAccount(t *testing.T) {
	res := createAccount(t)

	assert.NotEmpty(t, res)
}

func Test_PayForAccount(t *testing.T) {
	var account = NewAccountRepo(db)
	res := createAccount(t)

	_, err := account.PayForAccount(
		context.Background(),
		&as.PayForAccountRequest{
			AccountNumber: res.GetAccountNumber(),
			Balance:       50,
		},
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func Test_WithdrawFromAccount(t *testing.T) {
	var account = NewAccountRepo(db)
	res := createAccount(t)

	_, err := account.WithdrawFromAccount(
		context.Background(),
		&as.WithdrawFromAccountRequest{
			AccountNumber: res.GetAccountNumber(),
			Balance:       50,
		},
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)
}

func Test_TransferBalance(t *testing.T) {
	var account = NewAccountRepo(db)
	sender := createAccount(t)
	recipient := createAccount(t)

	_, err := account.TransferBalance(
		context.Background(),
		&as.TransferBalanceRequest{
			Sender:    sender.GetAccountNumber(),
			Recipient: recipient.GetAccountNumber(),
			Balance:   10,
		},
	)

	assert.NoError(t, err)
}
