package db

import (
	"database/sql"
)

// Delete - функция запроса на удаление машины
func Delete(db *sql.DB, id int64) error {
	insert, err := db.Query("DELETE FROM cars WHERE id = ?", id)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}
