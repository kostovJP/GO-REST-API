package main

import (
	"net/http"
	"strconv"

	"example.com/REST-API/db"
	"example.com/REST-API/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)
	server.GET("/events/:id", getEvent)

	server.Run(":8080")

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Something went wrong while fetching events",
			},
		)
		return
	}
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

	//save the newly created event.
	err = event.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Something went wrong while saving event to database",
			},
		)
		return
	}

	//send back a response if everything ok:
	context.JSON(http.StatusCreated, gin.H{
		"message": "event created successfully",
		"event":   event, //sending the newly created event as response
	})
}

func getEvent(context *gin.Context) {
	// extracting the param named "id" from the request url, and then
	// converting it to an int64 value using strconv.ParseInt()
	evtID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message": "couldnot get event id",
			},
		)
		return
	}

	event, err := models.GetEventById(evtID)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "couldnot fetch event from database",
			},
		)
		return
	}

	// if everything goes well:
	context.JSON(
		http.StatusOK,
		gin.H{
			"message": "event found",
			"event":   event,
		},
	)
}
