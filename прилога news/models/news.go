package models

import (
	"mynews/db"
	"time"
)

// News представляет структуру новости
type News struct {
	ID              int       `json:"id"`
	Title           string    `json:"title"`
	Content         string    `json:"content"`
	PublicationData time.Time `json:"publication_date"`
	AuthorId        int       `json:"author_id"`
	CategoryId      int       `json:"category_id"`
}

// GetNewsHandler возвращает список всех новостей
func GetNewsHandler() ([]News, error) {
	// Запрос к базе данных для получения всех новостей
	rows, err := db.DB.Query("SELECT id, title, content,publication_date, author_id, category_id  FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Срез для хранения новостей
	var novost []News

	// Обход результатов запроса и добавление новостей в срез
	for rows.Next() {
		var news News
		if err := rows.Scan(&news.ID, &news.Title, &news.Content,&news.PublicationData,&news.AuthorId,&news.CategoryId); err != nil {
			return nil, err
		}
		novost = append(novost, news)
	}

	// Обработка ошибок, которые могли возникнуть в процессе обхода результатов
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return novost, nil
}

// CreateNewsHandler создает новую новость  в базе данных
func CreateNewsHandler(news *News) error {
	// SQL запрос для вставки новой новости
	query := `
        INSERT INTO news (title, content)
        VALUES ($1, $2)
        RETURNING id
    `

	// Подготавливаем запрос
	stm, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stm.Close()

	// Выполняем запрос и сохраняем ID новой новости
	err = stm.QueryRow(news.Title, news.Content).Scan(&news.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser обновляет информацию о пользователе в базе данных
func UpdateNewsHandler(news *News) error {
	// SQL запрос для обновления информации о пользователе
	query := `
        UPDATE news
        SET title = $2, content = $3, publication_data = $4, author_id = $5, category_id = $6
        WHERE id = $1
    `

	// Выполняем SQL запрос для обновления информации о пользователе
	_, err := db.DB.Exec(query, news.ID, news.Title, news.Content, news.PublicationData, news.AuthorId, news.CategoryId)
	if err != nil {
		return err
	}

	return nil
}

// / DeleteUser удаляет пользователя из базы данных по его ID
func DeleteNewsHandler(newsID int) error {
	// SQL запрос для удаления пользователя по его ID
	query := "DELETE FROM news WHERE id = $1"

	// Выполняем SQL запрос для удаления пользователя
	_, err := db.DB.Exec(query, newsID)
	if err != nil {
		return err
	}

	return nil
}

