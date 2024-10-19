package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/model"
	"main.go/utils"
	"net/http"
)

func GetAllUser(ctx *gin.Context) {

	var users []model.User
	fmt.Println("dd")
	db.Find(&users)
	ctx.JSON(200, gin.H{
		"message": users,
	})
}

func AddUser(ctx *gin.Context) {

	var users model.User
	err := ctx.BindJSON(&users)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}

	users.Password = utils.HashPassword(users.Password)

	err = db.Create(&users).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant create passwor",
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created",
		"user": users,
	})
}

func UpdateUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user model.User
	err := ctx.BindJSON(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
	}
	user.Password = utils.HashPassword(user.Password)
	err = db.Model(&model.User{}).Where("id=?", id).Updates(user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant find user",
			"err":     err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Updated created",
		"user":    user,
	})

}

func DeletUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.Delete(&model.User{}, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant find user",
			"err":     err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "deleted User",
	})
}
