// internal/handlers/student.go
package handlers

import (
	"net/http"
	"strconv"
	"time"

	"Final_project/internal/database"

	"Final_project/internal/models"

	"github.com/gin-gonic/gin"
)

func GetSchedule(c *gin.Context) {
	var schedules []models.Schedule
	// Здесь можно сузить по студенту, если будет логика
	err := database.DB.Select(&schedules, "SELECT * FROM schedules")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"schedule": schedules})
}

func GetLessonsStudent(c *gin.Context) {
	var lessons []models.Lesson
	err := database.DB.Select(&lessons, "SELECT * FROM lessons")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lessons": lessons})
}

func GetGrades(c *gin.Context) {
	userID := c.GetInt("userID")
	var grades []models.Grade
	err := database.DB.Select(&grades, "SELECT * FROM grades WHERE student_id=$1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"grades": grades})
}

// Сдача домашки
func SubmitHomework(c *gin.Context) {
	var hw models.Homework
	if err := c.ShouldBindJSON(&hw); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hw.SubmittedAt = time.Now()
	userID := c.GetInt("userID")
	hw.StudentID = userID
	_, err := database.DB.Exec(
		"INSERT INTO homeworks (student_id,lesson_id,content,submitted_at) VALUES ($1,$2,$3,$4)",
		hw.StudentID, hw.LessonID, hw.Content, hw.SubmittedAt,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось отправить домашку"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Домашнее задание отправлено"})
}
func EnrollCourse(c *gin.Context) {
	studentID := c.GetInt("userID")
	courseID, _ := strconv.Atoi(c.Param("id"))
	_, err := database.DB.Exec(
		"INSERT INTO student_course_progress (student_id, course_id) VALUES ($1,$2)",
		studentID, courseID,
	)
	if err != nil {
		c.JSON(400, gin.H{"error": "не удалось записаться"})
		return
	}
	c.JSON(200, gin.H{"message": "записано на курс"})
}

func UpdateProgress(studentID, courseID, delta int) error {
	// Увеличиваем progress на delta, но не более 100
	_, err := database.DB.Exec(
		"UPDATE student_course_progress SET progress = LEAST(progress + $1, 100) WHERE student_id=$2 AND course_id=$3",
		delta, studentID, courseID,
	)
	if err != nil {
		return err
	}
	// Если progress достиг 100, отмечаем completed=true
	_, err = database.DB.Exec(
		"UPDATE student_course_progress SET completed = TRUE WHERE student_id=$1 AND course_id=$2 AND progress >= 100",
		studentID, courseID,
	)
	return err
}
func GenerateCertificate(c *gin.Context) {
	var req struct{ CourseID int }
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": "некорректный запрос"})
		return
	}
	studentID := c.GetInt("userID")
	var progress models.StudentCourseProgress
	err := database.DB.Get(&progress,
		"SELECT * FROM student_course_progress WHERE student_id=$1 AND course_id=$2",
		studentID, req.CourseID)
	if err != nil || !progress.Completed {
		c.JSON(400, gin.H{"error": "курс не завершён или не найден"})
		return
	}
	// Здесь можно сгенерировать реальный сертификат (PDF, изображение и т.д.)
	c.JSON(200, gin.H{"message": "Сертификат успешно сформирован"})
}
