package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	mode, ok := os.LookupEnv("APP_MODE")

	if !ok {
		panic("APP_MODE env is not set")
	}

	if mode == "DEV" {
		initDevelopmentLogger()
	} else if mode == "PROD" {
		initProductionLogger()
	}
}

func initDevelopmentLogger() {
	log = logrus.New()
	log.Warn("Development mode logger")
	log.SetLevel(logrus.DebugLevel)
}

func initProductionLogger() {
	log = logrus.New()
	file, err := os.OpenFile("./logs/aitu_mom.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err == nil {
		log.Out = file
	} else {
		log.Panic("Failed to log to file in production mode")
	}

	log.Info("Production mode logger")
}

func GetLogger() *logrus.Logger {
	return log
}
