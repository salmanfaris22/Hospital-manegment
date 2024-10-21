package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/helpers"
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

	err, patient := helpers.UserFindHelp(ctx, db)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "pleas logine",
			"err":     err,
		})
		return
	}

	var tempdate model.Date
	err = db.Where("date_time = ? AND slot = ? AND doctor_id = ?", appointment.Date, appointment.Slot, appointment.DoctorID).First(&tempdate).Error
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "this cant find slot Booked",
			"err":     err,
		})
		return
	}
	if tempdate.Available != true {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "this slot Booked",
			"err":     err,
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
	tempdate.UserId = appointment.UserID
	db.Create(&tempdate)
	err = db.Model(&tempdate).
		Where("date_time = ? AND doctor_id = ? AND slot = ?", appointment.Date, appointment.DoctorID, appointment.Slot).
		Updates(map[string]interface{}{
			"available": false,
			"user_id":   appointment.User.ID,
		}).Error

	if err != nil {
		fmt.Println("Error updating availability:", err)
		return
	}
	fmt.Println(appointment.Age)
	appointment.Doctor = doc
	appointment.Doctor.TimeSlot2 = appointment.Slot
	appointment.DateID = tempdate.ID
	db.Create(&appointment)

	ctx.JSON(200, gin.H{
		"message":      "your slot is Booked Be Redy",
		"time":         tempdate.Slot,
		"date":         tempdate.DateTime,
		"docName":      doc.DoctName,
		"dep":          doc.Dep,
		"patient_name": appointment.PatientName,
		"age":          appointment.Age,
		"token_id":     appointment.TokenID,
	})

}
func GetAllApoiment(ctx *gin.Context) {
	var results []model.AppointmentWithDoctor
	err, user := helpers.UserFindHelp(ctx, db)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "pleas logine",
			"err":     err,
		})
		return
	}
	err = db.Table("appointments").
		Select("appointments.*,doctors.doct_name as doctor_name").
		Joins("JOIN doctors ON doctors.doct_id = appointments.doctor_id").
		Where("appointments.user_id = ?", user.ID).
		Scan(&results).Error
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": results,
	})
}

func CancellApoiment(ctx *gin.Context) {
	var results model.AppointmentWithDoctor
	err := ctx.BindJSON(&results)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "pleas slelect apoiment",
			"err":     err,
		})
		return
	}
	err = db.Model(&model.Date{}).
		Where("id=?", results.DateID).
		Update("available", true).Error

	if err != nil {
		fmt.Println("Error updating availability:", err)
		return
	}
	err = db.Where("date_id = ?", results.DateID).Delete(&model.Appointment{}).Error
	fmt.Println(results.DateID)

	if err != nil {
		fmt.Println("Error deleting appointment:", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": " Apoimnet  deleted",
	})
}
