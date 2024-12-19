package routes

import (
	"net/http"

	"github.com/ToxicSozo/InfoSecChallenge/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/register", handlers.RegisterHandler) // Маршрут для регистрации       // Маршрут для входа
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
