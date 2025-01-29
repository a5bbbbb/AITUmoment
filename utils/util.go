package utils

import (
	"aitu-moment/logger"
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
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

func GetUserFromClaims(c *gin.Context) (*int, error) {
	value, exists := c.Get("claims")
	if !exists {
		return nil, errors.New("Claims are inexistant")
	}

	claims, ok := value.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Error during claims parse")
	}

	userData, exists := claims["userID"]
	if !exists {
		return nil, errors.New("user claim not found")
	}

	userIDFloat, ok := userData.(float64)

	if !ok {
		return nil, errors.New("Error during parsing userID to float")
	}

	userID := int(userIDFloat)
	return &userID, nil

}
