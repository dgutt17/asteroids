package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"asteroids/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()
	r.Static("/", "./phaser")
  // r.GET("/", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  // })

  r.Run()
}