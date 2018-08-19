package handler

import (
	"net/http"

	"github.com/fatLama-backend-challenge/db"
	"github.com/fatLama-backend-challenge/model"
)

// SearchHandler handles the search and returns the top matching 20 items.
// Search currently consists of a search text, longitude and latitude.
func SearchHandler(w http.ResponseWriter, r *http.Request, itemsDB *db.Items) {
	vars := r.URL.Query()
	_, err := parseQuery(vars)
	if err != nil {
		//TODO:: return HTTP 400
	}
	//resultItems := doSearch(searchParams, itemsDB)

}

func writeJsonResponse(w http.ResponseWriter, bytes []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Write(bytes)
}

func doSearch(searchParams *searchParams, itemsDB *db.Items) []*model.Item {
	return nil
}
