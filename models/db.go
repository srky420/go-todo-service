package models

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
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

	fmt.Println("MySql DB connection established")
}

// Create DB tables
func CreateTables() error {
	// Exec table creation queries
	_, userErr := DB.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
			username VARCHAR(128) UNIQUE NOT NULL,
			email VARCHAR(255) UNIQUE NOT NULL,
			password VARCHAR(255) NOT NULL,
			is_admin BOOL NOT NULL
		);
	`)
	if userErr != nil {
		return userErr
	}
	_, todoErr := DB.Exec(`
		CREATE TABLE IF NOT EXISTS task (
			id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
			title VARCHAR(255) NOT NULL,
			description TEXT NOT NULL,
			user_id INT NOT NULL,
			FOREIGN KEY (user_id) REFERENCES user(id)
		);
	`)
	if todoErr != nil {
		return todoErr
	}
	return nil
}

// Create admin row in DB
func CreateAdmin() error {
	// Insert Admin into the table if not exists
	_, err := GetUserByEmail(os.Getenv("ADMIN_EMAIL"))
	if err != nil {
		if err == sql.ErrNoRows {
			// Generate password hash for admin
			hashedPass, hashErr := bcrypt.GenerateFromPassword([]byte(os.Getenv("ADMIN_PASS")), bcrypt.DefaultCost)
			if hashErr != nil {
				return err
			}

			// Insert into admin into DB
			_, inserErr := DB.Exec(`
				INSERT INTO user
					(username, email, password, is_admin)
					VALUES
					(?, ?, ?, ?);
			`, os.Getenv("ADMIN_USER"), os.Getenv("ADMIN_EMAIL"), hashedPass, true)

			if inserErr != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}
