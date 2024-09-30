package routes

import (
	"net/http"

	"github.com/Alejoboga20/go-api-fundamentals/models"
	"github.com/gin-gonic/gin"
)

func singup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})
}
