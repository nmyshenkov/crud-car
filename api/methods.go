package api

import (
	"crud-car/cars"
	"crud-car/db"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// List - метод получения списка машин
func (m Typ) List(w http.ResponseWriter, r *http.Request) {
	var in ListRequest

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if checkHttpPost(w, r) != nil {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if in.Limit > 100 {
		log.Println(err)
		setError(w, "Значение limit не может быть больше 100")
		return
	}

	if in.Limit == 0 {
		in.Limit = 100
	}

	res, err := db.Select(m.db, in.Limit, in.Offset)
	if err != nil {
		log.Println(err)
		setError(w, InternalServiceError)
		return
	}
	setResult(w, res)
	return
}

// Get - метод получения машины по идентификатору
func (m Typ) Get(w http.ResponseWriter, r *http.Request) {
	var (
		in  IDRequest
		out []cars.Car
	)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if checkHttpPost(w, r) != nil {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if in.ID == 0 {
		setError(w, "Не передан обязательный параметр id")
		return
	}

	res, err := db.SelectOne(m.db, in.ID)
	if err != nil {
		log.Println(err)
		setError(w, InternalServiceError)
		return
	}
	out = append(out, res)
	setResult(w, out)
	return
}

// Create - метод создания машины
func (m Typ) Create(w http.ResponseWriter, r *http.Request) {
	var (
		in  InsertRequest
		car cars.Car
	)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if checkHttpPost(w, r) != nil {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if in.Company == "" {
		setError(w, "Не передан обязательный параметр company")
		return
	}
	car.Company = in.Company

	if in.Model == "" {
		setError(w, "Не передан обязательный параметр model")
		return
	}
	car.Model = in.Model

	if in.Capacity == 0 {
		setError(w, "Не передан обязательный параметр capacity")
		return
	}
	car.Capacity = in.Capacity

	if err := db.Insert(m.db, car); err != nil {
		log.Println(err)
		setError(w, InternalServiceError)
		return
	}

	setResult(w, nil)
	return
}

// Update - метод обновления автомобиля
func (m Typ) Update(w http.ResponseWriter, r *http.Request) {
	var (
		in  UpdateRequest
		car cars.Car
	)
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if checkHttpPost(w, r) != nil {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if in.ID == 0 {
		setError(w, "Не передан обязательный параметр id")
		return
	}
	car.ID = in.ID

	if in.Company == "" {
		setError(w, "Не передан обязательный параметр company")
		return
	}
	car.Company = in.Company

	if in.Model == "" {
		setError(w, "Не передан обязательный параметр model")
		return
	}
	car.Model = in.Model

	if in.Capacity == 0 {
		setError(w, "Не передан обязательный параметр capacity")
		return
	}
	car.Capacity = in.Capacity

	if err := db.Update(m.db, car); err != nil {
		log.Println(err)
		setError(w, InternalServiceError)
		return
	}

	setResult(w, nil)
	return
}

// Delete - метод обновления автомобиля
func (m Typ) Delete(w http.ResponseWriter, r *http.Request) {
	var in IDRequest

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if checkHttpPost(w, r) != nil {
		return
	}

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, BadRequest)
		return
	}

	if in.ID == 0 {
		setError(w, "Не передан обязательный параметр id")
		return
	}

	if err := db.Delete(m.db, in.ID); err != nil {
		log.Println(err)
		setError(w, InternalServiceError)
		return
	}

	setResult(w, nil)
	return
}
