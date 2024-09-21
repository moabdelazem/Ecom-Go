package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(cfg mysql.Config) (*sql.DB, error) {
	// Open the connection
	db, err := sql.Open("mysql", cfg.FormatDSN())
	// Handle the error if any
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
