// internal/routes/admin.go
package routes

import (
	"Final_project/internal/handlers"

	"Final_project/internal/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutes(r *gin.Engine) {
	admin := r.Group("/admin", middleware.AuthMiddleware(), middleware.AdminOnly())
	{
		// Верификация
		admin.GET("/verify/teachers", handlers.GetUnverifiedTeachers)
		admin.PATCH("/verify/teachers/:id", handlers.VerifyTeacher)
		admin.GET("/verify/students", handlers.GetUnverifiedStudents)
		admin.PATCH("/verify/students/:id", handlers.VerifyStudent)
		// Teachers CRUD
		admin.GET("/teachers", handlers.GetTeachers)
		admin.POST("/teachers", handlers.CreateTeacher)
		admin.PUT("/teachers/:id", handlers.UpdateTeacher)
		admin.DELETE("/teachers/:id", handlers.DeleteTeacher)
		// Students CRUD
		admin.GET("/students", handlers.GetStudents)
		admin.POST("/students", handlers.CreateStudent)
		admin.PUT("/students/:id", handlers.UpdateStudent)
		admin.DELETE("/students/:id", handlers.DeleteStudent)
		// Lessons CRUD
		admin.GET("/lessons", handlers.GetLessons)
		admin.POST("/lessons", handlers.CreateLesson)
		admin.PUT("/lessons/:id", handlers.UpdateLesson)
		admin.DELETE("/lessons/:id", handlers.DeleteLesson)
		// Schedules CRUD
		admin.GET("/schedules", handlers.GetSchedules)
		admin.POST("/schedules", handlers.CreateSchedule)
		admin.PUT("/schedules/:id", handlers.UpdateSchedule)
		admin.DELETE("/schedules/:id", handlers.DeleteSchedule)
		// Payments CRUD
		admin.GET("/payments", handlers.GetPayments)
		admin.POST("/payments", handlers.CreatePayment)
		admin.PUT("/payments/:id", handlers.UpdatePayment)
		admin.DELETE("/payments/:id", handlers.DeletePayment)
	}
}
