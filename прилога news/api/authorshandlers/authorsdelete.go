package authorshandlers

import (
    "net/http"
    "strconv"

    "mynews/models"
)

// DeleteAuthors обрабатывает запрос на удаление автора по его ID
func DeleteAuthorsHandler(w http.ResponseWriter, r *http.Request) {
    // Получаем ID автора из URL запроса
    authorIDStr := r.URL.Query().Get("id")
    authorID, err := strconv.Atoi(authorIDStr)
    if err != nil {
        http.Error(w, "Invalid author ID", http.StatusBadRequest)
        return
    }

    // Вызываем функцию Delete из модели для удаления автора
    err = models.DeleteAuthorsHandler(authorID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Возвращаем успешный статус
    w.WriteHeader(http.StatusOK)
}
