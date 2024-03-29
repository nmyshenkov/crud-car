package db

import (
	"crud-car/cars"
	"database/sql"
)

// SelectOne - функция запроса на получение машины по идентификатору
func SelectOne(db *sql.DB, id int64) (cars.Car, error) {
	var c cars.Car

	err := db.QueryRow("SELECT id, company, model, capacity, create_time FROM cars WHERE id = ?", id).Scan(&c.ID, &c.Company, &c.Model, &c.Capacity, &c.CreateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return c, nil
		}
		return c, err
	}
	return c, nil
}

// Select - функция запроса на получения списка машин
func Select(db *sql.DB, limit, offset int64) ([]cars.Car, error) {
	var res []cars.Car

	rows, err := db.Query("SELECT id, company, model, capacity, create_time FROM cars LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		return res, err
	}

	for rows.Next() {
		var c cars.Car
		if err := rows.Scan(&c.ID, &c.Company, &c.Model, &c.Capacity, &c.CreateTime); err != nil {
			return res, err
		}
		res = append(res, c)
	}
	defer rows.Close()

	return res, nil
}
