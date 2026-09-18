package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort        string
	ServerHost        string
	DBHost            string
	DBPort            string
	DBUser            string
	DBPassword        string
	DBName            string
	JWTSecret         string
	JWTLifeSpan       string
	PrinterType       string
	PrinterPaperWidth string
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		ServerPort:        getEnv("SERVER_PORT", "8080"),
		ServerHost:        getEnv("SERVER_HOST", "localhost"),
		DBHost:            getEnv("DB_HOST", "127.0.0.1"),
		DBPort:            getEnv("DB_PORT", "3306"),
		DBUser:            getEnv("DB_USER", "root"),
		DBPassword:        getEnv("DB_PASSWORD", ""),
		DBName:            getEnv("DB_NAME", "mypos_db"),
		JWTSecret:         getEnv("JWT_SECRET", "super_secret_local_jwt_key_mypos_2026"),
		JWTLifeSpan:       getEnv("JWT_LIFESPAN", "24"),
		PrinterType:       getEnv("PRINTER_TYPE", "usb"),
		PrinterPaperWidth: getEnv("PRINTER_PAPER_WIDTH", "58mm"),
	}
}
