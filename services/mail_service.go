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

func (m *MailService) sendMail(to, subject, body string) error {

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

	logger.GetLogger().Info("Email sent successfully")
	return nil
}

func (m *MailService) SendEmailVerification(to, link string) error {
	logger.GetLogger().Info("Sending email verification letter to ", to)

	subject := "Email verification for AITUmoment"
	body := `
	<p> Hello,
	This is an email verification letter for AITUmoment social media. 
	To verify your email click on the link below!
	If you did not register on our platform, then please ignore this message.</p>
	<a href="` + link + `">Verify email</a>`

	err := m.sendMail(to, subject, body)

	if err != nil {
		logger.GetLogger().Errorf("Error sending email verification letter to %s with link %s: %s", to, link, err.Error())
		return err
	}

	logger.GetLogger().Info("Successfully sent email verification letter to ", to)

	return nil
}
