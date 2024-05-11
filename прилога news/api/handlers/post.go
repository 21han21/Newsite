package handlers

import (
    "encoding/json"
    "net/http"

    "mynews/models"
)

// CreateUser создает нового пользователя
func CreateUser(w http.ResponseWriter, r *http.Request) {
    // Декодируем JSON тело запроса в структуру User
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Вызываем функцию CreateUser из модели для добавления нового пользователя в базу данных
    err = models.CreateUser(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Отправляем ответ с кодом статуса 201 Created
    w.WriteHeader(http.StatusCreated)
}