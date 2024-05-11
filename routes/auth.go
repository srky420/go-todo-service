package routes

import (
	"todo-service/controllers"

	"github.com/gin-gonic/gin"
)

// Defines auth routes
func AuthRoutes(router *gin.Engine) {
	// Create router configuration
	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/admin", controllers.Admin)
	router.GET("/home", controllers.Home)
}
