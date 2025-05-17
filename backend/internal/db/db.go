package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

var DB *sql.DB

func Connect() error {
	var err error
	DB, err = sql.Open("postgres", os.Getenv("DB_SOURCE"))
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Successfully connected to DB")

	return nil
}
