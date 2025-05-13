package handlers

import (
	"aitu-moment/logger"
	"aitu-moment/models"
	"aitu-moment/services"
	"aitu-moment/utils"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	userService  *services.UserService
	eduService   *services.EduService
	groupService *services.GroupService
	mailService  *services.MailService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		userService:  services.NewUserService(),
		eduService:   services.NewEduService(),
		groupService: services.NewGroupService(),
		mailService:  services.NewMailService(),
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
		logger.GetLogger().Errorf("Erorr during getting register page %v", err.Error())
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "You are a filty little hoe",
		})
		return
	}

	c.HTML(http.StatusOK, "register.html", gin.H{
		"edu_list": programs,
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
		logger.GetLogger().Errorf("Erorr during getting register page %v", err.Error())
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "You are a filty little hoe",
		})
		return
	}

	if len(groups) > 0 {
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
			"error": "You are a filty little hoe",
		})
		return
	}

	user, err := h.userService.UpdateUser(user)
	if err != nil {
		logger.GetLogger().Errorf("Error during updating user %v", err.Error())
		c.HTML(http.StatusBadRequest, "userProfile.html", gin.H{
			"error": "You are a filty little hoe",
		})
		return
	}

	user, programs, group, err := h.userService.GetFullUserInfo(user.Id)

	if err != nil {
		logger.GetLogger().Errorf("Error during getting user info %v", err.Error())
		c.HTML(http.StatusBadRequest, "userProfile.html", gin.H{
			"error": "You are a filty little hoe",
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
			"error": "Wrong data format",
		})
		return
	}

	encryptedEmail, err := utils.Encrypt(user.Email)

	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Error during email encryption: " + err.Error(),
		})
		return
	}

	err = h.mailService.SendEmailVerification(user.Email, "http://"+c.Request.Host+"/verify?data="+encryptedEmail)

	if err != nil {
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Error during email verification: " + err.Error(),
		})
		return
	}

	user, err = h.userService.CreateUser(user)

	if err != nil {
		logger.GetLogger().Errorf("Error during registration %v", err.Error())
		c.HTML(http.StatusBadRequest, "register.html", gin.H{
			"error": "Error during registration",
		})
		return
	}

	c.HTML(http.StatusOK, "auth.html", gin.H{
		"fromLoginName": fmt.Sprintf("%s, please verify your account with the link sent on- %s", user.Name, user.Email),
	})
}

func (h *AuthHandler) Verify(c *gin.Context) {
	email := c.Query("data")

	if email == "" {
		c.HTML(http.StatusBadRequest, "auth.html", gin.H{
			"error": "Wrong verification link!",
		})
		return
	}

	decryptedEmail, err := utils.Decrypt(email)

	if err != nil {
		c.HTML(http.StatusBadRequest, "auth.html", gin.H{
			"error": "Error during email decryption",
		})
		return
	}

	err = h.userService.VerifyUser(decryptedEmail)

	if err != nil {
		c.HTML(http.StatusBadRequest, "auth.html", gin.H{
			"error": "Error during verification",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}

func (h *AuthHandler) Login(c *gin.Context) {

	email, _ := c.GetPostForm("email")
	passwd, _ := c.GetPostForm("passwd")

	user, err := h.userService.Authorize(email, passwd)

	if err != nil || user == nil {
		c.HTML(http.StatusUnauthorized, "auth.html", gin.H{
			"error": "Invalid credentials!",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
		"userID": user.Id,
	})

	tokenString, err := token.SignedString([]byte(utils.GetFromEnv("JWT_SECRET", "super_duper")))
	if err != nil {
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
