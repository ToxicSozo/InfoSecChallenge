package handler

import (
	"net/http"

	"github.com/ToxicSozo/InfoSecChallenge/internal/handler/view/home"
)

type homeHandler struct{}

type usersHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
