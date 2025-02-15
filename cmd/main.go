package main

import (
	"aitu-moment/db"
	"aitu-moment/handlers"
	"aitu-moment/logger"
	"aitu-moment/middleware"

	"github.com/gin-gonic/gin"
)

var (
	redirectHandler *handlers.RedirectHandler
	threadHandler   *handlers.ThreadsHandler
	authMiddleware  *middleware.Middleware
)

func init() {
	redirectHandler = handlers.NewRedirectHandler()
	threadHandler = handlers.NewThreadsHandler()
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
	e.GET("/login", redirectHandler.Redirect("8081", "/login"))
	e.POST("/login", redirectHandler.Redirect("8081", "/login"))
	e.GET("/register", redirectHandler.Redirect("8081", "/register"))
	e.POST("/register", redirectHandler.Redirect("8081", "/register"))
	e.GET("/groupsList", redirectHandler.Redirect("8081", "/groupsList"))
	e.PUT("/user", redirectHandler.Redirect("8081", "/user"))
	e.GET("/user", redirectHandler.Redirect("8081", "/user"))
	e.POST("/logout", redirectHandler.Redirect("8081", "/logout"))
	e.Use(authMiddleware.AuthMiddleware)
	e.GET("/", threadHandler.FeedPage)
	e.GET("/thread/sub", threadHandler.GetSubThreads)
	e.GET("/thread/new", threadHandler.CreateThreadPage)
	e.POST("/thread/new", threadHandler.SaveThread)
	e.POST("/upvote", threadHandler.SaveUpvote)

}
