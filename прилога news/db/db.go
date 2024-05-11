package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var (
	// DB соединение с базой данных
	DB *sql.DB
)

// функция InitDB инициализ., подключение к базе данных PostgreSQL
func InitDB() {
	// Строка подключения к базе данных PostgreSQL
	// "postgres://" - протокол для раб с PostgreSQL
	// "username:password" - учетные данные пользователя базы данных
	// "@host:port" - адрес хоста базы данных и порт
	// "/database" - название базы данных
	connectionString := "postgres://postgres:solo2121@localhost:5432/news_site?sslmode=disable"

	// Открытие соединения с базой данных
	var err error
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}
	// Проверка подключения к базе данных
	err = DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to the database подключение к базе")
}

// CloseDB закрытие соединение с базой данных
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}

