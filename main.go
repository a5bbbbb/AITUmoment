package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"aitu-moment/db"
	"aitu-moment/db/repository"
	"aitu-moment/handlers"
)

func main() {

	databaseConnection, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err.Error())
	}

	defer databaseConnection.Close()

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	userRepo := repository.NewUserRepository(databaseConnection.GetDB())
	threadRepo := repository.NewThreadRepository(databaseConnection.GetDB())
	userHandler := handlers.NewUserHandler(userRepo)
	threadHandler := handlers.NewThreadHandler(threadRepo)

	e.GET("/", userHandler.GetHome)
	e.POST("/users", userHandler.SaveUser)
	e.GET("/threads", threadHandler.GetThreads)
	e.POST("/threads", threadHandler.CreateThread)

	e.Run(":8080")
}
