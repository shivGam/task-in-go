package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/shivGam/task-in-go/internal/db"
)

type QueryData struct {
	Query string `json:"query"`
}

func RunQuery(w http.ResponseWriter,r *http.Request) {
	var req QueryData
	w.Header().Set("Content-Type","applicaton/json")

	if err:= json.NewDecoder(r.Body).Decode(&req); err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error":"Invalid request"})
		return
	}
	
	rows,err := db.DB.Query(context.Background(),req.Query)
	if err!=nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error":err.Error()})
		return
	}
	defer rows.Close()

	var result []map[string]interface{}

	for rows.Next() {
		value,err := rows.Values()
		if err!=nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error":err.Error()})
			return
		}

		rowMap := make(map[string]interface{})
		field := rows.FieldDescriptions()

		for i, val := range value {
            rowMap[string(field[i].Name)] = val
        }
        result = append(result, rowMap)
	}
	json.NewEncoder(w).Encode(result)
}