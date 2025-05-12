package server

import (
	"context"
	"fmt"

	"log"
	"net"

	"github.com/a5bbbbb/AITUmoment/email_service/config"
	"github.com/a5bbbbb/AITUmoment/email_service/internal/adapter/grpc/server/frontend"

	emailsvc "github.com/a5bbbbb/AITUmoment/common/pkg/grpc/proto/email_service/gen/service/frontend/email_service/v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type API struct {
	s            *grpc.Server
	cfg          config.GRPCServer
	addr         string
	emailService EmailService
}

func New(
	cfg config.GRPCServer,
	emailService EmailService,
) *API {
	return &API{
		cfg:          cfg,
		addr:         fmt.Sprintf("0.0.0.0:%d", cfg.Port),
		emailService: emailService,
	}
}

func (a *API) Run(ctx context.Context, errCh chan<- error) {
	go func() {
		log.Println(ctx, "gRPC server starting listen", fmt.Sprintf("addr: %s", a.addr))

		if err := a.run(ctx); err != nil {
			errCh <- fmt.Errorf("can't start grpc server: %w", err)

			return
		}
	}()
}

// Stop method gracefully stops grpc API server. Provide context to force stop on timeout.
func (a *API) Stop(ctx context.Context) error {
	if a.s == nil {
		return nil
	}

	stopped := make(chan struct{})
	go func() {
		a.s.GracefulStop()
		close(stopped)
	}()

	select {
	case <-ctx.Done(): // Stop immediately if the context is terminated
		a.s.Stop()
	case <-stopped:
	}

	return nil
}

// run starts and runs GRPCServer server.
func (a *API) run(ctx context.Context) error {
	a.s = grpc.NewServer(a.setOptions(ctx)...)

	// Register services
	emailsvc.RegisterEmailServiceServer(a.s, frontend.NewEmailServiceHandler(a.emailService))

	// Register reflection service
	reflection.Register(a.s)

	listener, err := net.Listen("tcp", a.addr)
	if err != nil {
		return fmt.Errorf("failed to create listener: %w", err)
	}

	err = a.s.Serve(listener)
	if err != nil {
		return fmt.Errorf("failed to serve grpc: %w", err)
	}

	return nil
}
