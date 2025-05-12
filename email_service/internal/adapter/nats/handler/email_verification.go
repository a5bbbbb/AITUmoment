package handler

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/nats/handler/dto"

	"github.com/nats-io/nats.go"
)

type EmailVerification struct {
	evs EmailVerificationService
}

func NewEmailVerification(evs EmailVerificationService) *EmailVerification {
	return &EmailVerification{
		evs: evs,
	}
}

func (ev *EmailVerification) Handler(ctx context.Context, msg *nats.Msg) error {
	emailVerification, err := dto.ToEmailVerification(msg)
	if err != nil {
		logger.GetLogger().Println("failed to convert Client NATS msg: ", err)
		return err
	}

	err = ev.evs.SendEmailVerification(emailVerification)
	if err != nil {
		logger.GetLogger().Println("failed to send email verification ", err)
		return err
	}

	return nil
}
