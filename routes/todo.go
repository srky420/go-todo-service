package routes

import (
	"todo-service/controllers"

	"github.com/gin-gonic/gin"
)

// Define todo routes
func TodoRoutes(router *gin.Engine) {
	router.POST("/todo", controllers.Verify, controllers.CreateTodo)
	router.GET("/", controllers.Verify, controllers.GetAllTodos)
	router.PUT("/todo/:idTodo", controllers.Verify, controllers.UpdateTodo)
	router.DELETE("/todo/:idTodo", controllers.Verify, controllers.DeleteTodo)
}
