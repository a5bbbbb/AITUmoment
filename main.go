package main

import (
	"html/template"
	"os"

	"golang.org/x/time/rate"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"

	"aitu-moment/db"
	"aitu-moment/db/repository"
	"aitu-moment/handlers"
	"aitu-moment/middleware"
)

var log = logrus.New()

var limiter = rate.NewLimiter(1, 3) // Rate limit of 1 request per second with a burst of 3 requests

func init() {

	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func main() {
	log.SetLevel(logrus.DebugLevel)
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found or error loading it: %v", err)
	}

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

	// For file upload apparently https://github.com/gin-gonic/examples/blob/master/upload-file/multiple/main.go
	e.MaxMultipartMemory = 8 << 20 // 8 MiB

	e.Use(middleware.Ratelimited())
	e.SetFuncMap(template.FuncMap{
		"sub": func(a, b int) int { return a - b },
		"add": func(a, b int) int { return a + b },
	})
	e.LoadHTMLGlob("templates/*")

	userRepo := repository.NewUserRepository(databaseConnection.GetDB())
	userHandler := handlers.NewUserHandler(userRepo)
	threadHandler := handlers.NewThreadHandler(userRepo)
	e.GET("/", userHandler.GetHome)
	e.POST("/users", userHandler.SaveUser)
	e.GET("/mail", userHandler.GetMail)
	e.POST("/mail", userHandler.SendMail)
	e.GET("/threads", threadHandler.GetThreads)
	e.Run(":8080")
}
