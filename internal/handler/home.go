package handler

import (
	"database/sql"
	"fmt"
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
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		return home.Test(0).Render(r.Context(), w)
	}

	score, err := models.GetUserScore(h.db, userID)
	if err != nil {
		return err
	}

	return home.Test(score).Render(r.Context(), w)
}

func (h homeHandler) handleGetScore(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-name")
	userID, ok := session.Values["user_id"].(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return nil
	}

	score, err := models.GetUserScore(h.db, userID)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"score": %d}`, score)))
	return nil
}
