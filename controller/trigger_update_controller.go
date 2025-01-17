package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sourceanddestination/repository"
)

func TriggerBackgroundUpdate(w http.ResponseWriter, r *http.Request) {
	if dbSource == nil {
		log.Println("Database connection not initialized")
		http.Error(w, "Database connection not initialized", http.StatusInternalServerError)
		return
	}
	// jalankan goroutine
	go repository.ProcessBackgroundUpdate(dbSource, dbDestination)

	response := map[string]interface{}{
		"status":  "OK",
		"code":    http.StatusOK,
		"message": "Background update started",
	}

	json.NewEncoder(w).Encode(response)
}
