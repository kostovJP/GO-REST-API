package routes

import (
	"net/http"

	"example.com/REST-API/models"
	"example.com/REST-API/utils"
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

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(
			http.StatusBadRequest,
			gin.H{
				"message":"could not parse request data",
			},
		)
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(
			http.StatusUnauthorized,
			gin.H{
				"message":"invalid credentials ! Try again..",
			},
		)
		return
	}
	
	token, err := utils.GenerateJWT(user.Email, user.ID)

	if err != nil {
		context.JSON(
			http.StatusInternalServerError,
			gin.H{
				"message":"couldnot generate token..",
			},
		)

		return
	}

	context.JSON(
		http.StatusOK,
		gin.H{
			"message":"login successful",
			"token": token,
		},
	)
}