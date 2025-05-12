package email_verification

import (
	"fmt"

	"github.com/a5bbbbb/AITUmoment/core_service/config"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/models"
	"github.com/a5bbbbb/AITUmoment/core_service/pkg/security"
)

type EmailVerificationManager struct {
	jwtManager security.JWTManager
	cfg        config.VerificationLinkGenerator
}

func NewEmailVerificationGenerator(cfg config.VerificationLinkGenerator) *EmailVerificationManager {
	return &EmailVerificationManager{
		jwtManager: *security.NewJWTManager(cfg.JWTSecret),
		cfg:        cfg,
	}
}

func (evg *EmailVerificationManager) Generate(user models.User) (*models.EmailVerification, error) {
	verificationToken, err := evg.jwtManager.GenerateEmailVerificationToken(user.Email, user.PublicName)
	if err != nil {
		logger.GetLogger().Errorln("error while generating email verification token ", err.Error())
		return nil, err
	}

	verification := models.EmailVerification{
		Email:            user.Email,
		PublicName:       user.PublicName,
		VerificationLink: evg.cfg.Origin + "/verify?token=" + verificationToken,
	}

	return &verification, nil
}

func (evg *EmailVerificationManager) GetEmail(token string) (string, error) {
	claims, err := evg.jwtManager.Verify(token)
	if err != nil {
		return "", fmt.Errorf("invalid email verification jwt token %w", err)
	}

	email, ok := claims["email"].(string)
	if !ok {
		return "", fmt.Errorf("cannot get email from token in verification jwt token: %v", token)
	}

	return email, nil
}
