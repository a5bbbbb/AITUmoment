package config

import (
	"os"
	"strconv"

	"email_service/internal/logger"

	"github.com/joho/godotenv"
)

var (
	RabbitMQURL    string
	SendEmailQueue string
	RoutingKey     string
	EmailExchange  string

	EmailSender   string
	EmailPasswd   string
	EmailHost     string
	EmailHostPort int
)

func init() {
	if err := godotenv.Load(); err != nil {
		logger.GetLogger().Warnf("Warning: .env file not found or error loading it: %v", err)
	}

	RabbitMQURL = getEnv("RABBITMQ_URL", "amqp://guest:guest@localhost")
	SendEmailQueue = getEnv("SendEmailQueue", "my_queue")
	RoutingKey = getEnv("SendEmailQueueRoutingKey", "send_email")
	EmailExchange = getEnv("EmailExchange", "email_exchange")
	EmailSender = getEnv("EMAIL", "serafbbs@gmail.com")
	EmailPasswd = getEnv("EMAIL_PASSWORD", "rkan gcwq klam zwiw")
	EmailHost = getEnv("EMAIL_HOST", "smtp.gmail.com")

	var err error
	EmailHostPort, err = strconv.Atoi(getEnv("EMAIL_HOST_PORT", "587"))

	if err != nil {
		EmailHostPort = 587
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}
