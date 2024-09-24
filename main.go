package main

import (
	"net/http"

	gin "github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8080") // listen and serve on localhost:8080
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}
