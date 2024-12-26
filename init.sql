-- Создание таблицы users
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    pass TEXT NOT NULL,
    username TEXT NOT NULL UNIQUE
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

-- Добавление начальных данных для таблицы types
INSERT INTO types (type, objective) VALUES
('Crypto', 'Задачи на криптографию'),
('OSINT', 'Задачи на поиск информации в открытых источниках');

-- Добавление начальных данных для таблицы practice (флаги)
INSERT INTO practice (type, description, flag) VALUES
(1, 'Никак не могу разглядеть', 'InfoSec_CTF{6l1nd_w0nt_s3e_th1s}'),
(1, 'F1Z1K1 ОТДЫХАЮТ', 'InfoSec_CTF{flag2}'),
(2, 'Тайна Гермеуса Моры', 'InfoSec_CTF{flag3}'),
(2, 'Кто же он такой???', 'InfoSec_CTF{flag4}');

-- Добавление начальных данных для таблицы theory (тестовые вопросы)
INSERT INTO theory (type, description, correct_answer, answers) VALUES
(1, 'Какой тип шифрования использует асимметричные ключи?', 'RSA', '["DNS Spoofing", "ARP Poisoning", "SQL Injection", "RSA"]'),
(1, 'Какой из перечисленных методов является наиболее эффективным для защиты от атак типа "человек посередине" (Man-in-the-Middle)?', 'Использование VPN', '["Использование VPN", "Регулярное обновление антивируса", "Установка брандмауэра", "Отключение JavaScript в браузере"]'),
(2, 'Какой из следующих стандартов регулирует управление информационной безопасностью в организации?', 'ISO 27001', '["ISO 27001", "PCI DSS", "GDPR", "HIPAA"]');