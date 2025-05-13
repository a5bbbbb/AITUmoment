package services

import (
	config "email_service/configs"
	"email_service/internal/logger"
	"errors"

	gomail "gopkg.in/gomail.v2"
)

func SendMail(to, subject, body string) error {

	logger.GetLogger().Debug("Sending email to ", to, " with subject: ", subject, "\n and body: ", body)

	username := config.EmailSender
	password := config.EmailPasswd

	message := gomail.NewMessage()

	message.SetHeader("From", username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/html", body)

	dialer := gomail.NewDialer(config.EmailHost, config.EmailHostPort, username, password)

	err := dialer.DialAndSend(message)
	if err != nil {
		logger.GetLogger().Error(err)
		return errors.New("Failed to send email: " + err.Error())
	}

	logger.GetLogger().Info("Email sent successfully")
	return nil
}
