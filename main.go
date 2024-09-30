package main

import (
	"github.com/Alejoboga20/go-api-fundamentals/db"
	"github.com/Alejoboga20/go-api-fundamentals/routes"
	gin "github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") // listen and serve on localhost:8080
}
