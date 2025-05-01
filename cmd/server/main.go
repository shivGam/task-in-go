package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/shivGam/task-in-go/internal/config"
	"github.com/shivGam/task-in-go/internal/db"
	"github.com/shivGam/task-in-go/internal/handlers"
)

func main() {

	config.LoadEnv()
	db.Connect()
	defer db.Close()

	r:=mux.NewRouter()
	r.HandleFunc("/health",handlers.HealthCheckup).Methods("GET")
	r.HandleFunc("/query",handlers.RunQuery).Methods("POST")

	port := os.Getenv("SERVER_PORT")

	log.Println("🚀 Server running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}