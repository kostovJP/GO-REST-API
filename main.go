package main

import (
	"net/http"

	"example.com/REST-API/models"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)
	server.GET("/events", createEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	//the second argument may be anything.
	context.JSON(http.StatusOK, gin.H{"events": events})
}

func createEvent(context *gin.Context) {
	var event models.Event

	// gin internally will scan the request body and then store the data
	// from that request body into the event object passed to ShouldBindJSON.
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "improper request fields"})
		return 
	}

	event.ID = 1 //a dummy ID for now
	event.UserId = 1

	//send back a response if everything ok:
	context.JSON(http.StatusCreated, gin.H{
		"message": "event created successfully",
		"event": event, //sending the newly created event as response
	})
}
