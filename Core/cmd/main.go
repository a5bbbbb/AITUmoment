package main

import (
	"aitu-moment/db"
	"aitu-moment/handlers"
	"aitu-moment/logger"
	"aitu-moment/middleware"
	"aitu-moment/publisher"

	"github.com/gin-gonic/gin"
)

var (
	authHandler    *handlers.AuthHandler
	threadHandler  *handlers.ThreadsHandler
	authMiddleware *middleware.Middleware
)

func init() {
	authHandler = handlers.NewAuthHandler()
	threadHandler = handlers.NewThreadsHandler()
	authMiddleware = middleware.NewMiddleware()
}

func main() {

	defer db.Close()

	logger.GetLogger().Info("Setting up gin engine...")
	engine := gin.Default()

	engine.LoadHTMLGlob("./view/*")

	registerRoutes(engine)

	publisher.InitProvider()
	defer publisher.CloseConn()
	defer publisher.GetProvider().EmailPublisher.ClosePublisher()

	logger.GetLogger().Info("Successfully started the server")

	engine.Run(":8080")

}

func registerRoutes(e *gin.Engine) {
	e.GET("/login", authHandler.AuthPage)
	e.POST("/login", authHandler.Login)
	e.GET("/register", authHandler.RegisterPage)
	e.POST("/register", authHandler.Register)
	e.GET("/groupsList", authHandler.GroupsListPage)
	e.GET("/verify", authHandler.Verify)
	e.Use(authMiddleware.AuthMiddleware)
	e.GET("/", threadHandler.FeedPage)
	e.PUT("/user", authHandler.UpdateUser)
	e.GET("/user", authHandler.ProfilePage)
	e.POST("/logout", authHandler.Logout)
	e.GET("/thread/sub", threadHandler.GetSubThreads)
	e.GET("/thread/new", threadHandler.CreateThreadPage)
	e.POST("/thread/new", threadHandler.SaveThread)
	e.POST("/upvote", threadHandler.SaveUpvote)

}
