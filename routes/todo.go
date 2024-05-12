package routes

import (
	"todo-service/controllers"
	"todo-service/middlewares"

	"github.com/gin-gonic/gin"
)

// Define todo routes
func TodoRoutes(router *gin.Engine) {
	router.POST("/todos", middlewares.Verify, controllers.CreateTodo)
	router.GET("/", middlewares.Verify, controllers.GetAllTodos)
	router.PUT("/todos/:id", middlewares.Verify, controllers.UpdateTodo)
	router.DELETE("/todos/:id", middlewares.Verify, controllers.DeleteTodo)
}
