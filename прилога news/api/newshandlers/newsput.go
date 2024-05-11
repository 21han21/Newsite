package newshandlers


import (
	"encoding/json"
	"net/http"

	"mynews/models"
)

// UpdateNews обновляет информацию об опр новости
func UpdateNewsHandler(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON тело запроса в структуру News
	var news models.News
	err := json.NewDecoder(r.Body).Decode(&news)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию UpdateNews из модели для обновления информации о новосте
	err = models.UpdateNewsHandler(&news)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
}
