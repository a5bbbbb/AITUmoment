package service

import (
	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/email_service/internal/model"
)

type EmailService struct {
	transmitter EmailTransmitter
}

func NewEmailService(transmitter EmailTransmitter) *EmailService {
	return &EmailService{
		transmitter: transmitter,
	}
}

func (es *EmailService) SendEmailVerification(verification model.EmailVerification) error {
	logger.GetLogger().Debug("Sending email verification letter to ", verification.Email)

	subject := "Email verification for " + verification.PublicName + " at AITUmoment"
	body := `
	<p> Hello, ` + verification.PublicName + `
	This is an email verification letter for AITUmoment social media platform. 
	To verify your email click on the link below!
	If you did not register on our platform, then please ignore this message.</p>
	<a href="` + verification.VerificationLink + `">Verify email</a>`

	err := es.transmitter.Send(verification.Email, subject, body)

	if err != nil {
		logger.GetLogger().Errorf("Error sending email verification letter to %s with link %s: %s", verification.Email, verification.VerificationLink, err.Error())
		return err
	}

	logger.GetLogger().Info("Successfully sent email verification letter to ", verification.Email)

	return nil
}
