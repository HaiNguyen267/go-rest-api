package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hainguyen267/go-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)

	authenticatedRoutes := server.Group("/")
	authenticatedRoutes.Use(middlewares.Authenticate)
	authenticatedRoutes.POST("/events", createEvent) // protected route
	authenticatedRoutes.PUT("/events/:id", updateEvent) // protected route
	authenticatedRoutes.DELETE("/events/:id", deleteEvent) // protected route
	authenticatedRoutes.POST("/events/:id/register", registerForEvent) // protected route
	authenticatedRoutes.DELETE("/events/:id/cancel", cancelRegistration) // protected route

	server.POST("/signup", signup)
	server.POST("/login", login)
}