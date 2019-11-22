package db

import (
	"crud-car/cars"
	"database/sql"
)

// Insert - функция запроса на создание машины
func Insert(db *sql.DB, r cars.Car) error {
	insert, err := db.Query("INSERT INTO cars (company, model, capacity) VALUES (?, ?, ?)", r.Company, r.Model, r.Capacity)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}
