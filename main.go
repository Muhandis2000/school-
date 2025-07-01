// @title Online School API
// @version 1.0
// @description API для платформы
// @host localhost:8080
// @BasePath /
// @schemes http
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main

import (
	"fmt"

	"Final_project/config"

	"Final_project/internal/database"
	"Final_project/internal/logger"
	"Final_project/internal/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"                      // PostgreSQL driver
	swagFiles "github.com/swaggo/files"        // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func main() {
	// Загрузка конфигурации
	config.LoadConfig()

	// Инициализация логгера
	log := logger.Log
	log.Info("Starting application...")

	// Подключение к БД
	database.ConnectDB()
	log.Info("Connected to database")

	// Инициализация маршрутов
	router := gin.Default()
	router.Use(gin.Logger(), gin.Recovery())

	// Роуты
	routes.AuthRoutes(router)
	routes.AdminRoutes(router)
	routes.TeacherRoutes(router)
	routes.StudentRoutes(router)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swagFiles.Handler))
	// Запуск сервера

	port := config.Cfg.Server.Port
	log.Infof("Server running at :%d", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal(err)
	}
}
