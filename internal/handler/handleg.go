package handler

import (
	"database/sql"
	"log/slog"
	"net/http"

	"github.com/ToxicSozo/InfoSecChallenge/internal/models"
	"github.com/ToxicSozo/InfoSecChallenge/internal/view/auth"
	"github.com/go-chi/chi/v5"
	"golang.org/x/crypto/bcrypt"
)

type Dependencies struct {
	AssetsFS http.FileSystem
	DB       *sql.DB
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

func RegisterRoutes(r *chi.Mux, deps Dependencies) {
	home := homeHandler{}

	r.Get("/", handler(home.handleIndex))
	r.Get("/about", handler(home.handleAbout))
	r.Get("/test", handler(home.handleTest))

	r.Handle("/assets/*", http.StripPrefix("/assets", http.FileServer(deps.AssetsFS)))

	r.Get("/register", func(w http.ResponseWriter, r *http.Request) {
		auth.Register().Render(r.Context(), w)
	})
	r.Get("/login", func(w http.ResponseWriter, r *http.Request) {
		auth.Login().Render(r.Context(), w)
	})

	r.Post("/register", handler(RegisterHandler(deps.DB)))
	r.Post("/login", handler(LoginHandler(deps.DB)))
	r.Post("/logout", handler(LogoutHandler))
}

func handler(h handlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			handleError(w, r, err)
		}
	}
}

func handleError(w http.ResponseWriter, _ *http.Request, err error) {
	slog.Error("error during request", slog.String("err", err.Error()))
	http.Error(w, "Something went wrong", http.StatusInternalServerError)
}

func RegisterHandler(db *sql.DB) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := r.ParseForm()
		if err != nil {
			return err
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return nil
		}

		if err := models.CreateUser(db, username, password); err != nil {
			return err
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User created successfully"))
		return nil
	}
}

func LoginHandler(db *sql.DB) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := r.ParseForm()
		if err != nil {
			return err
		}

		username := r.FormValue("username")
		password := r.FormValue("password")

		if username == "" || password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return nil
		}

		user, err := models.GetUserByUsername(db, username)
		if err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return nil
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return nil
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Login successful"))
		return nil
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Logout successful"))
	return nil
}
