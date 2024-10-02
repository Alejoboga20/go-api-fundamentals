package middlewares

import (
	"net/http"

	"github.com/Alejoboga20/go-api-fundamentals/utils"
	"github.com/gin-gonic/gin"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized.",
		})
		return
	}

	userId, err := utils.VerifyJWT(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Not authorized.",
		})
		return
	}

	context.Set("userId", userId)
	context.Next()
}
