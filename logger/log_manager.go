package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	file, err := os.OpenFile("./logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func GetLogger() *logrus.Logger {
	return log
}
