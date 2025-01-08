package main

import (
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"aitu-moment/db"
	"aitu-moment/db/repository"
	"aitu-moment/handlers"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func init() {

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	// log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	log.WithFields(logrus.Fields{
		"status": "OK!",
		"msg":    "HEY",
	}).Info("HEY2")

	databaseConnection, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer databaseConnection.Close()

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	userRepo := repository.NewUserRepository(databaseConnection.GetDB())
	userHandler := handlers.NewUserHandler(userRepo)
	e.GET("/", userHandler.GetHome)
	e.POST("/users", userHandler.SaveUser)
	e.GET("/mail", userHandler.GetMail)
	e.POST("/mail", userHandler.SendMail)

	e.Run(":8080")
}
