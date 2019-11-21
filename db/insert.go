package db

import (
	"database/sql"
	"private/crud-car/cars"
)

func Insert(db *sql.DB, r cars.Car) error {
	insert, err := db.Query("INSERT INTO cars (company, model, capacity) VALUES (?, ?, ?)", r.Company, r.Model, r.Capacity)
	if err != nil {
		return err
	}
	defer insert.Close()
	return nil
}
