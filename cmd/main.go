package main

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/moabdelazem/ecom/cmd/api"
	"github.com/moabdelazem/ecom/config"
	"github.com/moabdelazem/ecom/db"
)

func main() {
	// Create New MySQL Storage
	db, err := db.NewMySQLStorage(mysql.Config{
		User:                 config.Envs.DBUser,
		Passwd:               config.Envs.DBPass,
		Addr:                 config.Envs.DBAddr,
		DBName:               config.Envs.DBName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	// Handle the error if any occurs on creating the MySQL storage
	if err != nil {
		log.Fatal("Error creating MySQL storage")
	}

	// Initialize the storage
	initStorage(db)

	// Create a new server
	server := api.NewServer(":8080", db)

	// Start the server
	// Handle the error if any occurs on starting the server
	if err := server.Run(); err != nil {
		log.Fatal("Error starting the server")
	}
}

func initStorage(db *sql.DB) {
	err := db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database")
	}

	log.Println("Database is connected")
}
