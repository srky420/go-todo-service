package routes

import (
	"todo-service/controllers"

	"github.com/gin-gonic/gin"
)

// Define todo routes
func TodoRoutes(router *gin.Engine) {
	router.POST("/todos", controllers.Verify, controllers.CreateTodo)
	router.GET("/", controllers.Verify, controllers.GetAllTodos)
	router.PUT("/todos/:id", controllers.Verify, controllers.UpdateTodo)
	router.DELETE("/todos/:id", controllers.Verify, controllers.DeleteTodo)
}
