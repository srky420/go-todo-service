package models

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

// Initialize global DB var
var DB *sql.DB

// Initializes connection to DB and return DB variable
func InitDB(cfg mysql.Config) {

	// Attempts to connect to DB and
	// initialize db var
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// Check connection to DB
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// Assign global DB to the connection db var
	DB = db
}
