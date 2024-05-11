package categorieshandlers

import (
	"encoding/json"
	
	"net/http"
	"mynews/models"
)

// CreateCategoriesHandler создает новую категорию
func CreateCategoriesHandler(w http.ResponseWriter, r *http.Request) {
    // Декодируем JSON тело запроса в структуру Categories
    var categ models.Categorie
    err := json.NewDecoder(r.Body).Decode(&categ)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // Вызываем функцию CreateCategoriesHandler из модели для добавления новой категорий в базу данных
    err = models.CreateCategoriesHandler(&categ)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Отправляем ответ с кодом статуса 201 Created
    w.WriteHeader(http.StatusCreated)
}