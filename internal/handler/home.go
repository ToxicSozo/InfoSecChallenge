package handler

import (
	"myapp/internal/handler/view/home"
	"net/http"
)

type homeHandler struct{}

type usersHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
}
