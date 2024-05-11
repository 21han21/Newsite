package models

import(
	"mynews/db"
)
// User структура представляет пользователя сайта
type Categorie struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
}


// GetAllUsers возвращает список всех пользователей
func GetCategoriesHandler() ([]Categorie, error) {
	// Запрос к базе данных для получения всех пользователей
	rows, err := db.DB.Query("SELECT id, name FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Срез для хранения пользователей
	var categg []Categorie

	// Обход результатов запроса и добавление пользователей в срез
	for rows.Next() {
		var categ Categorie
		if err := rows.Scan(&categ.ID, &categ.Name); err != nil {
			return nil, err
		}
		categg = append(categg, categ)
	}

	// Обработка ошибок, которые могли возникнуть в процессе обхода результатов
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categg, nil
}


// CreateUser создает нового пользователя в базе данных
func CreateCategoriesHandler(categ *Categorie) error {
	// SQL запрос для вставки нового пользователя
	query := `
        INSERT INTO categories (name)
        VALUES ($1)
        RETURNING id
    `

	// Подготавливаем запрос
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Выполняем запрос и сохраняем ID нового пользователя
	err = stmt.QueryRow(categ.Name).Scan(&categ.ID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateUser обновляет информацию о пользователе в базе данных
func UpdateCategoriesHandler(categ *Categorie) error {
	// SQL запрос для обновления информации о пользователе
	query := `
        UPDATE categories
        SET name = $2
        WHERE id = $1
    `

	// Выполняем SQL запрос для обновления информации о пользователе
	_, err := db.DB.Exec(query, categ.ID, categ.Name)
	if err != nil {
		return err
	}

	return nil
}

// / DeleteUser удаляет пользователя из базы данных по его ID
func DeleteCategoriesHandler(categID int) error {
	// SQL запрос для удаления пользователя по его ID
	query := "DELETE FROM categories WHERE id = $1"

	// Выполняем SQL запрос для удаления пользователя
	_, err := db.DB.Exec(query, categID)
	if err != nil {
		return err
	}

	return nil
}


