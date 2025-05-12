package frontend

import (
	"context"

	"github.com/a5bbbbb/AITUmoment/email_service/internal/model"

	pb "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen/service/frontend/email_service/v2"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmailServiceHandler struct {
	es EmailService

	pb.UnimplementedEmailServiceServer
}

func NewEmailServiceHandler(es EmailService) *EmailServiceHandler {
	return &EmailServiceHandler{
		es: es,
	}
}

func (esh *EmailServiceHandler) SendEmailVerification(ctx context.Context, req *pb.SendEmailVerificationRequest) (*pb.SendEmailVerificationResponse, error) {
	verification := model.EmailVerification{
		Email:            req.Verification.Email,
		PublicName:       req.Verification.PublicName,
		VerificationLink: req.Verification.VerificationLink,
	}

	err := esh.es.SendEmailVerification(verification)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.SendEmailVerificationResponse{}, nil
}
