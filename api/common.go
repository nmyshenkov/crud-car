package api

import (
	"crud-car/cars"
	"encoding/json"
	"errors"
	"log"
	"net/http"
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

func checkHttpPost(w http.ResponseWriter, r *http.Request) error {
	if r.Method != "POST" {
		var e cars.Error
		e.Error = MethodNotAllow
		w.WriteHeader(http.StatusMethodNotAllowed)
		if err := json.NewEncoder(w).Encode(e); err != nil {
			log.Println(err)
		}
		log.Println(e.Error)
		return errors.New(e.Error)
	}
	return nil
}
