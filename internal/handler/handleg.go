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
	home := homeHandler{db: deps.DB} // Передаем db в homeHandler

	r.Get("/", handler(home.handleIndex))
	r.Get("/about", AuthMiddleware(handler(home.handleAbout)))
	r.Get("/test", AuthMiddleware(handler(home.handleTest)))
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
	r.Post("/submit-flag", AuthMiddleware(handler(SubmitFlagHandler(deps.DB))))
	r.Get("/get-score", AuthMiddleware(handler(home.handleGetScore)))
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

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session-name")
		_, ok := session.Values["user_id"]
		if !ok {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	}
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

		http.Redirect(w, r, "/login", http.StatusFound)
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

		session, _ := store.Get(r, "session-name")
		session.Values["user_id"] = user.ID
		session.Save(r, w)

		http.Redirect(w, r, "/", http.StatusFound)
		return nil
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-name")
	delete(session.Values, "user_id")
	session.Save(r, w)

	http.Redirect(w, r, "/login", http.StatusFound)
	return nil
}

func SubmitFlagHandler(db *sql.DB) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		err := r.ParseForm()
		if err != nil {
			return err
		}

		flag := r.FormValue("flag")
		taskID := r.FormValue("task_id")
		userID := r.Context().Value("user_id").(int)

		// Проверка флага
		var correctFlag string
		err = db.QueryRow("SELECT flag FROM practice WHERE id = $1", taskID).Scan(&correctFlag)
		if err != nil {
			return err
		}

		if flag == correctFlag {
			// Начисление баллов
			err = models.UpdateUserScore(db, userID, 1, 10) // 10 баллов за правильный ответ
			if err != nil {
				return err
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Флаг верный! Баллы начислены."))
		} else {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Неверный флаг."))
		}

		return nil
	}
}
