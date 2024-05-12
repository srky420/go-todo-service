package utils

import "golang.org/x/crypto/bcrypt"

// Compares string password and password hash
func CompareHashPassword(password, hash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err
}
