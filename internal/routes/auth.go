// internal/routes/auth.go
package routes

import (
	"Final_project/internal/handlers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.GET("/me", handlers.Me)
	}
}
