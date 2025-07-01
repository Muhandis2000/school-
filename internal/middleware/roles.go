// internal/middleware/roles.go
package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetString("role") != role {
			c.AbortWithStatusJSON(403, gin.H{"error": "доступ запрещен"})
			return
		}
		c.Next()
	}
}

// AdminOnly позволяет доступ только роли admin
func AdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ только для администратора"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// TeacherOnly позволяет доступ только роли teacher
func TeacherOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "teacher" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ только для учителя"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// StudentOnly позволяет доступ только роли student
func StudentOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != "student" {
			c.JSON(http.StatusForbidden, gin.H{"error": "Доступ только для студента"})
			c.Abort()
			return
		}
		c.Next()
	}
}
