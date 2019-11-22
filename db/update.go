package db

import (
	"crud-car/cars"
	"database/sql"
)

// Update - функция запроса на обновление машины
func Update(db *sql.DB, r cars.Car) error {
	up, err := db.Query("UPDATE cars SET company = ?, model = ?, capacity = ? WHERE id = ?", r.Company, r.Model, r.Capacity, r.ID)
	if err != nil {
		return err
	}
	defer up.Close()
	return nil
}
