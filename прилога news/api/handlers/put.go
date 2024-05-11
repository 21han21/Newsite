package handlers

import (
	"encoding/json"
	"net/http"

	"mynews/models"
)

// UpdateUser обновляет информацию о пользователе
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Декодируем JSON тело запроса в структуру User
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Вызываем функцию UpdateUser из модели для обновления информации о пользователе
	err = models.UpdateUser(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Отправляем успешный ответ
	w.WriteHeader(http.StatusOK)
}
