package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"private/crud-car/api"
)

func InitDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	// db, err := sql.Open("mysql", )
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	var dsn = "username:password@tcp(127.0.0.1:3306)/vehicals"

	db, err := InitDB(dsn)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	m := api.Init(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/get", m.Get)
	mux.HandleFunc("/create", m.Create)
	mux.HandleFunc("/update", m.Update)
	mux.HandleFunc("/list", m.List)
	mux.HandleFunc("/delete", m.Delete)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
