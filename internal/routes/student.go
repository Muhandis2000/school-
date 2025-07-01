// internal/routes/student.go
package routes

import (
	"Final_project/internal/handlers"
	"Final_project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(r *gin.Engine) {
	student := r.Group("/student", middleware.AuthMiddleware(), middleware.StudentOnly())
	{
		student.GET("/schedule", handlers.GetSchedule)
		student.GET("/lessons", handlers.GetLessonsStudent)
		student.GET("/grades", handlers.GetGrades)
		student.POST("/homeworks", handlers.SubmitHomework)
		student.POST("/certificate", handlers.GenerateCertificate) // выдать сертификат
	}
}
