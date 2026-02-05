package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type TimeResponse struct {
	Time   string `json:"time"`
	Status string `json:"status"`
}

func TimeHandler(w http.ResponseWriter, r *http.Request) {

	response := TimeResponse{
		Time:   time.Now().Format(time.RFC3339),
		Status: "running",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
