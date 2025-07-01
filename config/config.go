// config/config.go
package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	DB struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
		SSLMode  string `yaml:"sslmode"`
	} `yaml:"db"`
	JWT struct {
		Secret string `yaml:"secret"`
	} `yaml:"jwt"`
}

var Cfg Config

// LoadConfig читает config/config.yaml и .env для настройки приложения
func LoadConfig() {
	file, err := os.Open("config/config.yaml")
	if err != nil {
		log.Fatalf("Не удалось открыть конфиг: %v", err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&Cfg); err != nil {
		log.Fatalf("Не удалось декодировать конфиг: %v", err)
	}

	// Перезаписываем из переменных окружения, если заданы
	if v := os.Getenv("DB_HOST"); v != "" {
		Cfg.DB.Host = v
	}
	if v := os.Getenv("DB_PORT"); v != "" {
		Cfg.DB.Port = v
	}
	if v := os.Getenv("DB_USER"); v != "" {
		Cfg.DB.User = v
	}
	if v := os.Getenv("DB_PASSWORD"); v != "" {
		Cfg.DB.Password = v
	}
	if v := os.Getenv("DB_NAME"); v != "" {
		Cfg.DB.Name = v
	}
	if v := os.Getenv("JWT_SECRET"); v != "" {
		Cfg.JWT.Secret = v
	}
}
