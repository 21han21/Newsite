package authorshandlers

import (
	"encoding/json"
	
	"net/http"
	"mynews/models"
)

// CreateAuthorsHandler создает нового автора
func CreateAuthorsHandler(w http.ResponseWriter, r *http.Request) {
    // Декодируем JSON тело запроса в структуру Authors
    var author models.Authors
    err := json.NewDecoder(r.Body).Decode(&author)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Вызываем функцию CreateAuthorsHandler из модели для добавления нового автора в базу данных
    err = models.CreateAuthorsHandler(&author)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Отправляем ответ с кодом статуса 201 Created
    w.WriteHeader(http.StatusCreated)
}