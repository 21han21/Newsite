package models

import (
	"mynews/db"
)

// User структура представляет пользователя сайта
type Authors struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Bio     string `json:"bio"`
}

// GetAllUsers возвращает список всех пользователей
func GetAuthorsHandler() ([]Authors, error) {
	// Запрос к базе данных для получения всех пользователей
	rows, err := db.DB.Query("SELECT id, name, surname, email, bio FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Срез для хранения пользователей
	var authorr []Authors

	// Обход результатов запроса и добавление пользователей в срез
	for rows.Next() {
		var author Authors
		if err := rows.Scan(&author.ID, &author.Name, &author.Surname, &author.Email, &author.Bio); err != nil {
			return nil, err
		}
		authorr = append(authorr, author)
	}

	// Обработка ошибок, которые могли возникнуть в процессе обхода результатов
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return authorr, nil
}

// CreateNewsHandler создает новую новость  в базе данных
func CreateAuthorsHandler(author *Authors) error {
	// SQL запрос для вставки новой новости
	query := `
        INSERT INTO authors (name, surname,email,bio)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `

	// Подготавливаем запрос
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	// Выполняем запрос и сохраняем ID новой новости
	err = stm.QueryRow(author.Name, author.Surname, author.Email, author.Bio).Scan(&author.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser обновляет информацию о пользователе в базе данных
func UpdateAuthorsHandler(author *Authors) error {
	// SQL запрос для обновления информации о пользователе
	query := `
        UPDATE authors
        SET name = $2, surname = $3, email = $4, bio = $5
        WHERE id = $1
    `

	// Выполняем SQL запрос для обновления информации о пользователе
	_, err := db.DB.Exec(query, author.ID, author.Name, author.Surname, author.Email, author.Bio)
	if err != nil {
		return err
	}

	return nil
}

// / DeleteUser удаляет пользователя из базы данных по его ID
func DeleteAuthorsHandler(authorID int) error {
	// SQL запрос для удаления пользователя по его ID
	query := "DELETE FROM authors WHERE id = $1"

	// Выполняем SQL запрос для удаления пользователя
	_, err := db.DB.Exec(query, authorID)
	if err != nil {
		return err
	}

	return nil
}
