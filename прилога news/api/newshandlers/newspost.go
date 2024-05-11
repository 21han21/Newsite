package newshandlers

import (
	"encoding/json"
	
	"net/http"
	"mynews/models"
)

// CreateNewsHandler создает новую новость
func CreateNewsHandler(w http.ResponseWriter, r *http.Request) {
    // Декодируем JSON тело запроса в структуру News
    var news models.News
    err := json.NewDecoder(r.Body).Decode(&news)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Вызываем функцию CreateNewsHandler из модели для добавления новой новости в базу данных
    err = models.CreateNewsHandler(&news)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Отправляем ответ с кодом статуса 201 Created
    w.WriteHeader(http.StatusCreated)
}