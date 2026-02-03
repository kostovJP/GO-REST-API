package routes

import (
	"net/http"
	"strconv"

	"example.com/REST-API/models"
	"github.com/gin-gonic/gin"
)

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

// protected
// the token should be part of the request header( Authorization header ).
// only a logged in user (a user having a valid token can create an event)
func createEvent(context *gin.Context) {
	var event models.Event

	// gin internally will scan the request body and then store the data
	// from that request body into the event object passed to ShouldBindJSON.
	// We are passing the address because we want the actual event object to be populated
	// with data not it's copy.
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "improper request fields"})
		return
	}

	// event id will be generated automatically, since we are creating a new event here,
	// and we have autoincrement set to event ID.
	// but we need to find the id of the user who created this event, and assign it to 
	// this event object so that it can be saved to the database.
	userId := context.GetInt64("userId") //getting the user id from context as we saved it there in the auth middleware.
	event.UserId = userId

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

func updateEvent(context *gin.Context) {
	// only the user who have created the event, should be able to update or delete that
	// event. The id of the user that is trying to update can be obtained from the 
	// context.Get("userId"). And we need to extract the id of the user mentioned in
	// the event.
	// eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	// if err != nil {
	// 	context.JSON(
	// 		http.StatusBadRequest,
	// 		gin.H{
	// 			"message":"id invalid!!",
	// 		},
	// 	)

	// 	return
	// }
	
	// event, err := models.GetEventById(eventID)

	// if err != nil {
	// 	context.JSON(
	// 		http.StatusInternalServerError,
	// 		gin.H{
	// 			"message":"Coudnot fetch event from database...",
	// 		},
	// 	)

	// 	return 
	// }

	// evtUser := event.UserId
	// user := context.GetInt64("userId")

	// if evtUser != user {
	// 	context.JSON(
	// 		http.StatusUnauthorized,
	// 		gin.H{
	// 			"message":"Not authorized",
	// 		},
	// 	)
	// 	return
	// }

	//in updateEvent() we are expecting a body along with the request.
	var updatedEvent models.Event
	err := context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message":"Couldnot parse event data...",
			},
		)

		return
	}

	eventID := context.GetInt64("eventId")
	// the id might not be included in the body, but we know that it has to be the 
	// id mentioned in the URL param
	updatedEvent.ID = eventID
	err = updatedEvent.Update()	

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message":"Couldnot update event..",
			},
		)

		return 
	}

	//success response:
	context.JSON(
		http.StatusOK,
		gin.H{
			"message":"Event updated successfully..",
			"event": updatedEvent,
		},
	)
}

func deleteEvent(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message":"Couldnot parse event data...",
			},
		)

		return
	}

	event, err := models.GetEventById(eventID)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message":"Couldnot find event with specified id..",
			},
		)

		return
	}

	//here we are not expecting any body, just the event id in the params.
	err = event.Delete()

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message":"couldnot delete event..",
			},
		)

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message":"Event deleted successfully...",
			"event": event,
		},
	)
}