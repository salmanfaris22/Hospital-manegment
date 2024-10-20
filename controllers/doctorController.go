package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"main.go/model"
)

type GetToken struct {
	ID string `json:"id"`
}

func DoctorController(ctx *gin.Context) {
	var tokenID GetToken
	err := ctx.BindJSON(&tokenID)
	if err != nil {
		ctx.JSON(200, gin.H{
			"message": "error tokenID",
		})
		return
	}
	var token model.Appointment
	if err := db.First(&token, tokenID.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(200, gin.H{
				"message": "token not exister pleas get token",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "token not exister pleas get token",
		})
		return
	}

	var patient model.User
	if err := db.First(&patient, token.UserID).Error; err != nil {
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
	if err := db.First(&doc, token.DoctorID).Error; err != nil {
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

	var medicine model.Medicine
	if err := db.First(&medicine, 2).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.JSON(200, gin.H{
				"message": "medicn not availible not exister",
			})
			return
		}
		ctx.JSON(200, gin.H{
			"message": "medicn not availible not exister",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"TokenID":      token.DoctorID,
		"hellow":       doc,
		"doc_name":     doc.DoctName,
		"pationt_name": patient.FirstName,
		"medicin":      "Aspirin,Ibuprofen",
	})
}
