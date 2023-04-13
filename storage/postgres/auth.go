package postgres

import (
	"context"
	"fmt"

	us "online_banking_service/genproto/user_service"
)

func (u *userRepo) Login(ctx context.Context, req *us.LoginRequest) (resp string, err error) {
	resp = ""

	query := `
		SELECT
			guid
		FROM users
		WHERE username = $1 AND password = crypt($2, password)
	`

	row := u.db.QueryRow(
		ctx,
		query,
		req.GetUsername(),
		req.GetPassword(),
	)

	err = row.Scan(&resp)
	if err != nil {
		return "", fmt.Errorf("error while checking login: %v", err)
	}

	return
}
