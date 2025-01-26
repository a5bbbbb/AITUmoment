package main

import (
	"aitu-moment/db"
	"aitu-moment/handlers"
	"aitu-moment/logger"
	"aitu-moment/middleware"

	"github.com/gin-gonic/gin"
)

var (
	authHandler    *handlers.AuthHandler
	authMiddleware *middleware.Middleware
)

func init() {
	authHandler = handlers.NewAuthHandler()
	authMiddleware = middleware.NewMiddleware()
}

func main() {

	defer db.Close()

	logger.GetLogger().Info("Setting up gin engine...")
	engine := gin.Default()

	engine.LoadHTMLGlob("./view/*")

	registerRoutes(engine)

	logger.GetLogger().Info("Successfully started the server")

	engine.Run(":8080")
}

func registerRoutes(e *gin.Engine) {
	e.GET("/login", authHandler.AuthPage)
	e.POST("/login", authHandler.Login)
	e.GET("/register", authHandler.RegisterPage)
	e.POST("/register", authHandler.Register)
	e.GET("/groupsList", authHandler.GroupsListPage)
	e.Use(authMiddleware.AuthMiddleware)
	e.GET("/", authHandler.MainPage)
	e.PUT("/user", authHandler.UpdateUser)
	e.GET("/user", authHandler.ProfilePage)
	e.POST("/logout", authHandler.Logout)

}
