package main

import (
	"aitu-moment/db"
	"aitu-moment/handlers"
	"aitu-moment/logger"
	"fmt"

	"github.com/gin-gonic/gin"
)

var (
	authHandler *handlers.AuthHandler
)

func init() {

	authHandler = handlers.NewAuthHandler()
}

func main() {

	defer db.Close()

	if err := db.GetDB().DB.Ping(); err != nil {
		fmt.Println("Couldn't pind db.")
	}

	logger.GetLogger().Info("Setting up gin engine...")
	engine := gin.Default()

	engine.LoadHTMLGlob("./view/*")

	registerRoutes(engine)

	logger.GetLogger().Info("Successfully started the server")

	engine.Run(":8081")
}

func registerRoutes(e *gin.Engine) {
	e.GET("/login", authHandler.AuthPage)
	e.POST("/login", authHandler.Login)
	e.GET("/register", authHandler.RegisterPage)
	e.POST("/register", authHandler.Register)
	e.GET("/groupsList", authHandler.GroupsListPage)
	e.PUT("/user", authHandler.UpdateUser)
	e.GET("/user", authHandler.ProfilePage)
	e.POST("/logout", authHandler.Logout)
}
