// internal/middleware/auth.go
package middleware

import (
	"net/http"
	"strings"

	"Final_project/config"
	"Final_project/internal/utils"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// AuthMiddleware проверяет JWT и кладёт в контекст userID и role
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Отсутствует токен авторизации"})
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат токена"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		token, err := jwt.ParseWithClaims(tokenString, &utils.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.Cfg.JWT.Secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
			c.Abort()
			return
		}

		claims := token.Claims.(*utils.JWTClaims)
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}
