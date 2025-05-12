package producer

import (
	"context"
	"fmt"
	"time"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen/base/frontend/v2"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
	"github.com/a5bbbbb/AITUmoment/core_service/pkg/nats"
	"google.golang.org/protobuf/proto"
)

const PushTimeout = 30 * time.Second

type EmailVerification struct {
	natsClient *nats.Client
	subject    string
}

func NewEmailVerification(
	natsClient *nats.Client,
	subject string,
) *EmailVerification {
	return &EmailVerification{
		natsClient: natsClient,
		subject:    subject,
	}
}

func (ev *EmailVerification) Push(ctx context.Context, verification models.EmailVerification) error {

	pbEmailVerification := base.EmailVerification{
		Email:            verification.Email,
		PublicName:       verification.PublicName,
		VerificationLink: verification.VerificationLink,
	}

	logger.GetLogger().Debug("EmailVerification.Push: Marshaling message")

	data, err := proto.Marshal(&pbEmailVerification)
	if err != nil {
		logger.GetLogger().Error("EmailVerification.Push: failed to marshal message: ", err.Error())
		return fmt.Errorf("EmailVerification.Push: failed to marshal message: %w", err)
	}

	logger.GetLogger().Debug("EmailVerification.Push: Publishing message")

	err = ev.natsClient.Conn.Publish(ev.subject, data)
	if err != nil {
		logger.GetLogger().Error("EmailVerification.Push: failed to published message: ", err.Error())
		return fmt.Errorf("EmailVerification.Push: failed to published message: %w", err)
	}

	logger.GetLogger().Debug("EmailVerification.Push: message published")

	return nil
}
