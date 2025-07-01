// internal/handlers/auth.go
package handlers

import (
	"Final_project/internal/logger"
	"net/http"

	"Final_project/internal/database"
	"Final_project/internal/models"

	"Final_project/internal/utils"

	"github.com/gin-gonic/gin"
)

type RegisterInput struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Role     string `json:"role" binding:"required"` // "admin","teacher","student"
}

// @Summary Регистрация пользователя
// @Description Создает нового пользователя с ролью (student/teacher/admin).
// @Tags auth
// @Accept json
// @Produce json
// @Param registerRequest body RegisterRequest true "Данные для регистрации"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
// Register создаёт нового пользователя
func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Хешируем пароль
	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось хешировать пароль"})
		return
	}

	user := models.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: hashedPassword,
		Role:     input.Role,
		Verified: false,
	}

	// Вставка пользователя в БД
	err = database.DB.QueryRow(
		"INSERT INTO users (name, email, password, role, verified) VALUES ($1,$2,$3,$4,$5) RETURNING id",
		user.Name, user.Email, user.Password, user.Role, user.Verified,
	).Scan(&user.ID)
	if err != nil {
		logger.Log.Error(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось зарегистрироваться"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Пользователь зарегистрирован", "id": user.ID})
}

type LoginInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login аутентифицирует пользователя и возвращает JWT
func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	// Ищем пользователя по email
	err := database.DB.Get(&user, "SELECT * FROM users WHERE email=$1", input.Email)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}
	// Проверяем пароль
	if !utils.CheckPassword(user.Password, input.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный email или пароль"})
		return
	}
	if !user.Verified {
		c.JSON(http.StatusForbidden, gin.H{"error": "Пользователь не подтвержден администратором"})
		return
	}
	// Генерируем JWT
	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать токен"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

// Me возвращает данные текущего пользователя
func Me(c *gin.Context) {
	userID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "user_id не найден в контексте"})
		return
	}

	var user models.User
	err := database.DB.Get(&user, "SELECT id, name, email, role FROM users WHERE id=$1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить профиль"})
		return
	}

	c.JSON(http.StatusOK, user)
}
