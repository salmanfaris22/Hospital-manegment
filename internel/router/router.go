package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/config"
	"main.go/controllers"
)

type Router interface {
	Start()
}
type impel struct {
	gin *gin.Engine
}

func (i impel) Start() {
	db := config.DbInit()
	i.gin.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * 3600,
		AllowCredentials: true,
	}))
	i.gin.POST("/patient", func(context *gin.Context) {
		controllers.PatientRegister(context, db)
	})
	i.gin.GET("/appoiment", func(context *gin.Context) {
		controllers.GetAppointment(context, db)
	})
	i.gin.Run()
}

func NewRouter() Router {
	return &impel{
		gin: gin.New(),
	}
}
