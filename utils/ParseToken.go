package utils

import (
	"todo-service/models"

	"github.com/dgrijalva/jwt-go"
)

// Parses JWT token and returns Claims struct
func ParseToken(tokenString string) (*models.Claims, error) {
	// Parse JWT token
	token, err := jwt.ParseWithClaims(tokenString, &models.Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("my_secret_key"), nil
	})
	if err != nil {
		return nil, err
	}

	// Create Claims from token
	claims, ok := token.Claims.(*models.Claims)
	if !ok {
		return nil, err
	}
	return claims, nil
}
