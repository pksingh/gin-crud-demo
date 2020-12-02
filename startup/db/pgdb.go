package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pksingh/gin-curd-demo/log"
	"go.uber.org/zap"
)

// Postgres - this var will be used in the rest of code to execute DB operations

type PgxIface interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	QueryRow(ctx context.Context, sql string, optionsAndArgs ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Close(context.Context) error
}

var Postgres PgxIface

// InitDB - initializes the global var Postgres
func InitDB(ctx context.Context) error {
	logger := log.GetLogger(ctx)
	// props := appProps.GetAll()

	host := "localhost"
	dbname := "test"
	port := "5432"
	user := "user"
	pwd := "password"
	sslmode := "disabled"

	logger.Info("connection string:", zap.String("host", host), zap.String("port", port),
		zap.String("user", user), zap.String("dbname", dbname), zap.String("sslmode", sslmode))

	fmt.Println("cleartextPwd PWD=", pwd)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pwd, dbname, sslmode)

	var err error
	Postgres, err = pgx.Connect(ctx, psqlInfo)
	if err != nil {
		logger.Error("sql.Open failed, error: ", zap.Error(err))
		return err
	}
	return nil

}
