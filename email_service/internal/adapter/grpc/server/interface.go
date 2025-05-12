package server

import (
	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/grpc/server/frontend"
)

type EmailService interface {
	frontend.EmailService
}
