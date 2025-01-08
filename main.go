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
	userHandler := handlers.NewUserHandler(userRepo)
	e.GET("/", userHandler.GetHome)
	e.POST("/users", userHandler.SaveUser)

	e.Run(":8080")
}
