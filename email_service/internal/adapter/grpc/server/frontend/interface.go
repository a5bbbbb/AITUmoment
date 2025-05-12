package frontend

import "github.com/a5bbbbb/AITUmoment/email_service/internal/model"

type EmailService interface {
	SendEmailVerification(verification model.EmailVerification) error
}
