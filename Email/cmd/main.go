package main

import (
	"email_service/internal/consumer"
	"email_service/internal/logger"

	"os"
	"os/signal"
	"syscall"
)

func main() {
	logger.GetLogger().Info("Email service: setting up the service....")

	defer consumer.CloseConn()

	emailConsumer, err := consumer.NewEmailConsumer()
	if err != nil {
		logger.GetLogger().Error("Email service: failed to initialize email consumer.")
		return
	}
	defer emailConsumer.CloseConsumer()

	go emailConsumer.RunConsumer()

	logger.GetLogger().Info("Email service: Ready and serving")

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	<-sigs
	logger.GetLogger().Info("Email service: shutting down gracefully.")

}
