package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"main.go/config"
	"main.go/model"
	"main.go/utils"
	"net/http"
	"time"
)

var db *gorm.DB
var validate *validator.Validate

func init() {
	db = config.DbInit()
	validate = validator.New()
}
func PassDb() *gorm.DB {
	return db
}
func Signup(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := validate.Struct(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.Password = utils.HashPassword(user.Password)
	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "email alredy used"})
		return
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err = db.Model(&model.User{}).Where("id=?", user.ID).Update("token", token).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update token"})
		return
	}
	if user.UserType == "admin" {
		c.Set("userType", "admin")
	} else if user.UserType == "user" {
		c.Set("userType", "user")
	}
	userType, exists := c.Get("userType")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User type not found"})
		return
	}
	c.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, gin.H{
		"token":    token,
		"userType": userType,
		"message":  "logine successfully",
	})
}

func Logine(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	var dbUser model.User
	err := db.Where("email=?", user.Email).First(&dbUser).Error
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "maile not fount pleas register"})
		return
	}
	pserr := utils.CheckPasswordHash(user.Password, dbUser.Password)
	if !pserr {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "password not match",
			"use": user,
			"db":  dbUser,
		})
		return
	}
	token, err := utils.GenerateToken(dbUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	err = db.Model(&model.User{}).Where("id=?", dbUser.ID).Update("token", token).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Could not update token"})
		return
	}

	if dbUser.UserType == "admin" {
		ctx.Set("userType", "admin")
	} else if dbUser.UserType == "user" {
		ctx.Set("userType", "user")
	}

	userType, exists := ctx.Get("userType")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User type not found"})
		return
	}

	ctx.SetCookie("token", token, int(24*time.Hour.Seconds()), "/", "localhost", false, true)
	ctx.JSON(200, gin.H{
		"token":    token,
		"userType": userType,
		"message":  "logine successfully",
	})
}

func LogOut(ctx *gin.Context) {
	ctx.SetCookie("token", "", -1, "/", "", false, true)
	ctx.JSON(200, gin.H{
		"message": "log out",
	})
}
