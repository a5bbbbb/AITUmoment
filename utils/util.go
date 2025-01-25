package utils

import (
	"aitu-moment/logger"
	"os"

	"github.com/joho/godotenv"
)



func init(){
    if err := godotenv.Load(); err != nil {
        logger.GetLogger().Warnf("Warning: .env file not found or error loading it: %v", err)
    }
}

func GetFromEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
