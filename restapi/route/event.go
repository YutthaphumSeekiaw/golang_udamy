package route

import (
	"fmt"
	"net/http"
	"restapi/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(ctx *gin.Context) {
	events, err := model.GetAllEvents()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not query data"})
		return
	}
	ctx.JSON(http.StatusOK, events)
}

func getEvent(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not prase Id"})
		return
	}

	event, err := model.GetEventByID(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not prase get data by id"})
		return
	}

	ctx.JSON(http.StatusOK, event)

}

func createEvent(ctx *gin.Context) {
	var event model.Event
	if err := ctx.ShouldBindJSON(&event); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserId = userId

	err := event.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not create data"})
		return
	}

	ctx.JSON(http.StatusCreated, event)
}

func updateEvents(ctx *gin.Context) {
	var event model.Event
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not prase Id"})
		return
	}

	eventdata, err := model.GetEventByID(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not prase get data by id"})
		return
	}

	userId := ctx.GetInt64("userId")
	event.UserId = userId
	if eventdata.UserId != userId {
		ctx.JSON(http.StatusUnauthorized, gin.H{"Message": "Not Autrorize"})
		return
	}

	fmt.Println("get data by request")
	var updateEvent model.Event

	err = ctx.ShouldBindJSON(&updateEvent)

	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("update process")
	updateEvent.ID = eventId

	err = event.UpdateEvents()
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not update data by id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Update Success"})
}

func deleteEvents(ctx *gin.Context) {
	var event model.Event
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not prase Id"})
		return
	}

	_, err = model.GetEventByID(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not prase get data by id"})
		return
	}

	err = event.DeleteEvent(eventId)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"Message": "Could not delete data by id"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Message": "Delete Success"})
}
