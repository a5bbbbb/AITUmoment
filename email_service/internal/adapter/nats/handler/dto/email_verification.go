package dto

import (
	"fmt"

	"github.com/a5bbbbb/AITUmoment/email_service/internal/model"

	base "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen/base/frontend/v2"
	"github.com/nats-io/nats.go"
	"google.golang.org/protobuf/proto"
)

func ToEmailVerification(msg *nats.Msg) (model.EmailVerification, error) {
	var pbEmailVerificationRequest base.EmailVerification
	err := proto.Unmarshal(msg.Data, &pbEmailVerificationRequest)
	if err != nil {
		return model.EmailVerification{}, fmt.Errorf("ToEmailVerification: proto.Unmarshall: %w", err)
	}

	return model.EmailVerification{
		Email:            pbEmailVerificationRequest.Email,
		PublicName:       pbEmailVerificationRequest.PublicName,
		VerificationLink: pbEmailVerificationRequest.VerificationLink,
	}, nil
}
