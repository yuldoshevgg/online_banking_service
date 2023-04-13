package postgres

import (
	"context"
	"fmt"

	as "online_banking_service/genproto/account_service"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jackc/pgx/v4/pgxpool"
	"google.golang.org/protobuf/types/known/emptypb"
)

type accountRepo struct {
	db *pgxpool.Pool
}

func NewAccountRepo(db *pgxpool.Pool) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (i *accountRepo) CreateAccount(ctx context.Context, req *as.CreateAccountRequest) (resp *as.CreateAccountResponse, err error) {
	resp = &as.CreateAccountResponse{}

	query := `
		INSERT INTO accounts (
			account_number,
			user_id
		) VALUES ($1, (SELECT id FROM users WHERE guid = $2))
		RETURNING guid, account_number
	`

	row := i.db.QueryRow(
		ctx,
		query,
		req.GetAccountNumber(),
		req.GetUserId(),
	)

	err = row.Scan(&resp.AccountId, &resp.AccountNumber)
	if err != nil {
		return nil, fmt.Errorf("error while creating account: %v", err)
	}

	return
}

func (i *accountRepo) PayForAccount(ctx context.Context, req *as.PayForAccountRequest) (resp *empty.Empty, err error) {

	query := `
		UPDATE accounts SET 
			balance = balance + $2
		WHERE account_number = $1
	`

	_, err = i.db.Exec(
		ctx,
		query,
		req.GetAccountNumber(),
		req.GetBalance(),
	)

	if err != nil {
		return nil, fmt.Errorf("error while update account balance: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (i *accountRepo) WithdrawFromAccount(ctx context.Context, req *as.WithdrawFromAccountRequest) (resp *empty.Empty, err error) {

	query := `
		UPDATE accounts SET 
			balance = balance - $2
		WHERE account_number = $1
	`

	_, err = i.db.Exec(
		ctx,
		query,
		req.GetAccountNumber(),
		req.GetBalance(),
	)

	if err != nil {
		return nil, fmt.Errorf("error while update account balance: %v", err)
	}

	return &emptypb.Empty{}, nil
}

func (i *accountRepo) CheckAccountBalance(ctx context.Context, req *as.CheckAccountBalanceRequest) (resp *as.CheckAccountBalanceResponse, err error) {
	resp = &as.CheckAccountBalanceResponse{}

	query := `
		SELECT 
			(
				SELECT 
					balance 
				FROM accounts 
				WHERE account_number = $1
			) >= $2
	`

	err = i.db.QueryRow(
		ctx,
		query,
		req.GetAccountNumber(),
		req.GetBalance(),
	).Scan(&resp.Exist)

	if err != nil {
		return nil, fmt.Errorf("error while update account balance: %v", err)
	}

	return
}

func (i *accountRepo) GetAccounts(ctx context.Context, req *as.GetAccountsRequest) (resp *as.GetAccountsResponse, err error) {
	resp = &as.GetAccountsResponse{}

	query := `
		SELECT
			account_number,
			balance,
			TO_CHAR(created_at, 'YYYY-MM-DD HH24:MI:SS')
		FROM accounts
		WHERE user_id = (
			SELECT id FROM users WHERE username = $1
		) AND deleted_at IS NULL
	`

	rows, err := i.db.Query(ctx, query, req.GetUsername())
	if err != nil {
		return nil, fmt.Errorf("error while getting user accounts: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var c as.GetAccounts

		err = rows.Scan(
			&c.AccountNumber,
			&c.Balance,
			&c.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("error while scanning: %v", err)
		}

		resp.Accounts = append(resp.Accounts, &c)
	}

	return
}

func (i *accountRepo) TransferBalance(ctx context.Context, req *as.TransferBalanceRequest) (*as.TransferBalanceResponse, error) {
	resp := &as.TransferBalanceResponse{}

	tx, err := i.db.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("Failed to start transaction")
	}

	defer func() {
		if err != nil {
			err = tx.Rollback(ctx)
			if err != nil {
				return
			}
		} else {
			err = tx.Commit(ctx)
			if err != nil {
				return
			}
		}
	}()

	up1 := `
		UPDATE accounts SET
			balance = balance - $2
		WHERE account_number = $1
	`

	_, err = tx.Exec(
		ctx,
		up1,
		req.GetSender(),
		req.GetBalance(),
	)

	if err != nil {
		return nil, fmt.Errorf("error while updating sender balance: %v", err)
	}

	up2 := `
		UPDATE accounts SET
			balance = balance + $2
		WHERE account_number = $1
	`

	_, err = tx.Exec(
		ctx,
		up2,
		req.GetRecipient(),
		req.GetBalance(),
	)

	if err != nil {
		return nil, fmt.Errorf("error while updating recipient balance: %v", err)
	}

	query := `
		INSERT INTO transactions (
			sender,
			recipient,
			amount
		) VALUES ($1, $2, $3)
		RETURNING guid
	`

	err = tx.QueryRow(
		ctx,
		query,
		req.GetSender(),
		req.GetRecipient(),
		req.GetBalance(),
	).Scan(&resp.TransactionId)

	if err != nil {
		return nil, fmt.Errorf("error while inserting transaction: %v", err)
	}

	return resp, nil
}

func (i *accountRepo) CheckAccountExist(ctx context.Context, req *as.CheckAccountExistRequest) (resp *as.CheckAccountExistResponse, err error) {
	resp = &as.CheckAccountExistResponse{}

	query := `
		SELECT
			EXISTS (
				SELECT guid FROM accounts WHERE account_number = $1
			)
	`

	row := i.db.QueryRow(
		ctx,
		query,
		req.GetAccountNumber(),
	)

	err = row.Scan(&resp.Exist)
	if err != nil {
		return nil, fmt.Errorf("error while checking account exist: %v", err)
	}

	return
}
