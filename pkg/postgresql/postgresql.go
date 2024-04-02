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

//type DsnConfig struct {
//	Host        string
//	Port        string
//	User        string
//	Password    string
//	Database    string
//	MaxAttempts int
//}

type Client interface {
	Exec(context.Context, string, ...interface{}) (pgconn.CommandTag, error)
	Query(context.Context, string, ...interface{}) (pgx.Rows, error)
	QueryRow(context.Context, string, ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

func NewClient(ctx context.Context, cfg config.DataBaseconfig) (pool *pgxpool.Pool, err error) {

	//var pool *pgxpool.Pool
	//var err error
	s := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)

	for i := 0; i < cfg.MaxAttempts; i++ {
		time.Sleep(5 * time.Second)
		pool, err = pgxpool.New(ctx, s)
		if err != nil {
			fmt.Println("failed to connect to postgres")
		}

	}

	return pool, err
}
