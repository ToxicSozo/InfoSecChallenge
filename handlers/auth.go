package handlers

import (
	"html/template"
	"net/http"

	"github.com/ToxicSozo/InfoSecChallenge/db"
	"github.com/ToxicSozo/InfoSecChallenge/models"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// Отображение формы регистрации
		tmpl := template.Must(template.ParseFiles("templates/register.html"))
		tmpl.Execute(w, nil)
	} else if r.Method == "POST" {
		// Обработка данных формы
		username := r.FormValue("username")
		password := r.FormValue("password")
		email := r.FormValue("email")

		// Регистрация пользователя
		err := models.RegisterUser(db.DB, username, password, email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Перенаправление на страницу входа
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
