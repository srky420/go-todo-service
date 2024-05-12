package controllers

import (
	"database/sql"
	"time"
	"todo-service/models"
	"todo-service/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my_secret_key")

func Login(c *gin.Context) {
	// Define body struct
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON body to userSignup struct
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Get user by email address
	user, err := models.GetUserByEmail(body.Email)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}

	// Compare provided and stored passwords
	hashErr := utils.CompareHashPassword(body.Password, user.Password)
	if hashErr != nil {
		c.JSON(400, gin.H{"error": "Invalid username or password"})
		return
	}

	// Token expiration time
	expiration := time.Now().Add(5 * time.Minute)

	// Create claims for the user
	claims := models.Claims{
		IsAdmin: user.IsAdmin,
		StandardClaims: jwt.StandardClaims{
			Subject:   user.Email,
			ExpiresAt: expiration.Unix(),
		},
	}

	// Generate jwt token for new user
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(400, gin.H{"message": "Error generating token"})
		return
	}

	// Set token in cookie
	c.SetCookie("token", tokenString, int(expiration.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "User logged in successfully!"})
}

func Signup(c *gin.Context) {
	// Define body struct
	var body struct {
		Username string `json:"username" binding:"required"`
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	// Bind JSON body to userSignup struct
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	// Check if user exists
	_, err := models.GetUserByEmail(body.Email)
	if err != sql.ErrNoRows {
		c.JSON(400, gin.H{"message": "User already exists"})
		return
	}

	// Insert new user to DB
	_, err = models.InsertUser(body.Username, body.Email, body.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": "Error creating new user"})
		return
	}

	c.JSON(201, gin.H{"message": "User created successfully!"})
}

func Admin(c *gin.Context) {

}

func Logout(c *gin.Context) {
	// Remove cookie
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"message": "User logged out successfully"})
}
