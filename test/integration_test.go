package tests

import (
	"aitu-moment/db"
	"aitu-moment/handlers"
	"aitu-moment/middleware"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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

func TestPingRoute(t *testing.T) {
	defer db.Close()
	router := gin.Default()

	router.LoadHTMLGlob("../view/*")

	registerRoutes(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/login", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}

// Test for POST /user/add
func TestPostUser(t *testing.T) {
	defer db.Close()
	router := gin.Default()

	router.LoadHTMLGlob("../view/*")

	registerRoutes(router)

	w := httptest.NewRecorder()

	type Data struct {
		Email  string `json:"email"`
		Passwd string `json:"passwd"`
	}

	a := Data{
		Email:  "a%40a.a",
		Passwd: "a",
	}

	formData := "email=" + a.Email + "&passwd=" + a.Passwd
	req, _ := http.NewRequest("POST", "/login", strings.NewReader(formData))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)

	assert.Equal(t, 302, w.Code)
}
