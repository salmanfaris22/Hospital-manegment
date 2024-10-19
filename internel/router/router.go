package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.go/controllers"
	"main.go/middleware"
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

	i.gin.POST("/signup", controllers.Signup)
	i.gin.POST("/logine", controllers.Logine)

	user := i.gin.Group("/user")
	user.Use(middleware.AuthMiddleware())
	{
		user.POST("/patient", controllers.PatientRegister)
		user.GET("/appoiment", controllers.GetAppointment)
		user.GET("/doctor", controllers.DoctorController)
		user.GET("/medicin", controllers.GetMedicine)
	}
	admin := i.gin.Group("/admin")
	admin.Use(middleware.AdminMidleWare())
	{
		admin.GET("/users", controllers.GetAllUser)
		admin.POST("/addUser", controllers.AddUser)
		admin.PUT("/update/:id", controllers.UpdateUser)
		admin.DELETE("/delete/:id", controllers.DeletUser)

	}
	i.gin.Run()
}

func NewRouter() Router {
	return &impel{
		gin: gin.New(),
	}
}
