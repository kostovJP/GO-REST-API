package routes

import (
	"example.com/REST-API/middleware"
	"github.com/gin-gonic/gin"
)

// takes the server pointer as input and registers the required routes.
func RegisterRoutes(server *gin.Engine) {
	authenticated := server.Group("/") 
	authenticated.Use(middleware.Authenticate)

	validatedUser := authenticated.Group("/")
	validatedUser.Use(middleware.CheckUser)
	
	// "/" is the base path. 
	// For eg: "/" + "/events" = "/events",
	// "/api" + "/events" = "/api/events" etc.

	// middleware execute order : authenticate -> checkUser -> handler..
	authenticated.POST("/events", createEvent) //protected 
	validatedUser.PUT("/events/:id", updateEvent) //protected + validUserCheck
	validatedUser.DELETE("/events/:id", deleteEvent) //protected + validUserCheck
	

	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent) 
	server.POST("/signup", signup)
	server.POST("/login", login) 
	
	//another way: 
	//server.POST("/events", middleware.Authenticate, createEvent)

	// executing takes place from left to write, so first the 
	// middleware handler will be executed.
	// we can do this for every protected route. 
}
