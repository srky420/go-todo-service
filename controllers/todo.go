package controllers

import (
	"strconv"
	"todo-service/models"

	"github.com/gin-gonic/gin"
)

// Add todo
func GetAllTodos(c *gin.Context) {
	// Get user from middleware
	userId, _ := c.MustGet("id").(int64)
	isAdmin, _ := c.MustGet("isAdmin").(bool)

	// Get all todos for current user
	todos, err := models.GetAllTodos(userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error retreiving data"})
		return
	}
	c.JSON(200, gin.H{"id": userId, "isAdmin": isAdmin, "todos": todos})
}

func CreateTodo(c *gin.Context) {
	// Get user from middleware
	userId, _ := c.MustGet("id").(int64)

	// Req body struct
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// Bind JSON to struct
	if err := c.BindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Add todo
	_, err := models.AddTodo(body.Title, body.Description, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error creating todo"})
		return
	}
	c.JSON(201, gin.H{"message": "Todo created successfully"})
}

func UpdateTodo(c *gin.Context) {
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

	// Get user from middleware
	userId, _ := c.MustGet("id").(int64)

	// Get todo id from param and parse it to int64
	id := c.Param("id")
	todoId, convErr := strconv.ParseInt(id, 10, 64)
	if convErr != nil {
		c.JSON(500, gin.H{"error": "Error updating todo"})
		return
	}

	// Delete todo
	err := models.UpdateTodo(body.Title, body.Description, todoId, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error updating todo"})
		return
	}
	c.JSON(200, gin.H{"message": "Todo updated successfully"})
}

func DeleteTodo(c *gin.Context) {
	// Get user from middleware
	userId, _ := c.MustGet("id").(int64)

	// Get todo id from param and parse it to int64
	id := c.Param("id")
	todoId, convErr := strconv.ParseInt(id, 10, 64)
	if convErr != nil {
		c.JSON(500, gin.H{"error": "Error deleting todo"})
		return
	}

	// Delete todo
	err := models.RemoveTodo(todoId, userId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Error deleting todo"})
		return
	}
	c.JSON(200, gin.H{"message": "Todo deleted successfully"})
}
