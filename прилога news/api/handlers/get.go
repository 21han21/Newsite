package handlers

import (
	"encoding/json"
	"net/http"

	"mynews/models"
)

// GetUsers возвращает список всех пользователей
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех пользователей из базы данных
	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех пользователей из базы данных
	users, err := models.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}
