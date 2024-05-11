package newshandlers

import (
	"encoding/json"
	"mynews/models"
	"net/http"
)

// GetNews возвращает список всех новостей
func GetNewsHandler(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех новостей из базы данных
	news, err := models.GetNewsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}

func GetNews(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех новостей из базы данных
	news, err := models.GetNewsHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(news)
}
