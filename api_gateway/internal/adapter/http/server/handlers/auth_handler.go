package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/http/server/handlers/utils"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/models"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService  *services.UserService
	eduService   *services.EduService
	groupService *services.GroupService
	jwtSecret    string
}

func NewAuthHandler(
	userService *services.UserService,
	eduService *services.EduService,
	groupService *services.GroupService,
	jwtSecret string,
) *AuthHandler {
	return &AuthHandler{
		userService:  userService,
		eduService:   eduService,
		groupService: groupService,
		jwtSecret:    jwtSecret,
	}
}

func (h *AuthHandler) AuthPage(c *gin.Context) {

	userID, err := utils.GetUserFromClaims(c)

	if err != nil || userID == nil {
		logger.GetLogger().Errorf("Error during getting claims in profile page %v", err.Error())
		c.HTML(http.StatusOK, "auth.html", nil)
		return
	}

	c.Redirect(http.StatusFound, "/")

}

func (h *AuthHandler) RegisterPage(c *gin.Context) {
	programs, err := h.eduService.GetPrograms()
	if err != nil {
		logger.GetLogger().Errorf("Error during getting register page %v", err.Error())
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Something went very wrong...",
		})
		return
	}

	c.HTML(http.StatusOK, "register.html", gin.H{
		"edu_list": programs,
	})
}

func (h *AuthHandler) VerifyEmailPage(c *gin.Context) {
	token := c.Query("token")
	c.HTML(http.StatusOK, "emailVerificationIndex.html", gin.H{
		"token": token,
	})
}

func (h *AuthHandler) GroupsListPage(c *gin.Context) {
	eduProgStr := c.Query("educational_program")

	eduProg, err := strconv.Atoi(eduProgStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{"error": "Invalid id parameter"})
		return
	}

	groups, err := h.groupService.GetGroups(uint8(eduProg))
	if err != nil {
		logger.GetLogger().Errorf("Error during getting register page %v", err.Error())
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Something went very wrong...",
		})
		return
	}

	if len(*groups) > 0 {
		logger.GetLogger().Info("Got groups")
		c.HTML(http.StatusOK, "groupsList.html", gin.H{
			"group_list": groups,
		})
	}
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("auth_token", "", -1, "/", "", false, true)
	c.Set("claims", nil)
	c.Redirect(http.StatusFound, "/login")
}

func (h *AuthHandler) ProfilePage(c *gin.Context) {

	userID, err := utils.GetUserFromClaims(c)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting claims in profile page %v", err.Error())
		c.HTML(http.StatusBadRequest, "auth.html", nil)
		return
	}

	user, programs, group, err := h.userService.GetFullUserInfo(*userID)

	if err != nil {
		logger.GetLogger().Errorf("Error from userService in profile page %v", err.Error())
		c.HTML(http.StatusBadRequest, "auth.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.HTML(http.StatusOK, "userProfile.html", gin.H{
		"user":     user,
		"edu_list": programs,
		"group":    group,
	})
}

func (h *AuthHandler) UpdateUser(c *gin.Context) {

	var user *models.User
	if err := c.ShouldBind(&user); err != nil {
		logger.GetLogger().Errorf("Error during bind in update handler %v", err.Error())
		c.HTML(http.StatusBadRequest, "userProfile.html", gin.H{
			"error": "Something went very wrong...",
		})
		return
	}

	user, err := h.userService.UpdateUser(user)
	if err != nil {
		logger.GetLogger().Errorf("Error during updating user %v", err.Error())
		c.HTML(http.StatusBadRequest, "userProfile.html", gin.H{
			"error": "Something went very wrong...",
		})
		return
	}

	user, programs, group, err := h.userService.GetFullUserInfo(user.Id)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting user info %v", err.Error())
		c.HTML(http.StatusBadRequest, "userProfile.html", gin.H{
			"error": "Something went very wrong...",
		})
		return
	}

	c.HTML(http.StatusOK, "userProfile.html", gin.H{
		"user":     user,
		"edu_list": programs,
		"group":    group,
	})
}

func (h *AuthHandler) Register(c *gin.Context) {

	var user *models.User
	if err := c.ShouldBind(&user); err != nil {
		logger.GetLogger().Errorf("Error during bind in registration %v", err.Error())
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Something went very wrong...",
		})
		return
	}

	user, err := h.userService.CreateUser(user)

	if err != nil {
		logger.GetLogger().Errorf("Error during registration %v", err.Error())
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Something went very wrong...",
		})
		return
	}

	c.HTML(http.StatusOK, "auth.html", gin.H{
		"message":       "Check your email for email verification letter.",
		"fromLoginName": user.PublicName,
	})
}

func (h *AuthHandler) VerifyEmail(c *gin.Context) {
	token, ok := c.GetPostForm("token")
	if !ok {
		logger.GetLogger().Error("Error during email verification: 'token' field is absent from the request")
		c.HTML(http.StatusOK, "emailVerification.html", gin.H{
			"error": "We couldn't verify your account. Try to open link from your email again or apply for new verification link. If that does not work, then please reach out to support.",
		})
	}

	logger.GetLogger().Debug("Token: ", token)

	user, err := h.userService.VerifyEmail(token)
	if err != nil {
		logger.GetLogger().Errorln("Error during email verification token: ", token, " err: ", err.Error())
		c.HTML(http.StatusOK, "emailVerification.html", gin.H{
			"error": "We couldn't verify your account. Try to open link from your email again or apply for new verification link. If that does not work, then please reach out to support.",
		})
	}

	c.HTML(http.StatusOK, "emailVerification.html", gin.H{
		"PublicName": user.PublicName,
	})
}

func (h *AuthHandler) Login(c *gin.Context) {

	email, _ := c.GetPostForm("email")
	passwd, _ := c.GetPostForm("passwd")

	user, err := h.userService.Authorize(email, passwd)

	if err != nil || user == nil {
		logger.GetLogger().Errorf("Error during login %v", err.Error())
		c.HTML(http.StatusUnauthorized, "auth.html", gin.H{
			"error": "Invalid credentials!",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
		"userID": user.Id,
	})

	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		logger.GetLogger().Errorf("Error during login %v", err.Error())
		c.HTML(http.StatusBadRequest, "auth.html", gin.H{
			"error": err.Error(),
		})
		return
	}

	c.SetCookie("auth_token", tokenString, 3600*24, "/", "", false, true)

	// c.HTML(http.StatusOK, "home.html", gin.H{
	// 	"name": user.Name,
	// })

	c.Redirect(http.StatusFound, "/")

}
