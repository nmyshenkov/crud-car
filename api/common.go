package api

import (
	"encoding/json"
	"log"
	"net/http"
	"private/crud-car/cars"
)

func setResult(w http.ResponseWriter, data []cars.Car) {
	var res cars.Result
	res.Status = "OK"
	if len(data) > 0 {
		res.Data = data
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
	return
}

func setError(w http.ResponseWriter, str string) {
	var e cars.Error
	e.Error = str
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(e); err != nil {
		log.Println(err)
	}
	return
}
