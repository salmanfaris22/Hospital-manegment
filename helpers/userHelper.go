package helpers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/model"
	"net/http"
)

func UserFindHelp(ctx *gin.Context, db *gorm.DB) (error, model.User) {
	var patient model.User
	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "token not font",
			"err":     err.Error(),
		})
		return err, patient
	}

	err = db.Where("token = ?", token).First(&patient).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Pleas Logine",
			"err":     err.Error(),
		})
		return err, patient
	}
	return nil, patient
}
