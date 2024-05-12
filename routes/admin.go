package routes

import (
	"todo-service/controllers"
	"todo-service/middlewares"

	"github.com/gin-gonic/gin"
)

// Define todo routes
func AdminRoutes(router *gin.Engine) {
	router.GET("/admin/todos", middlewares.Verify, controllers.GetAllTodosAdmin)
	router.GET("/admin/users", middlewares.Verify, controllers.GetAllUsers)
	router.PUT("/admin/todos/:id", middlewares.Verify, controllers.UpdateTodoAdmin)
	router.DELETE("/admin/todos/:id", middlewares.Verify, controllers.DeleteTodoAdmin)
}
