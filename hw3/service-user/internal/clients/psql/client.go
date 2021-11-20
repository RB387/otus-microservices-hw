package psql

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
)

func NewFromEnv() (*sqlx.DB, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	connString := fmt.Sprintf(
		"dbname=%s user=%s password=%s host=%s sslmode=disable",
		dbname,
		user,
		password,
		host,
	)

	return sqlx.Connect("postgres", connString)
}
