package middleware

import (
	"net/http"
	"strconv"

	"example.com/REST-API/models"
	"example.com/REST-API/utils"
	"github.com/gin-gonic/gin"
)

// This is a middleware and will be executed in between requests. So, other
// request handler after this one would still run.
func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	// since this is a middleware, so we cannot just send context.JSON as then it
	// will mean that we will be sending multiple responses. We want to abort if something
	// goes wrong here and not proceed to the next handler. That's why we are using
	// context.AbortWithStatusJSON() instead of just JSON()
	if token == "" {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Not authorized! Token not found....",
			},
		)

		return
	}

	userId, err := utils.VerifyJWT(token)

	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Not authorized! Token not found....",
			},
		)

		return
	}
	// since we need to somehow return the userId, we will attach it to context instead.
	// as the same context is used everywhere.
	context.Set("userId", userId)
	context.Next() // proceed to the next handler.
}

// Used to check if the user deleting or updating an event is the same user that has
// created the event.
func CheckUser(context *gin.Context) {
	eventID, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"message": "id invalid!!",
			},
		)

		return
	}

	event, err := models.GetEventById(eventID)

	if err != nil {
		context.AbortWithStatusJSON(
			http.StatusInternalServerError,
			gin.H{
				"message": "Coudnot fetch event from database...",
			},
		)

		return
	}

	evtUser := event.UserId
	user := context.GetInt64("userId")

	if evtUser != user {
		context.AbortWithStatusJSON(
			http.StatusUnauthorized,
			gin.H{
				"message": "Not authorized",
			},
		)
		return
	}

	context.Set("eventId", eventID)
	context.Next()
}
