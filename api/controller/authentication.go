package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pcorner10/my-finance/api/models"
	"github.com/pcorner10/my-finance/utility"
)

func Register(ctx *gin.Context) {
	var input models.User
	if err := ctx.ShouldBindJSON(&input); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedUser, err := input.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input models.User
	// create an interface to hold the input

	// bind the input to the interface

	if err := context.ShouldBindJSON(&input); err != nil {

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := &models.User{
		UserName: input.UserName,
	}

	user, err := user.GetUserByUsername()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := utility.GenerateJWT(*user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"jwt":      jwt,
		"fullname": user.Name + " " + user.LastName,
		"userid":   user.ID,
	})
}
