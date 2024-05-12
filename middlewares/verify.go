package middlewares

import (
	"todo-service/models"
	"todo-service/utils"

	"github.com/gin-gonic/gin"
)

func Verify(c *gin.Context) {
	// Get cookie token
	cookie, err := c.Cookie("token")
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not authorized"})
		return
	}

	// Parse token and get claims
	claims, err := utils.ParseToken(cookie)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not authorized"})
		return
	}

	// Get user from db using email
	email := claims.StandardClaims.Subject
	user, err := models.GetUserByEmail(email)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Not authorized"})
		return
	}

	// Set response for next
	c.Set("id", user.ID)
	c.Set("isAdmin", user.IsAdmin)
	c.Next()
}
