package authorshandlers

import (
	"encoding/json"
	"net/http"

	"mynews/models"
)

// Get возвращает список всех авторов
func GetAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех авторов из базы данных
	author, err := models.GetAuthorsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}

func GetAuthors(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех авторов из базы данных
	author, err := models.GetAuthorsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(author)
}
