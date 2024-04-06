package postgresql

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"project/config"
	"time"
)

type Client interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

var DB *pgxpool.Pool

func NewClient(ctx context.Context, cfg config.Config) (pool *pgxpool.Pool, err error) {

	//var pool *pgxpool.Pool
	//var err error
	s := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Name)
	for i := 0; i < cfg.Database.MaxAttempts; i++ {
		time.Sleep(1 * time.Second)
		pool, err = pgxpool.New(ctx, s)
		if err != nil {
			fmt.Println("failed to connect to postgres")
		}

	}
	DB = pool
	return DB, err
}
