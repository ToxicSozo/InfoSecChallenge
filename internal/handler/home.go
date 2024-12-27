package handler

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/ToxicSozo/InfoSecChallenge/internal/models"
	"github.com/ToxicSozo/InfoSecChallenge/internal/view/home"
)

type homeHandler struct {
	db *sql.DB
}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-name")
	if userID, ok := session.Values["user_id"]; ok {
		slog.Info("User ID from session", slog.Int("userID", userID.(int)))
	}

	return home.Index().Render(r.Context(), w)
}

func (h homeHandler) handleAbout(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-name")
	if userID, ok := session.Values["user_id"]; ok {
		slog.Info("User ID from session", slog.Int("userID", userID.(int)))
	}

	return home.About().Render(r.Context(), w)
}

func (h homeHandler) handleTest(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-name")
	var score int
	if userID, ok := session.Values["user_id"]; ok {
		slog.Info("User ID from session", slog.Int("userID", userID.(int)))
		var err error
		score, err = models.GetUserScore(h.db, userID.(int))
		if err != nil {
			return err
		}

		slog.Info("User score", slog.Int("userID", userID.(int)), slog.Int("score", score))
	}

	return home.Test().Render(r.Context(), w)
}
