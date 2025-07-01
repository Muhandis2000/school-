// internal/utils/jwt.go
package utils

import (
	"time"

	"Final_project/config"

	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	UserID int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// GenerateJWT создаёт токен для пользователя
func GenerateJWT(userID int, role string) (string, error) {
	secret := []byte(config.Cfg.JWT.Secret)
	claims := &JWTClaims{
		UserID: userID,
		Role:   role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
