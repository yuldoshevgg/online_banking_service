package postgres

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/assert"

	config "online_banking_service/config"
)

var (
	db *pgxpool.Pool
)

func CreateRandomId(t *testing.T) string {
	id, err := uuid.NewRandom()
	assert.NoError(t, err)
	return id.String()
}

func TestMain(m *testing.M) {
	cfg := config.Load()
	fmt.Println("env loaded for test!")
	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDatabase,
	))
	if err != nil {
		panic(err)
	}

	conf.MaxConns = cfg.PostgresMaxConnections

	db, err = pgxpool.ConnectConfig(context.Background(), conf)
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}
