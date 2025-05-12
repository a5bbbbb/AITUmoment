package app

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/a5bbbbb/AITUmoment/email_service/config"
	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/gomail"
	grpcserver "github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/grpc/server"
	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/logger"
	natshandler "github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/nats/handler"
	"github.com/a5bbbbb/AITUmoment/email_service/internal/service"
	natsconn "github.com/a5bbbbb/AITUmoment/email_service/pkg/nats"
	natsconsumer "github.com/a5bbbbb/AITUmoment/email_service/pkg/nats/consumer"
)

const serviceName = "email-service"

type App struct {
	grpcServer         *grpcserver.API
	natsPubSubConsumer *natsconsumer.PubSub
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger.GetLogger().Println("starting ", serviceName, " service")

	logger.GetLogger().Println("connecting to NATS", "hosts", strings.Join(cfg.Nats.Hosts, ","))

	natsClient, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		logger.GetLogger().Println("NATS connection failed")
		return nil, fmt.Errorf("nats.NewClient: %w", err)
	}
	logger.GetLogger().Println("NATS connection status is", natsClient.Conn.Status().String())

	natsPubSubConsumer := natsconsumer.NewPubSub(natsClient)

	gomailTransmitter := gomail.NewGomailTransmitter(cfg.Email)

	emailService := service.NewEmailService(gomailTransmitter)

	emailVerificationHandler := natshandler.NewEmailVerification(emailService)

	natsPubSubConsumer.Subscribe(natsconsumer.PubSubSubscriptionConfig{
		Subject: cfg.Nats.NatsSubjects.EmailVerificationCommandSubject,
		Handler: emailVerificationHandler.Handler,
	})

	gRPCServer := grpcserver.New(
		cfg.Server.GRPCServer,
		emailService,
	)

	app := &App{
		grpcServer:         gRPCServer,
		natsPubSubConsumer: natsPubSubConsumer,
	}

	return app, nil
}

func (a *App) Close(ctx context.Context) {
	err := a.grpcServer.Stop(ctx)
	if err != nil {
		logger.GetLogger().Println("failed to shutdown gRPC service ", err)
	}
	a.natsPubSubConsumer.Stop()
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()
	a.grpcServer.Run(ctx, errCh)
	a.natsPubSubConsumer.Start(ctx, errCh)

	logger.GetLogger().Printf("service %v started", serviceName)

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		logger.GetLogger().Printf("received signal: %v. Running graceful shutdown...", s)
		logger.GetLogger().Println("graceful shutdown completed!")
	}

	return nil
}
