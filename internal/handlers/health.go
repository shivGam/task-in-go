package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shivGam/task-in-go/internal/db"
)

func HealthCheckup(w http.ResponseWriter,r *http.Request){
	err := db.DB.Ping(context.Background())
	w.Header().Set("Content-Type" , "application/json")
	if err!=nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string {"status":"error","message":err.Error()})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}