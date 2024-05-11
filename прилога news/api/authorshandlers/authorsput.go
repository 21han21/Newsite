package authorshandlers

import (
	"encoding/json"
	"net/http"

	"mynews/models"
)

// UpdateAuthors обновляет информацию о авторе
func UpdateAuthorsHandler(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON тело запроса в структуру Authors
	var author models.Authors
	err := json.NewDecoder(r.Body).Decode(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию UpdateAuthorsHandler из модели для обновления информации о авторе
	err = models.UpdateAuthorsHandler(&author)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
}