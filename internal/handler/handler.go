package handler

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handlerfunc func(http.ResponseWriter, *http.Request) error

func RegisterRoutes(r *chi.Mux) {
	home := homeHandler{}

	r.Get("/", handler(home.handleIndex))
}

func handler(h handlerfunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handlerError(w, r, err)
		}
	}
}

func handlerError(w http.ResponseWriter, r *http.Request, err error) {
	slog.Error("error during request", slog.String("err", err.Error()))
}
