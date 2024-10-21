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

	var tempdate model.Date
	err = db.Where("date_time = ? AND slot = ? AND doctor_id = ?", appointment.Date, appointment.Slot, appointment.DoctorID).First(&tempdate).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "this date Booked",
			"err":     err,
		})
		return
	}
	if tempdate.Available == true {

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

	tempdate.Available = false
	tempdate.UserId = appointment.User.ID
	db.Create(&tempdate)
	err = db.Model(&tempdate).
		Where("date_time = ? AND doctor_id = ? AND slot = ?", appointment.Date, appointment.DoctorID, appointment.Slot).
		Update("available", false).Error

	if err != nil {

		fmt.Println("Error updating availability:", err)
		return
	}
	fmt.Println(appointment.Age)
	appointment.Doctor = doc
	ctx.JSON(200, gin.H{
		"message":      "your slot is Booked Be Redy",
		"time":         tempdate.Slot,
		"date":         tempdate.DateTime,
		"docName":      doc.DoctName,
		"dep":          doc.Dep,
		"patient_name": appointment.PatientName,
		"age":          appointment,
	})

}
