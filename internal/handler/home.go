package handler

import (
	"net/http"

	"github.com/Hollywood-Kid/InfoSecChallenge/internal/view/home"
)

type homeHandler struct{}

type usersHandler struct{}

func (h homeHandler) handleIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index()
}
