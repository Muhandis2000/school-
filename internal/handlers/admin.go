// internal/handlers/admin.go
package handlers

import (
	"net/http"
	"time"

	"Final_project/internal/database"
	"Final_project/internal/models"
	"Final_project/internal/utils"

	"github.com/gin-gonic/gin"
)

// ----- Верификация учителей и студентов -----
// GetUnverifiedTeachers возвращает список неподтверждённых учителей
func GetUnverifiedTeachers(c *gin.Context) {
	var teachers []models.User
	err := database.DB.Select(&teachers, "SELECT * FROM users WHERE role='teacher' AND verified=false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"unverified_teachers": teachers})
}

// VerifyTeacher устанавливает verified=true для учителя по ID
func VerifyTeacher(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("UPDATE users SET verified=true WHERE id=$1 AND role='teacher'", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Учитель подтвержден"})
}

// GetUnverifiedStudents и VerifyStudent аналогично
func GetUnverifiedStudents(c *gin.Context) {
	var students []models.User
	err := database.DB.Select(&students, "SELECT * FROM users WHERE role='student' AND verified=false")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка БД"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"unverified_students": students})
}

func VerifyStudent(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("UPDATE users SET verified=true WHERE id=$1 AND role='student'", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Студент подтверждён"})
}

// ----- CRUD для учителей -----
func GetTeachers(c *gin.Context) {
	var teachers []models.User
	err := database.DB.Select(&teachers, "SELECT * FROM users WHERE role='teacher'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teachers": teachers})
}

func CreateTeacher(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Role = "teacher"
	// Пароль должен быть захеширован
	hashed, _ := utils.HashPassword(input.Password)
	input.Password = hashed
	input.Verified = true // сразу подтверждаем при создании админом
	err := database.DB.QueryRow(
		"INSERT INTO users (name,email,password,role,verified) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		input.Name, input.Email, input.Password, input.Role, input.Verified,
	).Scan(&input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать учителя"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teacher": input})
}

func UpdateTeacher(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3 AND role='teacher'",
		input.Name, input.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Учитель обновлён"})
}

func DeleteTeacher(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM users WHERE id=$1 AND role='teacher'", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Учитель удалён"})
}

// ----- CRUD для студентов -----
func GetStudents(c *gin.Context) {
	var students []models.User
	err := database.DB.Select(&students, "SELECT * FROM users WHERE role='student'")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"students": students})
}

func CreateStudent(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	input.Role = "student"
	hashed, _ := utils.HashPassword(input.Password)
	input.Password = hashed
	input.Verified = true
	err := database.DB.QueryRow(
		"INSERT INTO users (name,email,password,role,verified) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		input.Name, input.Email, input.Password, input.Role, input.Verified,
	).Scan(&input.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать студента"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"student": input})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("UPDATE users SET name=$1, email=$2 WHERE id=$3 AND role='student'",
		input.Name, input.Email, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Студент обновлён"})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM users WHERE id=$1 AND role='student'", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Студент удалён"})
}

// ----- CRUD для уроков -----
func GetLessons(c *gin.Context) {
	var lessons []models.Lesson
	err := database.DB.Select(&lessons, "SELECT * FROM lessons")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lessons": lessons})
}

func CreateLesson(c *gin.Context) {
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.DB.QueryRow(
		"INSERT INTO lessons (title,description) VALUES ($1,$2) RETURNING id",
		lesson.Title, lesson.Description,
	).Scan(&lesson.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать урок"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"lesson": lesson})
}

func UpdateLesson(c *gin.Context) {
	id := c.Param("id")
	var lesson models.Lesson
	if err := c.ShouldBindJSON(&lesson); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("UPDATE lessons SET title=$1, description=$2 WHERE id=$3",
		lesson.Title, lesson.Description, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Урок обновлён"})
}

func DeleteLesson(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM lessons WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Урок удалён"})
}

// ----- CRUD для расписания -----
func GetSchedules(c *gin.Context) {
	var schedules []models.Schedule
	err := database.DB.Select(&schedules, "SELECT * FROM schedules")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"schedules": schedules})
}

func CreateSchedule(c *gin.Context) {
	var sch models.Schedule
	if err := c.ShouldBindJSON(&sch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.DB.QueryRow(
		"INSERT INTO schedules (lesson_id,teacher_id,start_time,end_time) VALUES ($1,$2,$3,$4) RETURNING id",
		sch.LessonID, sch.TeacherID, sch.StartTime, sch.EndTime,
	).Scan(&sch.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать расписание"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"schedule": sch})
}

func UpdateSchedule(c *gin.Context) {
	id := c.Param("id")
	var sch models.Schedule
	if err := c.ShouldBindJSON(&sch); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("UPDATE schedules SET lesson_id=$1, teacher_id=$2, start_time=$3, end_time=$4 WHERE id=$5",
		sch.LessonID, sch.TeacherID, sch.StartTime, sch.EndTime, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Расписание обновлено"})
}

func DeleteSchedule(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM schedules WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Расписание удалено"})
}

// ----- CRUD для платежей -----
func GetPayments(c *gin.Context) {
	var payments []models.Payment
	err := database.DB.Select(&payments, "SELECT * FROM payments")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка получения"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payments": payments})
}

func CreatePayment(c *gin.Context) {
	var p models.Payment
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := database.DB.QueryRow(
		"INSERT INTO payments (student_id,tariff_id,amount,date) VALUES ($1,$2,$3,$4) RETURNING id",
		p.StudentID, p.TariffID, p.Amount, time.Now(),
	).Scan(&p.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать платёж"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"payment": p})
}

func UpdatePayment(c *gin.Context) {
	id := c.Param("id")
	var p models.Payment
	if err := c.ShouldBindJSON(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec("UPDATE payments SET student_id=$1, tariff_id=$2, amount=$3 WHERE id=$4",
		p.StudentID, p.TariffID, p.Amount, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось обновить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Платёж обновлён"})
}

func DeletePayment(c *gin.Context) {
	id := c.Param("id")
	_, err := database.DB.Exec("DELETE FROM payments WHERE id=$1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось удалить"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Платёж удалён"})
}
