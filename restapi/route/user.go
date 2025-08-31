package route

import (
	"net/http"
	"restapi/model"
	"restapi/utils"

	"github.com/gin-gonic/gin"
)

func signup(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not prase request data."})
		return
	}

	err = user.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save data."})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Create User Success."})
}

func login(ctx *gin.Context) {
	var user model.User

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not prase request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authen."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authen."})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Login Success.", "token": token})
}
