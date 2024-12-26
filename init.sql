-- Создание таблицы users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    pass TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE
);

-- Создание таблицы sessions
CREATE TABLE IF NOT EXISTS sessions (
    session_id TEXT PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    expires_at TIMESTAMP NOT NULL
);

-- Создание таблицы types
CREATE TABLE IF NOT EXISTS types (
    id SERIAL PRIMARY KEY,
    type TEXT NOT NULL,
    objective TEXT
);

-- Создание таблицы theory
CREATE TABLE IF NOT EXISTS theory (
    id SERIAL PRIMARY KEY,
    type INT REFERENCES types(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    correct_answer TEXT NOT NULL,
    answers JSONB NOT NULL
);

-- Создание таблицы practice
CREATE TABLE IF NOT EXISTS practice (
    id SERIAL PRIMARY KEY,
    type INT REFERENCES types(id) ON DELETE CASCADE,
    description TEXT NOT NULL,
    flag TEXT NOT NULL
);

-- Создание таблицы user_answers
CREATE TABLE IF NOT EXISTS user_answers (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    theory_answers JSONB,
    practice_answers JSONB,
    PRIMARY KEY (user_id)
);

-- Создание таблицы для хранения итоговых результатов
CREATE TABLE IF NOT EXISTS results (
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    type INT REFERENCES types(id) ON DELETE CASCADE,
    score INT NOT NULL,
    completed_at TIMESTAMP DEFAULT NOW(),
    PRIMARY KEY (user_id, type)
);