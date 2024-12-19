package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Драйвер PostgreSQL
)

var DB *sql.DB

func InitDB(connStr string) {
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

	// Проверка соединения
	err = DB.Ping()
	if err != nil {
		log.Fatal("Не удалось проверить соединение с базой данных:", err)
	}

	log.Println("Успешно подключено к базе данных")
}
