package models

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int
	Username string
	Password string
}

func CreateUser(db *sql.DB, username, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("failed to hash password: %w", err)
	}

	query := `INSERT INTO users (username, pass) VALUES ($1, $2)`
	_, err = db.Exec(query, username, hashedPassword)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func GetUserByUsername(db *sql.DB, username string) (*User, error) {
	query := `SELECT id, username, pass FROM users WHERE username = $1`
	row := db.QueryRow(query, username)

	user := &User{}
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func UpdateUserScore(db *sql.DB, userID, typeID, score int) error {
	query := `INSERT INTO results (user_id, type, score) VALUES ($1, $2, $3) 
	          ON CONFLICT (user_id, type) DO UPDATE SET score = results.score + $3`
	_, err := db.Exec(query, userID, typeID, score)
	if err != nil {
		return fmt.Errorf("failed to update user score: %w", err)
	}
	return nil
}

func GetUserScore(db *sql.DB, userID int) (int, error) {
	query := `SELECT COALESCE(SUM(score), 0) FROM results WHERE user_id = $1`
	row := db.QueryRow(query, userID)

	var score int
	err := row.Scan(&score)
	if err != nil {
		return 0, fmt.Errorf("failed to get user score: %w", err)
	}

	return score, nil
}
