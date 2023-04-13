package postgres

import (
	"context"
	"testing"

	"online_banking_service/genproto/user_service"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
)

func createUser(t *testing.T) *user_service.CreateUserResponse {
	var (
		user = NewUserRepo(db)
	)

	res, err := user.Create(
		context.Background(),
		&user_service.CreateUserRequest{
			Username: gofakeit.Username(),
		},
	)

	assert.NoError(t, err)
	assert.NotEmpty(t, res)

	return res
}

func Test_CreateUser(t *testing.T) {
	res := createUser(t)

	assert.NotEmpty(t, res)
}
