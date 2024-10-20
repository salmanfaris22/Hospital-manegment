package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	var patient model.User

	token, err := ctx.Cookie("token")
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "token not font",
			"err":     err.Error(),
		})
		return
	}
	err = db.Where("token = ?", token).First(&patient).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "Appointment not found",
			"err":     err.Error(),
		})
		return
	}
	//appointment.PatientName=
	appointment.UserID = patient.ID
	appointment.User = patient

	var doc model.Doctor
	err = db.Where("doct_id=?", appointment.DoctorID).First(&doc).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "docter not found",
			"err":     err.Error(),
		})
		return
	}

	if appointment.Slot == "moring" {
		doc.TimeSlot1 = "booked"
	}
	if appointment.Slot == "evenig" {
		doc.TimeSlot2 = "booked"
	}

	err = db.Save(&doc).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "docter is bys",
			"err":     err.Error(),
		})
		return
	}
	appointment.Doctor = doc
	ctx.JSON(200, gin.H{
		"message": appointment,
	})

}
