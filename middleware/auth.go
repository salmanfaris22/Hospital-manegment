package middleware

import (
	"github.com/gin-gonic/gin"
	"main.go/utils"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			ctx.Abort()
			return
		}
		utils.ValidationToken(token)
		ctx.Next()
	}
}
