package config

import (
	"time"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type (
	Config struct {
		Server  Server
		LinkGen VerificationLinkGenerator
		Nats    Nats
		Redis   Redis
		Cache   Cache
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

	// Redis configuration for main application
	Redis struct {
		Host         string        `env:"REDIS_HOSTS,notEmpty" envSeparator:","`
		Password     string        `env:"REDIS_PASSWORD"`
		TLSEnable    bool          `env:"REDIS_TLS_ENABLE" envDefault:"true"`
		DialTimeout  time.Duration `env:"REDIS_DIAL_TIMEOUT" envDefault:"60s"`
		WriteTimeout time.Duration `env:"REDIS_WRITE_TIMEOUT" envDefault:"60s"`
		ReadTimeout  time.Duration `env:"REDIS_READ_TIMEOUT" envDefault:"30s"`
	}

	Cache struct {
		UserTTL time.Duration `env:"REDIS_CACHE_CLIENT_TTL" envDefault:"10s"`

		CMSVariableRefreshTime time.Duration `env:"CLIENT_REFRESH_TIME" envDefault:"1m"`
	}
)

func New() (*Config, error) {
	godotenv.Load()
	var cfg Config
	err := env.Parse(&cfg)

	return &cfg, err
}
