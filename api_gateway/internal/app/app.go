package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/a5bbbbb/AITUmoment/api_gateway/config"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/grpc"
	httpserver "github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/http/server"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/adapter/logger"
	"github.com/a5bbbbb/AITUmoment/api_gateway/internal/services"

	"github.com/a5bbbbb/AITUmoment/common/pkg/grpc/grpcconn"
	eduprogramclient "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/edu_program/v2"
	groupclient "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/group/v2"
	threadclient "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/thread/v2"
	userclient "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/core_service/gen/service/frontend/user/v2"
)

const serviceName = "api-gateway"

type App struct {
	httpServer *httpserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	logger.GetLogger().Println("starting ", serviceName, " service")

	coreMicroserviceConnection, err := grpcconn.New(cfg.GRPC.GRPCClient.CoreServiceURL)
	if err != nil {
		return nil, err
	}

	userServiceClient := grpc.NewAuthServiceClient(userclient.NewUserServiceClient(coreMicroserviceConnection))

	threadServiceClient := grpc.NewThreadServiceClient(threadclient.NewThreadServiceClient(coreMicroserviceConnection))

	groupServiceClient := grpc.NewGroupServiceClient(groupclient.NewGroupServiceClient(coreMicroserviceConnection))

	eduProgramServiceClient := grpc.NewEduProgramServiceClient(eduprogramclient.NewEduProgramServiceClient(coreMicroserviceConnection))

	userService := services.NewUserService(userServiceClient)

	threadService := services.NewThreadService(threadServiceClient)

	groupService := services.NewGroupService(groupServiceClient)

	eduProgramService := services.NewEduProgramService(eduProgramServiceClient)

	httpServer := httpserver.New(cfg.Server, *userService, *threadService, *groupService, *eduProgramService)

	app := &App{
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Close(ctx context.Context) {
	err := a.httpServer.Stop(ctx)
	if err != nil {
		logger.GetLogger().Error("failed to shutdown http server")
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

	ctx := context.Background()

	a.httpServer.Run(errCh)

	logger.GetLogger().Println("service", serviceName, " started")

	shutdownCh := make(chan os.Signal)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	select {
	case errRun := <-errCh:
		return errRun
	case s := <-shutdownCh:
		logger.GetLogger().Println("received signal: ", s, ". Running graceful shutdown...")

		a.Close(ctx)
		logger.GetLogger().Println("Graceful shutdown completed.")
	}

	return nil
}
