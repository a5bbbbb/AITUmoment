package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/a5bbbbb/AITUmoment/api_gateway/config"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/http/server/handlers"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/http/server/middleware"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/services"

	"github.com/gin-gonic/gin"
)

const serviceIPAddress = "0.0.0.0:%d"

type API struct {
	server        *gin.Engine
	cfg           config.HTTPServer
	address       string
	authHandler   *handlers.AuthHandler
	threadHandler *handlers.ThreadsHandler
	middleware    *middleware.Middleware
}

func New(cfg config.Server, userService services.UserService, threadService services.ThreadService, groupService services.GroupService, eduProgramService services.EduService) *API {
	gin.SetMode(cfg.HTTPServer.Mode)

	server := gin.New()

	server.Use(gin.Recovery())

	server.LoadHTMLGlob("./view/*")

	authHandler := handlers.NewAuthHandler(&userService, &eduProgramService, &groupService, cfg.HTTPServer.JWTsecret)

	threadHandler := handlers.NewThreadsHandler(&userService, &eduProgramService, &groupService, &threadService)

	middleware := middleware.NewMiddleware(cfg.HTTPServer.JWTsecret)

	api := &API{
		server:        server,
		cfg:           cfg.HTTPServer,
		address:       fmt.Sprintf(serviceIPAddress, cfg.HTTPServer.Port),
		authHandler:   authHandler,
		threadHandler: threadHandler,
		middleware:    middleware,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	a.server.GET("/login", a.authHandler.AuthPage)
	a.server.POST("/login", a.authHandler.Login)
	a.server.GET("/register", a.authHandler.RegisterPage)
	a.server.POST("/register", a.authHandler.Register)
	a.server.GET("/verify", a.authHandler.VerifyEmailPage)
	a.server.POST("/verify", a.authHandler.VerifyEmail)
	a.server.GET("/groupsList", a.authHandler.GroupsListPage)
	a.server.Use(a.middleware.AuthMiddleware)
	a.server.GET("/", a.threadHandler.FeedPage)
	a.server.PUT("/user", a.authHandler.UpdateUser)
	a.server.GET("/user", a.authHandler.ProfilePage)
	a.server.POST("/logout", a.authHandler.Logout)
	a.server.GET("/thread/sub", a.threadHandler.GetSubThreads)
	a.server.GET("/thread/new", a.threadHandler.CreateThreadPage)
	a.server.POST("/thread/new", a.threadHandler.SaveThread)
	a.server.POST("/upvote", a.threadHandler.SaveUpvote)
}

func (a *API) Run(errCh chan<- error) {
	go func() {
		logger.GetLogger().Println("Http server starting on: ", a.address)

		if err := a.server.Run(a.address); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to start HTTP server: %v", err)
			return
		}
	}()
}

func (a *API) Stop(ctx context.Context) error {
	logger.GetLogger().Println("HTTP server shutting down gracefully")

	logger.GetLogger().Println("HTTP server stopped successfully")
	return nil
}
