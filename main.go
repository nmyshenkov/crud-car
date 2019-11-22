package main

import (
	"context"
	"crud-car/api"
	"crud-car/db"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	// Setting up signal capturing
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	dbconn, err := db.InitDB(db.DSN)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		_ = dbconn.Close()
	}()

	m := api.Init(dbconn)

	router := http.NewServeMux()
	router.HandleFunc("/get", m.Get)
	router.HandleFunc("/create", m.Create)
	router.HandleFunc("/update", m.Update)
	router.HandleFunc("/list", m.List)
	router.HandleFunc("/delete", m.Delete)

	addr := ":8080"

	log.Println("Starting server on", addr)

	server := &http.Server{Addr: addr, Handler: router}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on %s: %v\n", addr, err)
		}
	}()

	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Shotdown error: %v\n", err)
	}

	log.Println("Server stopped")
}
