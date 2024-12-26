package app

import (
	"context"
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	"github.com/ToxicSozo/InfoSecChallenge/internal/config"
	"github.com/ToxicSozo/InfoSecChallenge/internal/handler"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

func Run(ctx context.Context) error {
	cfg := config.NewConfig()

	db, err := sql.Open("postgres", cfg.DBConnStr)
	if err != nil {
		return err
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		return err
	}

	if err := initDatabase(db); err != nil {
		return err
	}

	r := chi.NewRouter()
	handler.RegisterRoutes(r, handler.Dependencies{
		AssetsFS: http.Dir(cfg.AssetsDir),
		DB:       db,
	})

	s := http.Server{
		Addr:    cfg.ServerAddr,
		Handler: r,
	}

	go func() {
		<-ctx.Done()
		slog.Info("shutting down server")
		s.Shutdown(ctx)
	}()

	slog.Info("starting server", slog.String("addr", cfg.ServerAddr))
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func initDatabase(db *sql.DB) error {
	sqlFile, err := os.ReadFile("init.sql")
	if err != nil {
		return err
	}

	_, err = db.Exec(string(sqlFile))
	return err
}
