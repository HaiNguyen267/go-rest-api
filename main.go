package main

import (


	"github.com/gin-gonic/gin"
	"github.com/hainguyen267/go-rest-api/db"
	"github.com/hainguyen267/go-rest-api/routes"
)


func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}

