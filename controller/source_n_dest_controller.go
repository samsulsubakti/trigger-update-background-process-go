package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sourceanddestination/models"
	"sourceanddestination/repository"
)

func IndexSources(w http.ResponseWriter, r *http.Request) {
	var sources []models.SourceProduct
	response := repository.GetList(dbSource, sources, w)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func IndexDestinations(w http.ResponseWriter, r *http.Request) {
	var data []models.DestinationProduct
	response := repository.GetList(dbDestination, data, w)

	err := json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
