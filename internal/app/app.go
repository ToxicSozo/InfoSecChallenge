package app

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/Hollywood-Kid/InfoSecChallenge/internal/config"
	"github.com/Hollywood-Kid/InfoSecChallenge/internal/handler"
	"github.com/go-chi/chi/v5"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	r := chi.NewRouter()
	handler.RegisterRoutes(r)

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down serser")
		s.Shutdown(ctx)
	}()

	slog.Info("starting server", slog.String("addr", cfg.ServerAddr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
