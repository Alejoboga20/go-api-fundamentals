package routes

import (
	"net/http"
	"strconv"

	"github.com/Alejoboga20/go-api-fundamentals/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id."})
		return
	}

	event, err := models.GetEventById(eventId)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Event not found."})
		return
	}

	err = event.Register(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register for event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Successfully registered for event."})
}

func cancelRegistration(context *gin.Context) {
	// userId := context.GetInt64("userId")
	// eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	// if err != nil {
	// 	context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid event id."})
	// 	return
	// }
}
