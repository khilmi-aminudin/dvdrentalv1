package db

import (
	"context"

	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/khilmi-aminudin/dvdrentalv1/helper"
)

func ConnectDBWithPGX() *pgx.Conn {
	var (
		dbName   = os.Getenv("DB_NAME")
		dbUser   = os.Getenv("DB_USER")
		dbPass   = os.Getenv("DB_PASSWORD")
		dbHost   = os.Getenv("DB_HOST")
		dbPort   = os.Getenv("DB_PORT")
		dbDriver = os.Getenv("DB_DRIVER")
	)

	var connectionString = fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbDriver, dbUser, dbPass, dbHost, dbPort, dbName)
	connection, err := pgx.Connect(context.Background(), connectionString)
	helper.FatalError(err)
	return connection
}
