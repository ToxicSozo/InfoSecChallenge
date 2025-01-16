package handler

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"
	"text/template"

	"github.com/ToxicSozo/InfoSecChallenge/internal/models"
	"github.com/ToxicSozo/InfoSecChallenge/internal/view/auth"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

type Dependencies struct {
	AssetsFS http.FileSystem
	DB       *sql.DB
}

type handlerFunc func(http.ResponseWriter, *http.Request) error

var funcMap = template.FuncMap{
	"add": func(a, b int) int {
		return a + b
	},
}

var leaderboardTemplate = template.Must(template.New("leaderboard.html").Funcs(funcMap).ParseFiles("internal/view/leaderboard/leaderboard.html"))

var flags = map[string]string{
	"1": "InfoSec_CTF{6l1nd_w0nt_s3e_th1s}",
	"2": "InfoSec_CTF{Nfwq1aq_b03q_l0r_3v1qr}",
	"3": "InfoSec_CTF{HepBbl_He_u3_CTaJlu}",
	"4": "InfoSec_CTF{Meepo_Dota_2}",
}

// Инициализация хранилища сессий
var store = sessions.NewCookieStore([]byte("your-secret-key"))

func init() {
	// Настройка параметров куки
	store.Options = &sessions.Options{
		Path:     "/",       // Куки доступны для всех путей
		MaxAge:   86400 * 7, // Время жизни куки (7 дней)
		HttpOnly: true,      // Куки доступны только через HTTP (не через JavaScript)
		Secure:   false,     // Установите true, если используете HTTPS
	}
}

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
	r.Get("/get-score", AuthMiddleware(handler(GetUserScoreHandler(deps.DB))))
	r.Get("/leaderboard", LeaderboardHandler(deps.DB))
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
		if err := session.Save(r, w); err != nil {
			slog.Error("Ошибка при сохранении сессии", slog.String("err", err.Error()))
			return err
		}

		http.Redirect(w, r, "/", http.StatusFound)
		return nil
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) error {
	session, _ := store.Get(r, "session-name")
	delete(session.Values, "user_id")
	if err := session.Save(r, w); err != nil {
		slog.Error("Ошибка при сохранении сессии", slog.String("err", err.Error()))
		return err
	}

	http.Redirect(w, r, "/login", http.StatusFound)
	return nil
}

func SubmitFlagHandler(db *sql.DB) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		session, _ := store.Get(r, "session-name")
		userID, ok := session.Values["user_id"].(int)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return nil
		}

		taskIDStr := r.FormValue("task_id")
		flag := r.FormValue("flag")

		if taskIDStr == "" || flag == "" {
			http.Error(w, "Task ID and flag are required", http.StatusBadRequest)
			return nil
		}

		taskID, err := strconv.Atoi(taskIDStr)
		if err != nil {
			return err
		}

		hasCorrect, err := models.HasCorrectAnswer(db, userID, taskID)
		if err != nil {
			return err
		}
		if hasCorrect {
			w.Write([]byte("Вы уже решили эту задачу."))
			return nil
		}

		correctFlag, exists := flags[taskIDStr]
		if !exists {
			http.Error(w, "Task not found", http.StatusNotFound)
			return nil
		}

		if flag == correctFlag {
			err := models.UpdateUserScore(db, userID, 1, 10) // 10 баллов за правильный флаг
			if err != nil {
				return err
			}

			err = models.AddUserAnswer(db, userID, taskID, true)
			if err != nil {
				return err
			}

			w.Write([]byte("Флаг правильный! Баллы добавлены."))
		} else {
			err = models.AddUserAnswer(db, userID, taskID, false)
			if err != nil {
				return err
			}

			w.Write([]byte("Флаг неправильный."))
		}

		return nil
	}
}

func GetUserScoreHandler(db *sql.DB) handlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		session, _ := store.Get(r, "session-name")
		userID, ok := session.Values["user_id"].(int)
		if !ok {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return nil
		}

		// Получение счета из базы данных
		score, err := models.GetUserScore(db, userID)
		if err != nil {
			return err
		}

		// Возврат счета в формате JSON
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"score": %d}`, score)))
		return nil
	}
}

func LeaderboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetLeaderboard(db)
		if err != nil {
			http.Error(w, "Failed to fetch leaderboard", http.StatusInternalServerError)
			return
		}

		data := struct {
			Users []models.User
		}{
			Users: users,
		}

		w.Header().Set("Content-Type", "text/html")
		if err := leaderboardTemplate.Execute(w, data); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}
