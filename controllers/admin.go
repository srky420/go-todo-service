package controllers

import (
	"strconv"
	"todo-service/models"

	"github.com/gin-gonic/gin"
)

// Get all todos
func GetAllTodosAdmin(c *gin.Context) {
	// Get role from middleware
	isAdmin, _ := c.MustGet("isAdmin").(bool)

	// Check if user is admin
	if !isAdmin {
		c.JSON(401, gin.H{"error": "Not Authorized"})
		return
	}

	// Get all todos
	todos, err := models.GetAllTodosAdmin()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error retreiving todos"})
		return
	}
	c.JSON(200, todos)
}

func GetAllUsers(c *gin.Context) {
	// Get role from middleware
	isAdmin, _ := c.MustGet("isAdmin").(bool)

	// Check if user is admin
	if !isAdmin {
		c.JSON(401, gin.H{"error": "Not Authorized"})
		return
	}

	// Get all todos
	users, err := models.GetAllUsersAdmin()
	if err != nil {
		c.JSON(500, gin.H{"error": "Error retreiving todos"})
		return
	}
	c.JSON(200, users)
}

func UpdateTodoAdmin(c *gin.Context) {
	// Get role from middleware
	isAdmin, _ := c.MustGet("isAdmin").(bool)

	// Check if user is admin
	if !isAdmin {
		c.JSON(401, gin.H{"error": "Not Authorized"})
		return
	}

	// Create body struct
	var body struct {
		Title       string `json:"title" binding:"required"`
		Description string `json:"description" binding:"required"`
	}

	// Bind body to JSON
	if err := c.BindJSON(&body); err != nil {
		c.JSON(500, err.Error())
		return
	}

	// Get todo id from param and parse it to int64
	id := c.Param("id")
	todoId, convErr := strconv.ParseInt(id, 10, 64)
	if convErr != nil {
		c.JSON(500, gin.H{"error": "Error updating todo"})
		return
	}

	// Update todo using id
	err := models.UpdateTodoAdmin(body.Title, body.Description, todoId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating todo"})
		return
	}
	c.JSON(200, gin.H{"message": "Todo updated succesfully"})
}

func DeleteTodoAdmin(c *gin.Context) {
	// Get role from middleware
	isAdmin, _ := c.MustGet("isAdmin").(bool)

	// Check if user is admin
	if !isAdmin {
		c.JSON(401, gin.H{"error": "Not Authorized"})
		return
	}

	// Get todo id from param and parse it to int64
	id := c.Param("id")
	todoId, convErr := strconv.ParseInt(id, 10, 64)
	if convErr != nil {
		c.JSON(500, gin.H{"error": "Error deleting todo"})
		return
	}

	err := models.RemoveTodoAdmin(todoId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting todo"})
		return
	}
	c.JSON(200, gin.H{"message": "Todo deleted succesfully"})
}
