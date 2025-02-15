package utils

import (
	"aitu-moment/logger"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(".env", "../.env", "../../.env"); err != nil {
		logger.GetLogger().Warnf("Warning: .env file not found or error loading it: %v", err)
	}
	ENC_SECRET = GetFromEnv("ENC_SCRET", "ComaBomaComaBoma")
}

var ENC_SECRET string

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

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func Encrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(ENC_SECRET))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, []byte(ENC_SECRET))
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return encode(cipherText), nil
}

func Decrypt(text string) (string, error) {
	block, err := aes.NewCipher([]byte(ENC_SECRET))
	if err != nil {
		return "", err
	}
	cipherText := decode(text)
	cfb := cipher.NewCFBDecrypter(block, []byte(ENC_SECRET))
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
