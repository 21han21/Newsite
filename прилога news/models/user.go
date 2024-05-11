package models

import (
	"time"

	"mynews/db" // Подставьте правильный путь к пакету db
)

// User структура представляет пользователя сайта
type User struct {
	ID               int       `json:"id"`
	Username         string    `json:"username"`
	PasswordHash     string    `json:"password_hash"`
	Email            string    `json:"email"`
	RegistrationDate time.Time `json:"registration_date"`
	Role             string    `json:"role"`
}

// GetAllUsers возвращает список всех пользователей
func GetAllUsers() ([]User, error) {
	// Запрос к базе данных для получения всех пользователей
	rows, err := db.DB.Query("SELECT id, username, password_hash, email, registration_date, role FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Срез для хранения пользователей
	var users []User

	// Обход результатов запроса и добавление пользователей в срез
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Username, &user.PasswordHash, &user.Email, &user.RegistrationDate, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Обработка ошибок, которые могли возникнуть в процессе обхода результатов
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

// CreateUser создает нового пользователя в базе данных
func CreateUser(user *User) error {
	// SQL запрос для вставки нового пользователя
	query := `
        INSERT INTO users (username, password_hash, email, registration_date, role)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `

	// Подготавливаем запрос
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполняем запрос и сохраняем ID нового пользователя
	err = stmt.QueryRow(user.Username, user.PasswordHash, user.Email, user.RegistrationDate, user.Role).Scan(&user.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser обновляет информацию о пользователе в базе данных
func UpdateUser(user *User) error {
	// SQL запрос для обновления информации о пользователе
	query := `
        UPDATE users
        SET username = $2, password_hash = $3, email = $4, registration_date = $5, role = $6
        WHERE id = $1
    `

	// Выполняем SQL запрос для обновления информации о пользователе
	_, err := db.DB.Exec(query, user.ID, user.Username, user.PasswordHash, user.Email, user.RegistrationDate, user.Role)
	if err != nil {
		return err
	}

	return nil
}

// / DeleteUser удаляет пользователя из базы данных по его ID
func DeleteUser(userID int) error {
	// SQL запрос для удаления пользователя по его ID
	query := "DELETE FROM users WHERE id = $1"

	// Выполняем SQL запрос для удаления пользователя
	_, err := db.DB.Exec(query, userID)
	if err != nil {
		return err
	}

	return nil
}
