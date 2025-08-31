package route

import (
	"restapi/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", middleware.Authorization, getEvent)
	// middleware 2 can set process
	authen := server.Group("/")
	authen.Use(middleware.Authorization)
	authen.POST("/events", createEvent)
	authen.PUT("/events/:id", updateEvents)
	authen.DELETE("/events/:id", deleteEvents)
	authen.POST("/events/:id/register", registerForEvents)
	authen.DELETE("/events/:id/register", cancelRegistration)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
