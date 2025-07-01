// internal/routes/teacher.go
package routes

import (
	"Final_project/internal/handlers"

	"Final_project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func TeacherRoutes(r *gin.Engine) {
	teacher := r.Group("/teacher", middleware.AuthMiddleware(), middleware.TeacherOnly())
	{
		teacher.POST("/attendance", handlers.MarkAttendance)
		teacher.POST("/grades", handlers.AddGrade)
	}
}
