package main

import (
  "net/http"
  "github.com/gin-gonic/gin"
	"asteroids/models"
	"github.com/googollee/go-socket.io"
	"fmt"
)

func main() {
	r := gin.Default()
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected: ", s.ID())
		return nil
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	models.ConnectDatabase()
	r.Static("/", "./phaser")
  // r.GET("/", func(c *gin.Context) {
  //   c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  // })

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)

  r.Run()
}