package http_booking

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/juddbaguio/go-saga-choreography/business/domain/usecases/booking_usecase"
	"go.uber.org/zap"
)

type Container struct {
	logger         *zap.Logger
	bookingService *booking_usecase.Container
}

func NewHTTP(logger *zap.Logger, bookingService *booking_usecase.Container) *Container {
	return &Container{
		logger:         logger,
		bookingService: bookingService,
	}
}

func (c *Container) ListenAndServe() error {
	r := chi.NewRouter()

	c.InitRoutes(r)
	srv := http.Server{
		Addr:    "127.0.0.1:3031",
		Handler: r,
	}

	srvErr := make(chan error)
	sigCh := make(chan os.Signal, 1)

	signal.Notify(sigCh, syscall.SIGABRT, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		c.logger.Info("Starting server at port :3031")
		srvErr <- srv.ListenAndServe()
	}()

	select {
	case sig := <-sigCh:
		c.logger.Info(sig.String())
		c.logger.Warn("Starting Graceful Shutdown")
		defer c.logger.Info("shutdown complete")
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			c.logger.Warn("Initiating force server shutdown")
			if err = srv.Close(); err != nil {
				c.logger.Sugar().Errorw("server failed to shutdown", err)
				return err
			}
			return nil
		}
	case err := <-srvErr:
		c.logger.Warn("Something went wrong starting with the server: ")
		c.logger.Warn(err.Error())
		return err
	}

	return nil
}
