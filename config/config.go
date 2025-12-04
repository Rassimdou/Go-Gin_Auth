package config

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type DBconfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
}
type JWTconfig struct {
	Secret      string
	ExpiryHours time.Duration
	RefreshDays time.Duration
}
type SERVERconfig struct {
	Port string
}

func LoadConfig() (*DBconfig, *JWTconfig, *SERVERconfig, error) {
	//load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Note: .env file not found, using environment variables")
	}

	DB := &DBconfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("DB_USER", "rassim"),
		Password: getEnv("DB_PASSWORD", "rassim123"),
		Name:     getEnv("DB_NAME", "auth"),
	}
	JWT := &JWTconfig{
		Secret:      getEnv("JWT_SECRET", "CH_12andahwoadahdwa"),
		ExpiryHours: getEnvHours("JWT_EXPIRY_HOURS", 24*time.Hour),
		RefreshDays: getEnvDays("REFRESH_TOKEN_EXPIRY_DAYS", 7*24*time.Hour),
	}
	SERVER := &SERVERconfig{
		Port: getEnv("PORT", "8080"),
	}
	return DB, JWT, SERVER, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value

}

func getEnvHours(key string, defaultValue time.Duration) time.Duration {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer for %s: %s, using default", key, value)
		return defaultValue
	}
	return time.Duration(val) * time.Hour
}

func getEnvDays(key string, defaultValue time.Duration) time.Duration {
	value := getEnv(key, "")
	if value == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(value)
	if err != nil {
		log.Printf("Invalid integer for %s: %s, using default", key, value)
		return defaultValue
	}
	return time.Duration(val) * 24 * time.Hour
}
