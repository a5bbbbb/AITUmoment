package config

import (
	"time"

	env "github.com/caarlos0/env/v11"
)

type (
	Config struct {
		Server  Server
		LinkGen VerificationLinkGenerator
		Nats    Nats
	}

	Server struct {
		GRPCServer GRPCServer
	}

	GRPCServer struct {
		Port                  int16         `env:"GRPC_PORT,notEmpty"`
		MaxRecvMsgSizeMiB     int           `env:"GRPC_MAX_MESSAGE_SIZE_MIB" envDefault:"12"`
		MaxConnectionAge      time.Duration `env:"GRPC_MAX_CONNECTION_AGE" envDefault:"30s"`
		MaxConnectionAgeGrace time.Duration `env:"GRPC_MAX_CONNECTION_AGE_GRACE" envDefault:"10s"`
	}

	VerificationLinkGenerator struct {
		Origin    string `env:"VERIFICATION_ORIGIN",required`
		JWTSecret string `env:"VERIFICATION_JWT_SECRET",required`
	}

	// Nats configuration for main application
	Nats struct {
		Hosts        []string `env:"NATS_HOSTS,notEmpty" envSeparator:","`
		NKey         string   `env:"NATS_NKEY,notEmpty"`
		IsTest       bool     `env:"NATS_IS_TEST,notEmpty" envDefault:"true"`
		NatsSubjects NatsSubjects
	}

	// NatsSubjects for main application
	NatsSubjects struct {
		EmailVerificationCommandSubject string `env:"NATS_EMAIL_VERIFICATION_EVENT_SUBJECT,notEmpty"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
