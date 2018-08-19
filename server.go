package main

import (
	"log"
	"net/http"

	"github.com/fatLama-backend-challenge/db"
	"github.com/fatLama-backend-challenge/handler"
	"github.com/gorilla/mux"
)

const (

	// dataSourcePath is where the database is.
	dataSourcePath = "./fatlama.sqlite3"

	// searchEndpoint is the entry point of search
	searchEndpoint = "/search"
)

func main() {
	// initialize database
	itemsDB, err := db.NewItems(dataSourcePath)
	if err != nil {
		log.Fatal("Database could not start: ", err)
	}
	r := mux.NewRouter()
	// add search GET endpoint to the router
	r.HandleFunc(searchEndpoint, func(w http.ResponseWriter, r *http.Request) {
		handler.SearchHandler(w, r, itemsDB)
	}).Methods(http.MethodGet)
	// start the server
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("Server could not start: ", err)
	}
}
