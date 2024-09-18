package controllers

import (
	"net/http"
	"todo/database"
	"todo/database/models"

	"github.com/gin-gonic/gin"
)

func GetTodo (c *gin.Context) {
  id := c.Param("id") 
  todo := database.DB.First("id = ?", id)

  c.JSON(http.StatusOK, gin.H{"data": todo})
}

func GetTodos (c *gin.Context) {
  todo := database.DB.Find(&models.Todo{})

  c.JSON(http.StatusOK, gin.H{"data": todo})
}

