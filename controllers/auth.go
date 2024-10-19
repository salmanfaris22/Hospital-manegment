package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/config"
	"main.go/model"
	"main.go/utils"
	"net/http"
	"time"
)

var db *gorm.DB

// Initialize the database connection
func init() {
	db = config.DbInit()
}

func Signup(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash password
	user.Password = utils.HashPassword(user.Password)

	// Save user to DB
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User created"})
}

func Logine(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var dbUser model.User
	err := db.Where("email=?", user.Email).First(&dbUser).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pserr := utils.CheckPasswordHash(user.Password, dbUser.Password)
	if !pserr {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password not match",
			"use": user,
			"db":  dbUser,
		})
		return
	}
	token, err := utils.GenerateToken(dbUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = db.Model(&model.User{}).Where("id=?", dbUser.ID).Update("token", token).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update token"})
		return
	}
	ctx.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"token": token,
		"d":     dbUser,
	})
}
