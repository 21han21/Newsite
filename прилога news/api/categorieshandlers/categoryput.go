package categorieshandlers


import (
	"encoding/json"
	"net/http"

	"mynews/models"
)

// Update обновляет информацию о категориях
func UpdateCategoriesHandler(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON тело запроса в структуру Categories
	var categ models.Categorie
	err := json.NewDecoder(r.Body).Decode(&categ)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию UpdateCategoriesHandler из модели для обновления информации о категориях
	err = models.UpdateCategoriesHandler(&categ)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
}
