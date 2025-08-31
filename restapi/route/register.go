package route

import (
	"fmt"
	"net/http"
	"restapi/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func registerForEvents(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	fmt.Println(userId)
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id data"})
		return
	}

	event, err := model.GetEventByID(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not get event data"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Register Success"})
}

func cancelRegistration(ctx *gin.Context) {
	userId := ctx.GetInt64("userId")
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id data"})
		return
	}

	var event model.Event
	event.ID = eventId

	err = event.CancelRegis(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel register data"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Cancel Register Success"})

}
