package handler

import "github.com/a5bbbbb/AITUmoment/email_service/internal/model"

type EmailVerificationService interface {
	SendEmailVerification(verification model.EmailVerification) error
}
