package utils

import (
	"aitu-moment/logger"
	"errors"
	"os"
	"sync"
	"time"

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

type SafeMapStringTime struct {
	mu sync.Mutex
	v  map[string]time.Time
}

func NewSafeMapStringTime() *SafeMapStringTime {
	return &SafeMapStringTime{v: make(map[string]time.Time)}
}

func (c *SafeMapStringTime) Set(key string, deadline time.Time) {
	c.mu.Lock()
	c.v[key] = deadline
	c.mu.Unlock()
}

func (c *SafeMapStringTime) Unset(key string) {
	c.mu.Lock()
	delete(c.v, key)
	c.mu.Unlock()
}

func (c *SafeMapStringTime) Value(key string) (time.Time, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.v[key]
	return value, ok
}
