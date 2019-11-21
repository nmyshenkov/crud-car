package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"private/crud-car/cars"
	"private/crud-car/db"
)

func (t Typ) List(w http.ResponseWriter, r *http.Request) {
	var in cars.Car

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	res, err := db.Select(t.db)
	if err != nil {
		log.Println(err)
		setError(w, "Internal Service Error")
		return
	}
	setResult(w, res)
}

func (t Typ) Get(w http.ResponseWriter, r *http.Request) {
	var (
		in  cars.Car
		out []cars.Car
	)

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if in.ID == 0 {
		setError(w, "Не передан обязательный параметр id")
		return
	}

	res, err := db.SelectOne(t.db, in.ID)
	if err != nil {
		log.Println(err)
		setError(w, "Internal Service Error")
		return
	}
	out = append(out, res)
	setResult(w, out)
}

func (t Typ) Create(w http.ResponseWriter, r *http.Request) {
	var in cars.Car

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if in.Company == "" {
		setError(w, "Не передан обязательный параметр company")
		return
	}

	if in.Model == "" {
		setError(w, "Не передан обязательный параметр model")
		return
	}

	if in.Capacity == 0 {
		setError(w, "Не передан обязательный параметр capacity")
	}
	if err := db.Insert(t.db, in); err != nil {
		log.Println(err)
		setError(w, "Internal Service Error")
		return
	}

	setResult(w, nil)
}

func (t Typ) Update(w http.ResponseWriter, r *http.Request) {
	var in cars.Car

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if in.ID == 0 {
		setError(w, "Не передан обязательный параметр id")
		return
	}

	setResult(w, nil)
}

func (t Typ) Delete(w http.ResponseWriter, r *http.Request) {

	var in cars.Car

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if err := json.Unmarshal(reqBody, &in); err != nil {
		log.Println(err)
		setError(w, "Bad request")
		return
	}

	if in.ID == 0 {
		setError(w, "Не передан обязательный параметр id")
		return
	}

	setResult(w, nil)
}
