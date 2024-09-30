package main

import (
	"net/http"

	"github.com/Alejoboga20/go-api-fundamentals/db"
	"github.com/Alejoboga20/go-api-fundamentals/models"
	gin "github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") // listen and serve on localhost:8080
}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()

	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	event.ID = len(models.GetAllEvents()) + 1
	event.UserID = 1
	event.Save()
	context.JSON(http.StatusCreated, event)
}
