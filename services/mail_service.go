package services

import (
	"aitu-moment/logger"
	"errors"
	"os"

	gomail "gopkg.in/mail.v2"
)

type MailService struct {
}

func NewMailService() *MailService {
	return &MailService{}
}

/*
enctype="multipart/form-data"
always include this attribute into the <form> tag from which the file is send.
*/

func (m *MailService) SendMail(to, subject, body string) error {

	logger.GetLogger().Debug("Sending email to ", to, " with subject: ", subject, "\n and body: ", body)

	username := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")

	message := gomail.NewMessage()

	message.SetHeader("From", username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/html", body)

	// Set up the SMTP dialer 587 is TSL port for smtp.gmail.com
	dialer := gomail.NewDialer(os.Getenv("EMAIL_HOST"), 587, username, password)

	err := dialer.DialAndSend(message)
	if err != nil {
		logger.GetLogger().Error(err)
		return errors.New("Failed to send email: " + err.Error())
	}

	logger.GetLogger().Debug("Email sent successfully")
	return nil
}

func (m *MailService) SendEmailVerification(to, link string) error {
	logger.GetLogger().Debug("Sending email verification letter to ", to)

	subject := "Email verification for AITUmoment"
	body := `
	<p> Hello,
	This is an email verification letter for AITUmoment social media. 
	To verify your email click on the link below!
	If you did not register on our platform, then please ignore this message.</p>
	<a href="` + link + `">Verify email</a>`

	err := m.SendMail(to, subject, body)

	if err != nil {
		logger.GetLogger().Errorf("Error sending email verification letter to ", to, " with link ", link, " err: ", err)
		return err
	}

	logger.GetLogger().Debug("Successfully sent email verification letter to ", to)

	return nil
}

/*
func sendMail(to, subject, body, attachmentPath string) error {
	// Sender data.
	username := os.Getenv("PULLOEMAIL")
	password := os.Getenv("PULLOEMAIL_PASSWORD")

	// Receiver email address.
	message := gomail.NewMessage()

	// Set email headers
	message.SetHeader("From", username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	// Set email body
	message.SetBody("text/html", body)

	// Add attachments
	message.Attach(attachmentPath)

	// Set up the SMTP dialer
	dialer := gomail.NewDialer(os.Getenv("PULLO_EMAIL_PROVIDER"), 587, username, password)

	// Send the email
	err := dialer.DialAndSend(message)
	if err != nil {
		logger.WithFields(logrus.Fields{
			"Module": "SendMail",
			"Action": "dialer.DialAndSend",
		}).Error(err)
		return errors.New("Failed to send email: " + err.Error())
	}
	logger.WithFields(logrus.Fields{
		"to":      to,
		"subject": subject,
	}).Info("Email sent successfully")
	return nil
}

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
*/
