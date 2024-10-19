package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/model"
	"net/http"
	"strings"
)

func GetMedicine(ctx *gin.Context, db *gorm.DB) {
	var med string
	err := ctx.BindJSON(&med)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"messag": err,
		})
		return
	}

	res := strings.Split(med, ",")
	var tempMed []model.Medicine
	for _, v := range res {
		var medicin model.Medicine
		err := db.Where("med_name=?", v).Find(&medicin).Error
		if err != nil {
			ctx.JSON(200, gin.H{
				"message": "medicin not avalible",
				"meNAme":  v,
			})
		}
		tempMed = append(tempMed, medicin)
	}

	fmt.Println(res)
	totel := 0
	for _, p := range tempMed {
		totel += p.Price
	}
	ctx.JSON(200, gin.H{
		"medName": tempMed,
		"totoel":  totel,
	})
}
