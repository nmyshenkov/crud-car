package db

import (
	"database/sql"
	"private/crud-car/cars"
)

func SelectOne(db *sql.DB, id int64) (cars.Car, error) {
	var c cars.Car

	err := db.QueryRow("SELECT company, model, capacity, create_time FROM cars WHERE id = ?", id).Scan(&c.Company, &c.Model, &c.Capacity, &c.CreateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, nil
		}
		return c, err
	}
	return c, nil
}

func Select(db *sql.DB) ([]cars.Car, error) {
	var res []cars.Car

	rows, err := db.Query("SELECT company, model, capacity, create_time FROM cars")
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		var c cars.Car
		if err := rows.Scan(&c.Company, &c.Model, &c.Capacity, &c.CreateTime); err != nil {
			return res, err
		}
		res = append(res, c)
	}
	defer rows.Close()

	return res, nil
}
