package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/controllers"
)

type Router interface {
	Start()
}
type impel struct {
	gin *gin.Engine
}

func (i impel) Start() {

	i.gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600,
		AllowCredentials: true,
	}))

	i.gin.POST("signup", controllers.Signup)
	i.gin.POST("/logine", controllers.Logine)

	i.gin.POST("/patient", controllers.PatientRegister)
	i.gin.GET("/appoiment", controllers.GetAppointment)
	i.gin.GET("/doctor", controllers.DoctorController)
	i.gin.GET("/medicin", controllers.GetMedicine)

	i.gin.Run()
}

func NewRouter() Router {
	return &impel{
		gin: gin.New(),
	}
}
