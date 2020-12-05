package db

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/pksingh/gin-curd-demo/log"
	"github.com/pksingh/gin-curd-demo/startup/appProps"
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
	props := appProps.GetAll()

	host, ok := os.LookupEnv("DATABASE_HOST")
	if !ok {
		host = props.GetString("postgresql.host", "")
	}
	dbname, ok := os.LookupEnv("DATABASE_NAME")
	if !ok {
		dbname = props.GetString("postgresql.dbname", "")
	}
	port, ok := os.LookupEnv("DATABASE_PORT")
	if !ok {
		port = strconv.Itoa(props.GetInt("postgresql.port", 0))
	}
	user, ok := os.LookupEnv("DATABASE_USER")
	if !ok {
		user = props.GetString("postgresql.user", "")
	}
	cleartextPwd, ok := os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		cleartextPwd = props.GetString("postgresql.pwd", "")
	}

	sslmode, ok := os.LookupEnv("DATABASE_SSLMODE")
	if !ok {
		sslmode = props.GetString("postgresql.sslmode", "")
	}

	logger.Info("connection string:", zap.String("host", host), zap.String("port", port),
		zap.String("user", user), zap.String("dbname", dbname), zap.String("sslmode", sslmode))

	fmt.Println("cleartextPwd PWD=", cleartextPwd)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, cleartextPwd, dbname, sslmode)

	var err error
	Postgres, err = pgx.Connect(ctx, psqlInfo)
	if err != nil {
		logger.Error("sql.Open failed, error: ", zap.Error(err))
		return err
	}
	return nil

}
