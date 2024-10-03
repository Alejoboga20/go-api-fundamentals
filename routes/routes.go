package routes

import (
	"github.com/Alejoboga20/go-api-fundamentals/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEventById)

	protectedEventRoutes := server.Group("/events")
	protectedEventRoutes.Use(middlewares.Authenticate)
	protectedEventRoutes.POST("", createEvent)
	protectedEventRoutes.PUT("/:id", updateEvent)
	protectedEventRoutes.DELETE("/:id", deleteEvent)
	protectedEventRoutes.POST("/:id/register", registerForEvent)
	protectedEventRoutes.DELETE("/:id/register", cancelRegistration)

	server.POST("/users/register", register)
	server.POST("/users/login", login)
}
