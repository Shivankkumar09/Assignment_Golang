package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/Shivankkumar09/Assignment_Golang/controllers"
	"github.com/Shivankkumar09/Assignment_Golang/middlewares"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)

	receptionist := r.Group("/receptionist")
	receptionist.Use(middlewares.AuthMiddleware("receptionist"))
	{
		receptionist.POST("/patients/addpatient", controllers.CreatePatient)
		receptionist.GET("/patients", controllers.GetPatients)
		receptionist.PUT("/patients/:id", controllers.UpdatePatient)
		receptionist.DELETE("/patients/:id", controllers.DeletePatient)
	}

	doctor := r.Group("/doctor")
    doctor.Use(middlewares.AuthMiddleware("doctor"))
   {
	doctor.GET("/patients", controllers.DoctorGetPatients)
	doctor.PUT("/patients/:id", controllers.DoctorUpdatePatient)
   }

}
