package routes

import (
	"encoding/json"
	"net/http"
	"sourceanddestination/controller"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var _prefix string = "/api/v1"

func setContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") // Set Content-Type for all responses
		next.ServeHTTP(w, r)
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	messages := map[string]interface{}{
		"content":                             "Hello, DB_SOURCE and DB_DESTINATION!, Set Your .ENV",
		"author":                              "SAMSUL SUBAKTI",
		"_prefix":                             _prefix,
		"_api_list_source":                    _prefix + "/source",
		"_api_list_destination":               _prefix + "/destination",
		"_api_update_destination_from_source": _prefix + "/trigger-destination-update",
	}

	json.NewEncoder(w).Encode(messages)
}

func Init(router *mux.Router, dbSource *gorm.DB, dbDestination *gorm.DB) {
	router.Use(setContentTypeJson)
	router.HandleFunc("/", home).Methods("GET")
	apiRouter := router.PathPrefix(_prefix).Subrouter()
	apiRouter.HandleFunc("/source", controller.IndexSources).Methods("GET")
	apiRouter.HandleFunc("/destination", controller.IndexDestinations).Methods("GET")
	apiRouter.HandleFunc("/trigger-destination-update", controller.TriggerBackgroundUpdate).Methods("GET")
}
