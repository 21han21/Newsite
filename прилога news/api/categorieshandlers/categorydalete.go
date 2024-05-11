package categorieshandlers


import (
    "net/http"
    "strconv"

    "mynews/models"
)

// DeleteCategor обрабатывает запрос на удаление категории по ID
func DeleteCategoriesHandler(w http.ResponseWriter, r *http.Request) {
    // Получаем ID пользователя из URL запроса
    categIDStr := r.URL.Query().Get("id")
    categID, err := strconv.Atoi(categIDStr)
    if err != nil {
        http.Error(w, "Invalid categories ID", http.StatusBadRequest)
        return
    }

    // Вызываем функцию DeleteCategories из модели для удаления категории
    err = models.DeleteCategoriesHandler(categID)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Возвращаем успешный статус
    w.WriteHeader(http.StatusOK)
}