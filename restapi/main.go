package main

import (
	"restapi/db"
	"restapi/route"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Init()
	server := gin.Default()
	route.RegisterRoutes(server)

	server.Run(":8080")
}
