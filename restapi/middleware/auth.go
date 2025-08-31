package middleware

import (
	"fmt"
	"net/http"
	"restapi/utils"

	"github.com/gin-gonic/gin"
)

func Authorization(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	fmt.Println(token)
	if token == "" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": "Not Autrorization"})
		return
	}

	userId, err := utils.VerifyToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"Message": err})
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}
