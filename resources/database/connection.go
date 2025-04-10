package database

import (
	"context"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

var (
	conn *pgxpool.Pool
)

type Pgx interface {
	Exec(context context.Context, sql string, args ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(context context.Context, sql string, args ...interface{}) (rows pgx.Rows, err error)
	QueryRow(context context.Context, sql string, args ...interface{}) pgx.Row
	Begin(context context.Context) (pgx.Tx, error)
	Close()
}

func Connect() Pgx {
	if conn != nil {
		return conn
	}

	err := error(nil)
	// 'db' = docker | 'localhost' = local
	conn, err = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))

	conn.Config().MaxConns = 50
	conn.Config().MinConns = 40
	conn.Config().MaxConnIdleTime = time.Minute * 3

	if err != nil {
		panic(err)
	}

	log.Print("$$$$Connected to database")

	return conn
}
