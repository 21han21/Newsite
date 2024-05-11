package categorieshandlers

import (
	"encoding/json"
	"net/http"

	"mynews/models"
)

// Get возвращает список всех категорий
func GetCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех категорий из базы данных
	categ, err := models.GetCategoriesHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categ)
}

func GetCategori(w http.ResponseWriter, r *http.Request) {
	// Получение списка всех категорий из базы данных
	categ, err := models.GetCategoriesHandler()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправка ответа в формате JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(categ)
}
