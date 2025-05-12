package config

import (
	"time"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		Server Server
		Nats   Nats
		Email  EmailTransmitter

		Version string `env:"VERSION"`
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

	EmailTransmitter struct {
		Email         string `env:"EMAIL",required`
		EmailPassword string `env:"EMAIL_PASSWORD",required`
		EmailHost     string `env:"EMAIL_HOST",required`
		// Set up the SMTP dialer 587 is TSL port for smtp.gmail.com
		EmailHostPort int `env:"EMAIL_HOST_PORT" envDefault:"587"`
	}
)

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
