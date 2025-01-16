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
	Score    int
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

// Проверка, был ли уже правильный ответ
func HasCorrectAnswer(db *sql.DB, userID, taskID int) (bool, error) {
	query := `SELECT is_correct FROM user_answers WHERE user_id = $1 AND task_id = $2`
	row := db.QueryRow(query, userID, taskID)

	var isCorrect bool
	err := row.Scan(&isCorrect)
	if err == sql.ErrNoRows {
		return false, nil // Ответа нет
	}
	if err != nil {
		return false, fmt.Errorf("failed to check correct answer: %w", err)
	}

	return isCorrect, nil
}

// Добавление ответа пользователя
func AddUserAnswer(db *sql.DB, userID, taskID int, isCorrect bool) error {
	query := `INSERT INTO user_answers (user_id, task_id, is_correct) VALUES ($1, $2, $3)
	          ON CONFLICT (user_id, task_id) DO UPDATE SET is_correct = $3`
	_, err := db.Exec(query, userID, taskID, isCorrect)
	if err != nil {
		return fmt.Errorf("failed to add user answer: %w", err)
	}
	return nil
}

// Проверка, отвечал ли пользователь на тест
func HasAnsweredTest(db *sql.DB, userID, testID int) (bool, error) {
	query := `SELECT user_id FROM user_test_answers WHERE user_id = $1 AND test_id = $2`
	row := db.QueryRow(query, userID, testID)

	var id int
	err := row.Scan(&id)
	if err == sql.ErrNoRows {
		return false, nil // Пользователь не отвечал на тест
	}
	if err != nil {
		return false, fmt.Errorf("failed to check test answer: %w", err)
	}

	return true, nil
}

// Сохранение факта ответа на тест
func AddUserTestAnswer(db *sql.DB, userID, testID int) error {
	query := `INSERT INTO user_test_answers (user_id, test_id) VALUES ($1, $2)
              ON CONFLICT (user_id, test_id) DO NOTHING`
	_, err := db.Exec(query, userID, testID)
	if err != nil {
		return fmt.Errorf("failed to add user test answer: %w", err)
	}
	return nil
}

func GetLeaderboard(db *sql.DB) ([]User, error) {
	query := `
		SELECT u.id, u.username, COALESCE(SUM(r.score), 0) as score
		FROM users u
		LEFT JOIN results r ON u.id = r.user_id
		GROUP BY u.id
		ORDER BY score DESC
	`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get leaderboard: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		var score int
		if err := rows.Scan(&user.ID, &user.Username, &score); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		user.Score = score
		users = append(users, user)
	}

	return users, nil
}
