

package newshandlers
import (
    "net/http"
    "strconv"

    "mynews/models"
)

// DeleteNews обрабатывает запрос на удаление новости  по ID
func DeleteNewsHandler(w http.ResponseWriter, r *http.Request) {
    // Получаем ID новости из URL запроса
    newsIDStr := r.URL.Query().Get("id")
    newsID, err := strconv.Atoi(newsIDStr)
    if err != nil {
        http.Error(w, "Invalid news ID", http.StatusBadRequest)
        return
    }

    // Вызываем функцию DeleteNews из модели для удаления новости
    err = models.DeleteNewsHandler(newsID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Возвращаем успешный статус
    w.WriteHeader(http.StatusOK)
}
