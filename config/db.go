package config

import (
	_ "github.com/lib/pq"
	"database/sql"
	"fmt"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("postgres", "postgres://localhost/go_postgres?sslmode=disable")
	if err != nil {
		panic(err)
	}

	if err = DB.Ping(); err != nil {
		panic(err)
	}

	fmt.Println("Connected to Database.")
}
