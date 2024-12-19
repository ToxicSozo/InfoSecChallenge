package main

import (
	"log"
	"net/http"

	"github.com/ToxicSozo/InfoSecChallenge/db"
	"github.com/ToxicSozo/InfoSecChallenge/routes"
)

func main() {
	// Строка подключения к базе данных
	connStr := "user=myuser dbname=myapp sslmode=disable password=mypassword host=localhost port=5433"

	// Инициализация базы данных
	db.InitDB(connStr)

	// Настройка маршрутов
	routes.SetupRoutes()

	// Запуск сервера
	log.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
