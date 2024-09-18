package main

import (
	"net/http"
	"todo/controllers"
	"todo/database"

	"github.com/gin-gonic/gin"
)

func main() {
  g := gin.Default()

  todoGroup := g.Group("/todo")

  g.GET("/ping", func (c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Pong"})
  }) 

  todoGroup.GET("/", controllers.GetTodos)
  todoGroup.GET("/:id", controllers.GetTodo)

  // TODO: add .env setup
  database.Setup("")


  g.Run(":8000")
}
