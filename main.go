package main

import (
	"html/template"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"aitu-moment/db"
	"aitu-moment/db/repository"
	"aitu-moment/handlers"
)

var log = logrus.New()

func init() {

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
	}).Info("STARTED!!")

	databaseConnection, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer databaseConnection.Close()

	e := gin.Default()
	e.SetFuncMap(template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"add": func(a, b int) int { return a + b },
	})
	e.LoadHTMLGlob("templates/*")

	userRepo := repository.NewUserRepository(databaseConnection.GetDB())
	userHandler := handlers.NewUserHandler(userRepo)
	e.GET("/", userHandler.GetHome)
	e.POST("/users", userHandler.SaveUser)
	e.GET("/mail", userHandler.GetMail)
	e.POST("/mail", userHandler.SendMail)
	e.GET("/threads", userHandler.GetThreads)

	e.Run(":8080")
}
