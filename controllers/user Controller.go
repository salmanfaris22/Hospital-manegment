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
	fmt.Println(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "gtting eroror",
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
		"message": "User Updated",
		"user":    user,
	})

}

func DeletUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user model.User
	fmt.Println(id)
	err := db.Where("id=?", id).First(&user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant find user",
			"err":     err,
		})
	}
	if user.UserType == "admin" {
		ctx.JSON(200, gin.H{
			"message": "user cant delet Beacase He is Admin",
		})
		return
	}
	err = db.Delete(&model.User{}, id).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": "cant find user",
			"err":     err,
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": " User deleted",
	})
}
