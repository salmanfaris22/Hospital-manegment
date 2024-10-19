package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"main.go/model"
	"net/http"
)

func PatientRegister(ctx *gin.Context) {
	var patient model.Patient

	err := ctx.BindJSON(&patient)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}
	res := db.Create(&patient)
	if res.Error != nil {
		fmt.Println("\033[31mError creating user:\033[0m", res.Error) // Red for errors
		ctx.JSON(200, gin.H{
			"message": "somthin went wrong",
			"err":     err,
		})
	} else {
		fmt.Println("\033[32mUser added to the database successfully!\033[0m") // Green for success
		ctx.JSON(200, gin.H{
			"sister":  "thanks for registering",
			"message": "pleas waite",
			"id":      patient.ID,
		})
	}
}
