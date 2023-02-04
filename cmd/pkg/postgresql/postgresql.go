package postgresql

import (
	"context"
	"fmt"
	"github.com/Zavr22/goLangStudy/cmd/pkg/utills"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4"
	pgxpool "github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Client interface {
	Exec(sql string, arguments ...interface{}) pgconn.CommandTag
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(sql string, arguments ...interface{}) (pgx.Row, error)
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, maxAttempts int, username, password, host, port, database string) (con *pgxpool.Pool, err error) {

	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", username, password, host, port, database)
	err = repeatable.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		con, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttempts, 3*time.Second)

	if err != nil {

		log.Fatal("error while connect")
	}

	return con, nil
}
