package controllers

import "github.com/gin-gonic/gin"

// Add todo
func GetAllTodos(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(200, gin.H{"user": user})
}

func CreateTodo(c *gin.Context) {

}

func UpdateTodo(c *gin.Context) {

}

func DeleteTodo(c *gin.Context) {

}
