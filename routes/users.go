package routes

import (
	"net/http"

	"example.com/REST-API/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message":"couldnot parse JSON",
			},
		)

		return
	}

	err = user.Save()

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message":"User creating failed",
			},
		)

		return
	}

	context.JSON(
		http.StatusCreated,
		gin.H{
			"message":"New user created",
			"user": []any{user.ID, user.Email},
		},
	)
}