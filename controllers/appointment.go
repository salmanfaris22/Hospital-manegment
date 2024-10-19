package controllers

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/model"
	"net/http"
)

func GetAppointment(ctx *gin.Context) {
	var appointment model.Appointment
	err := ctx.BindJSON(&appointment)
	if err != nil {
		fmt.Println("errr")
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	var patient model.Patient
	if err := db.First(&patient, appointment.PatientID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(200, gin.H{
				"message": "patio ns not exist",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "patio ns not exist",
		})
		return
	}

	var doc model.Doctor
	if err := db.First(&doc, appointment.DoctorID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(200, gin.H{
				"message": "docter not exister",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "docter not exister",
		})
		return
	}

	res := db.Create(&appointment)
	if res.Error != nil {
		fmt.Println("ersrr")
		fmt.Println("\033[31mError creating user:\033[0m", res.Error) // Red for errors
		ctx.JSON(200, gin.H{
			"message": "somthin went wrong",
			"err":     err,
		})
	} else {
		fmt.Println("\033[32mUser added to the database successfully!\033[0m") // Green for success
		ctx.JSON(200, gin.H{
			"appointment_token": appointment.TokenID,
			"docter_name":       doc.DoctID,
			"docte_id":          doc.DoctName,
			"departMent":        doc.Dep,
		})
	}
}
