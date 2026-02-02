package routes

import "github.com/gin-gonic/gin"

// takes the server pointer as input and registers the required routes.
func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)
}
