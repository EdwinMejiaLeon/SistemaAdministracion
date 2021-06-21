package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

// get connection to data base
func GetConnection() *sql.DB {
	// nameUser:password@url:port/nameDataBase
	dsn := "postgres://postgres:12345@127.0.0.1:5432/sistemaRadioTaxis?sslmode=disable"
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
