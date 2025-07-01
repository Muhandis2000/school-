// internal/handlers/teacher.go
package handlers

import (
	"net/http"
	"time"

	"Final_project/internal/database"

	"Final_project/internal/models"

	"github.com/gin-gonic/gin"
)

// Отметка посещаемости
func MarkAttendance(c *gin.Context) {
	var att models.Attendance
	if err := c.ShouldBindJSON(&att); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	att.Date = time.Now()
	_, err := database.DB.Exec(
		"INSERT INTO attendance (student_id,lesson_id,date,present) VALUES ($1,$2,$3,$4)",
		att.StudentID, att.LessonID, att.Date, att.Present,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить посещаемость"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Посещаемость отмечена"})
}

// Выставление оценки
func AddGrade(c *gin.Context) {
	var grade models.Grade
	if err := c.ShouldBindJSON(&grade); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	grade.Date = time.Now()
	_, err := database.DB.Exec(
		"INSERT INTO grades (student_id,teacher_id,lesson_id,value,date) VALUES ($1,$2,$3,$4,$5)",
		grade.StudentID, grade.TeacherID, grade.LessonID, grade.Value, grade.Date,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить оценку"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Оценка выставлена"})
}
