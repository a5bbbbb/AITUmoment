package gomail

import (
	"errors"

	"github.com/a5bbbbb/AITUmoment/email_service/config"
	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/logger"

	"github.com/sirupsen/logrus"
	gomail "gopkg.in/mail.v2"
)

type GomailTransmitter struct {
	dialer *gomail.Dialer
}

func NewGomailTransmitter(cfg config.EmailTransmitter) *GomailTransmitter {
	return &GomailTransmitter{
		dialer: gomail.NewDialer(cfg.EmailHost, cfg.EmailHostPort, cfg.Email, cfg.EmailPassword),
	}
}

func (gt *GomailTransmitter) Send(to, subject, body string) error {

	logger.GetLogger().WithFields(logrus.Fields{
		"address": to,
		"subject": subject,
		"body":    body,
	}).Debug("GomailTransmitter is sending email")

	message := gomail.NewMessage()

	message.SetHeader("From", gt.dialer.Username)
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)

	message.SetBody("text/html", body)

	err := gt.dialer.DialAndSend(message)
	if err != nil {
		logger.GetLogger().Errorf("GomailTransmitter failed to send a letter: %v", err)
		return errors.New("Failed to send email: " + err.Error())
	}

	logger.GetLogger().Info("Email sent successfully")
	return nil
}
