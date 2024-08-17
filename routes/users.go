package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hainguyen267/go-rest-api/models"
	"github.com/hainguyen267/go-rest-api/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}


	err = user.Save()

	fmt.Println("User :", user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save the user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User signup successfully",
	})
}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data",
		})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})	
		return
	}


	jwt, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not authenticate user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successfully!",
		"token": jwt,
	})


}

