package main

import (
	"log"
	"net/http"
	"os"
	"sourceanddestination/routes"
	db "sourceanddestination/utils"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var dbSource *gorm.DB
var dbDestination *gorm.DB

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitDB()
	router := mux.NewRouter()
	routes.Init(router, dbSource, dbDestination)
	log.Println("Server running on http://localhost:" + os.Getenv("PORT"))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}
