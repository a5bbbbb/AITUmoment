package main

import (
	"context"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/a5bbbbb/AITUmoment/core_service/config"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/db/repository"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/email_verification"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/grpc/server"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/nats/producer"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/adapter/redis"
	"github.com/a5bbbbb/AITUmoment/core_service/internal/services"
	natsconn "github.com/a5bbbbb/AITUmoment/core_service/pkg/nats"
	redisconn "github.com/a5bbbbb/AITUmoment/core_service/pkg/redis"
	"github.com/a5bbbbb/AITUmoment/core_service/pkg/sqlx"
)

func main() {

	defer sqlx.Close()
	cfg, _ := config.New()

	logger.GetLogger().Println("connecting to NATS", "hosts", strings.Join(cfg.Nats.Hosts, ","))

	natsClient, err := natsconn.NewClient(context.TODO(), cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		logger.GetLogger().Println("NATS connection failed")
		logger.GetLogger().Panicf("nats.NewClient: %w", err)
	}
	logger.GetLogger().Println("NATS connection status is", natsClient.Conn.Status().String())

	// redis client
	redisClient, err := redisconn.NewClient(context.TODO(), (redisconn.Config)(cfg.Redis))
	if err != nil {
		logger.GetLogger().Errorf("redisconn.NewClient: %w", err)
		return
	}
	logger.GetLogger().Println("Redis is connected:", redisClient.Ping(context.TODO()) == nil)

	userRedisCache := redis.NewUser(redisClient, cfg.Cache.UserTTL)

	verifGen := email_verification.NewEmailVerificationGenerator(cfg.LinkGen)

	natsEmailVerification := producer.NewEmailVerification(natsClient, cfg.Nats.NatsSubjects.EmailVerificationCommandSubject)

	userRepo := repository.NewUserRepository()

	threadRepo := repository.NewThreadRepo()

	groupRepo := repository.NewGroupRepo()

	eduProgramRepo := repository.NewEduProgramRepo()

	userService := services.NewUserService(
		userRepo,
		groupRepo,
		eduProgramRepo,
		natsEmailVerification,
		verifGen,
		userRedisCache,
	)

	threadService := services.NewThreadService(threadRepo)

	groupService := services.NewGroupService(groupRepo)

	eduProgramService := services.NewEduService(eduProgramRepo)

	grpcServer := server.New(
		cfg.Server.GRPCServer,
		userService,
		threadService,
		groupService,
		eduProgramService,
	)

	errCh := make(chan error)
	grpcServer.Run(context.TODO(), errCh)

	shutdownCh := make(chan os.Signal, 1)

	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-errCh:
		logger.GetLogger().Info("oops, got an error: ", err)
	case <-shutdownCh:
		logger.GetLogger().Info("Shutting down gracefully")
	}

}
