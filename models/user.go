package models

import (
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
	Email    string
}

// Хэширование пароля
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func RegisterUser(db *sql.DB, username, password, email string) error {
	// Хэширование пароля
	hashedPassword, err := hashPassword(password)
	if err != nil {
		log.Println("Ошибка хэширования пароля:", err)
		return err
	}

	// Вставка данных в базу
	_, err = db.Exec("INSERT INTO users (username, password, email) VALUES ($1, $2, $3)", username, hashedPassword, email)
	if err != nil {
		log.Println("Ошибка вставки данных в базу:", err)
		return err
	}

	log.Println("Пользователь успешно зарегистрирован:", username)
	return nil
}
