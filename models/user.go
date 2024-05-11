package models

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// User struct
type User struct {
	ID       int64
	Username string
	Email    string
	Password string
	IsAdmin  bool
}

// Define jwt claims
type Claims struct {
	IsAdmin bool
	jwt.StandardClaims
}

// Add a user to the DB
func InsertUser(username string, email string, password string) (int64, error) {
	// Check if user already exists
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}

	// Insert the user into DB
	result, err := DB.Exec("INSERT INTO user (username, email, password) VALUES (?, ?, ?)", username, email, hashedPass)
	if err != nil {
		return 0, err
	}

	// Get the ID of inserted user
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// Get user by id
func GetUserByEmail(email string) (User, error) {
	// Declare album var
	var user User

	// Query DB using id to get album
	row := DB.QueryRow("SELECT * FROM user WHERE email = ?", email)

	// Check for field errors
	if err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.IsAdmin); err != nil {
		return user, err
	}
	return user, nil
}
