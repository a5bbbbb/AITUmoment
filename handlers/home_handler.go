package handlers

import (
	"aitu-moment/db/repository"
	"aitu-moment/models"
	"net/http"
	"net/smtp"
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
	//"github.com/wneessen/go-mail"

	"github.com/gin-gonic/gin"
)

// Create a new instance of the logger. You can have any number of instances.
var log = logrus.New()

func init() {

	// The API for setting attributes is a little different than the package level
	// exported logger. See Godoc.
	// log.Out = os.Stdout

	// You could set this to any `io.Writer` such as a file
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetHome(c *gin.Context) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		getErrorResponse(c, err.Error())
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"name":  "Awesome",
		"users": users,
	})

}

func (h *UserHandler) SaveUser(c *gin.Context) {
	edu_prog_int := c.PostForm("educational_program")
	educational_program_int, _ := strconv.Atoi(edu_prog_int)

	user := models.User{
		Name:               c.PostForm("username"),
		EducationalProgram: uint8(educational_program_int),
		Program_name:       "",
	}
	id, err := h.repo.CreateUser(user)
	if err != nil {
		getErrorResponse(c, err.Error())
		return
	}

	savedUser, err := h.repo.GetUser(id)
	if err != nil {
		getErrorResponse(c, err.Error())
		return
	}

	c.HTML(http.StatusOK, "user.html", gin.H{
		"Name":         savedUser.Name,
		"Program_name": savedUser.Program_name,
	})

}

func (h *UserHandler) SendMail(c *gin.Context) {

	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	from := "kantaidaulet@gmail.com"
	pass := "xfum mvti aixw atsn"

	email := "seraf5@bk.ru"
	log.Info(email)
	msg := c.PostForm("message")

	auth := smtp.PlainAuth("", from, pass, smtpHost)

	logrus.WithFields(logrus.Fields{
		"from":    from,
		"to":      email,
		"subject": msg,
	}).Info("Sending email started")

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{email}, []byte(msg))
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"from":    from,
			"to":      email,
			"subject": msg,
			"error":   err.Error(),
		}).Error("Failed to send email")
		getErrorResponse(c, err.Error())
		return
	}

	log.WithFields(logrus.Fields{
		"from":    from,
		"to":      email,
		"subject": msg,
	}).Info("Email sent successfully")

	c.HTML(http.StatusOK, "email.html", gin.H{})
}

// func (h *UserHandler) SendMail(c *gin.Context) {
//
// 	email := c.PostForm("mail")
// 	log.Info(email)
// 	msg := c.PostForm("message")
//
// 	// First we create a mail message
// 	m := mail.NewMsg()
// 	if err := m.From(email); err != nil {
// 		log.Errorf("failed to set From address: %s", err)
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
// 	if err := m.To("serafbbs@gmail.com"); err != nil {
// 		log.Errorf("failed to set To address: %s", err)
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
// 	m.Subject(msg)
// 	m.SetBodyString(mail.TypeTextPlain, msg)
//
// 	// Secondly the mail client
// 	host, err := mail.NewClient("smtp.gmail.com",
// 		mail.WithPort(587), mail.WithSMTPAuth(mail.SMTPAuthPlain),
// 		mail.WithUsername("kantaidaulet@gmail.com"), mail.WithPassword("xfum mvti aixw atsn"))
// 	if err != nil {
// 		log.Errorf("failed to create mail client: %s", err)
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
//
// 	// Finally let's send out the mail
// 	if err := host.DialAndSend(m); err != nil {
// 		log.Errorf("failed to send mail: %s", err)
// 		getErrorResponse(c, err.Error())
// 		return
// 	}
//
// 	c.HTML(http.StatusOK, "email.html", gin.H{})
//
// }

func (h *UserHandler) GetMail(c *gin.Context) {
	c.HTML(http.StatusOK, "email.html", gin.H{})

}

func getErrorResponse(c *gin.Context, errMessage string) {
	log.WithFields(logrus.Fields{
		"scope":   "Database",
		"message": errMessage,
	}).Error("ERROR")

	c.HTML(
		http.StatusInternalServerError,
		"error.html", gin.H{
			"error": "There is some error happened:" + errMessage,
		})

}
