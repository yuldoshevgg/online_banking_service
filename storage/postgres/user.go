package postgres

import (
	"context"
	"fmt"

	us "online_banking_service/genproto/user_service"

	"github.com/jackc/pgx/v4/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (i *userRepo) Create(ctx context.Context, req *us.CreateUserRequest) (resp *us.CreateUserResponse, err error) {
	resp = &us.CreateUserResponse{}

	query := `
		INSERT INTO users (
			username,
			password
		) VALUES ($1, crypt($2, gen_salt('bf', 8)))
		RETURNING guid
	`

	row := i.db.QueryRow(
		ctx,
		query,
		req.GetUsername(),
		req.GetPassword(),
	)

	err = row.Scan(&resp.UserId)
	if err != nil {
		return nil, fmt.Errorf("error while insert user: %v", err)
	}

	return
}

func (i *userRepo) CheckUserExist(ctx context.Context, req *us.CheckUserExistRequest) (resp *us.CheckUserExistResponse, err error) {
	resp = &us.CheckUserExistResponse{}

	query := `
		SELECT
			EXISTS (
				SELECT guid FROM users WHERE username = $1
			)
	`

	row := i.db.QueryRow(
		ctx,
		query,
		req.GetUsername(),
	)

	err = row.Scan(&resp.Exist)
	if err != nil {
		return nil, fmt.Errorf("error while user exist: %v", err)
	}

	return
}
