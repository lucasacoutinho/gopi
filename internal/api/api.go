package api

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lucasacoutinho/gopi/internal/config"
	"github.com/lucasacoutinho/gopi/internal/middleware"
	"github.com/pilu/xrequestid"
	"github.com/urfave/negroni"
	"go.uber.org/zap"
)

const TIMEOUT = 30 * time.Second

type ServerOption func(server *http.Server)

// Start a new http server with graceful shutdown and default parameters
func Start(cfg *config.Config, handler http.Handler, options ...ServerOption) error {
	n := negroni.New()
	n.Use(xrequestid.New(16))
	n.Use(middleware.NewZapSDLogger(cfg.Logger))
	n.Use(negroni.NewRecovery())
	n.UseHandler(handler)

	srv := &http.Server{
		ReadTimeout:  TIMEOUT,
		WriteTimeout: TIMEOUT,
		Addr:         cfg.Addr,
		Handler:      n,
		ErrorLog:     zap.NewStdLog(cfg.Logger.Desugar()),
	}

	for _, o := range options {
		o(srv)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	go func() {
		<-ctx.Done()
		cfg.Logger.Infow("shutdown", "status", "server stopped")
		err := srv.Shutdown(context.Background())
		if err != nil {
			panic(err)
		}
	}()

	cfg.Logger.Infow("startup", "status", "server started", "host", cfg.Addr)
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// WithReadTimeout configure http.Server parameter ReadTimeout
func WithReadTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.ReadTimeout = t
	}
}

// WithWriteTimeout configure http.Server parameter WriteTimeout
func WithWriteTimeout(t time.Duration) ServerOption {
	return func(srv *http.Server) {
		srv.WriteTimeout = t
	}
}
