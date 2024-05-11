package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"mynews/api/authorshandlers"
	"mynews/api/categorieshandlers"
	"mynews/api/handlers"
	"mynews/api/newshandlers"
	"mynews/db"
)

func main() {
	// Инициализация подключения к базе данных
	db.InitDB()
	defer db.CloseDB()

	// Инициализация маршрутизатора
	r := mux.NewRouter()

	// обработчики для таблицы пользователей
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/users", handlers.CreateUser).Methods("POST")
	//
	r.HandleFunc("/users/", handlers.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/", handlers.DeleteUser).Methods("DELETE")

	// обрабодчики для таблицы новостей
	r.HandleFunc("/news", newshandlers.GetNewsHandler).Methods("GET")
	r.HandleFunc("/news", newshandlers.CreateNewsHandler).Methods("POST")
	//
	r.HandleFunc("/news", newshandlers.UpdateNewsHandler).Methods("PUT")
	r.HandleFunc("/news", newshandlers.DeleteNewsHandler).Methods("DELETE")

	//обработчики для таблицы категории
	r.HandleFunc("/categories", categorieshandlers.CreateCategoriesHandler).Methods("POST")
	r.HandleFunc("/categories", categorieshandlers.GetCategoriesHandler).Methods("GET")
	//
	r.HandleFunc("/categories", categorieshandlers.UpdateCategoriesHandler).Methods("PUT")
	r.HandleFunc("/categories", categorieshandlers.DeleteCategoriesHandler).Methods("DELETE")

	//обработчик для таблицы авторов
	r.HandleFunc("/authors", authorshandlers.CreateAuthorsHandler).Methods("POST")
	r.HandleFunc("/authors", authorshandlers.GetAuthorsHandler).Methods("GET")
    //
    r.HandleFunc("/authors", authorshandlers.UpdateAuthorsHandler).Methods("PUT")
	r.HandleFunc("/authors", authorshandlers.DeleteAuthorsHandler).Methods("DELETE")


	// Запуск веб-сервера
	log.Fatal(http.ListenAndServe(":6666", r))
}
