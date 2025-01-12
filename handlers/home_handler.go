package handlers

import (
	"aitu-moment/db/repository"
	"aitu-moment/models"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-gomail/gomail"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func init() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
}

func SetLogLevelFromEnv() {
	level, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
	if err != nil {
		log.Error("Cannot set log level from .env. Error converting to int.")
		return
	}
	log.SetLevel(logrus.Level(level))
	log.Warn("Set log level to " + logrus.Level(level).String())
}

type UserHandler struct {
	repo *repository.UserRepository
}

func NewUserHandler(repo *repository.UserRepository) *UserHandler {
	SetLogLevelFromEnv()
	return &UserHandler{repo: repo}
}

func (h *UserHandler) GetHome(c *gin.Context) {
	users, err := h.repo.GetAllUsers()
	if err != nil {
		log.WithFields(logrus.Fields{
			"scope":   "handler",
			"message": err.Error(),
		}).Error("ERROR!")
		getErrorResponse(c, err.Error())
		return
	}
	log.WithFields(logrus.Fields{
		"scope":   "handler",
		"message": "Successfully got home screen",
	}).Error("ERROR!")
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
		log.WithFields(logrus.Fields{
			"scope":   "handler",
			"message": err.Error(),
		}).Error("ERROR!")
		getErrorResponse(c, err.Error())
		return
	}

	savedUser, err := h.repo.GetUser(id)
	if err != nil {
		log.WithFields(logrus.Fields{
			"scope":   "handler",
			"message": err.Error(),
		}).Error("ERROR!")
		getErrorResponse(c, err.Error())
		return
	}
	log.WithFields(logrus.Fields{
		"scope":   "handler",
		"message": "Successfully saved user"})

	c.HTML(http.StatusOK, "user.html", gin.H{
		"Name":         savedUser.Name,
		"Program_name": savedUser.Program_name,
	})

}

type Form struct {
	Receiver string                `form:"receiver" binding:"required`
	Message  string                `form:"message" binding:"required`
	File     *multipart.FileHeader `form:"file" binding:"required"`
}

/*
enctype="multipart/form-data"
always include this attribute into the <form> tag from which the file is send.
Otherwise gin does not expect a file to be a form parameter.
*/
func (h *UserHandler) SendMail(c *gin.Context) {
	logThis := log.WithFields(logrus.Fields{
		"package":  "home_handler",
		"function": "SendMail",
	})

	host := os.Getenv("MAIL_HOST")
	port, err := strconv.Atoi(os.Getenv("MAIL_HOST_PORT"))
	if err != nil {
		logThis.Error("Failed to convert port value to int.")
		getErrorResponse(c, err.Error())
		return
	}

	from := os.Getenv("MAIL_SENDER")
	pass := os.Getenv("MAIL_PASSWORD")
	pathToUploads := os.Getenv("UPLOADS_PATH")

	var form Form

	// form.Message = c.PostForm("message")
	// form.Receiver = c.PostForm("receiver")

	// // Source
	// form.File, err = c.FormFile("file")
	// if err != nil {
	// 	log.Error("Bad request or file is incorrect: ", err)
	// }

	err = c.ShouldBind(&form)

	if err != nil {
		logThis.Error("Unable to bind form values to form struct.")
		getErrorResponse(c, err.Error())
		return
	}
	m := gomail.NewMessage()
	m.SetHeader("From", from)
	m.SetHeader("To", form.Receiver)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", "This is an automatic message send to you from AITUmoment! This is the content:<br>"+form.Message+"<br>Nothing else here!")
	if form.File != nil {
		pathToFile := filepath.Join(pathToUploads, filepath.Base(form.File.Filename))

		logThis.WithFields(logrus.Fields{
			"Filename: ": form.File.Filename,
			"Header: ":   form.File.Header,
			"Size: ":     form.File.Size,
		}).Debug("Received file. Before save.")

		c.SaveUploadedFile(form.File, pathToFile)

		m.Attach(pathToFile)

		m.SetBody("text/html", "This is an automatic message send to you from AITUmoment! This is the content:<br>"+form.Message+"Checkout the attachment!")
	} else {
		logThis.Debug("No file received.")
	}

	d := gomail.NewDialer(host, port, from, pass)

	if err := d.DialAndSend(m); err != nil {
		logThis.WithFields(logrus.Fields{
			"from":    from,
			"to":      form.Receiver,
			"subject": form.Message,
			"error":   err.Error(),
		}).Error("Failed to send email")
		getErrorResponse(c, err.Error())
		return
	}

	logThis.WithFields(logrus.Fields{
		"from":    from,
		"to":      form.Receiver,
		"subject": form.Message,
	}).Info("Email sent successfully")

	c.HTML(http.StatusOK, "email.html", gin.H{})
}

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
