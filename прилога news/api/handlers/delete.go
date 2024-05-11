package handlers

import (
    "net/http"
    "strconv"

    "mynews/models"
)

// DeleteUser обрабатывает запрос на удаление пользователя по его ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
    // Получаем ID пользователя из URL запроса
    userIDStr := r.URL.Query().Get("id")
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    // Вызываем функцию DeleteUser из модели для удаления пользователя
    err = models.DeleteUser(userID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Возвращаем успешный статус
    w.WriteHeader(http.StatusOK)
}
