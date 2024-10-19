package middleware

import (
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/model"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}

func AdminMidleWare() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		var user model.User
		token, err := ctx.Cookie("token")
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization required"})
			ctx.Abort()
			return
		}

		db := controllers.PassDb()
		err = db.Where("token=?", token).First(&user).Error
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Your Not Admin"})
			ctx.Abort()
			return
		}

		if user.UserType == "admin" {
			ctx.Set("userType", "admin")
			ctx.Next()
		} else if user.UserType == "user" {
			ctx.Set("userType", "user")
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Your Not The Admin Bro"})
			ctx.Abort()
		} else {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid user type"})
			ctx.Abort()
		}
	}
}
