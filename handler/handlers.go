package handler

import (
	"net/http"

	"encoding/json"

	"github.com/fatLama-backend-challenge/db"
	"github.com/fatLama-backend-challenge/search"
)

// defaultPageSize is at most how many items will be returned as a response.
// If there are less than 20 items matching the query the result will be less than 20.
const defaultPageSize = 20

// noItemFound is returned as an error message when no item is found for search.
const noItemFound = "No item is found for the query!"

// SearchHandler handles the search and returns the top matching 20 items.
// Search currently consists of a search text, longitude and latitude.
// If no content is found 'http.StatusNoContent' status is returned.
func SearchHandler(w http.ResponseWriter, r *http.Request, itemsDB *db.Items) {
	vars := r.URL.Query()
	searchParams, err := parseQuery(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultItems, err := search.DoSearch(searchParams, itemsDB, defaultPageSize)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// If there is no content return http.StatusNoContent response.
	if len(resultItems) == 0 {
		http.Error(w, noItemFound, http.StatusNoContent)
		return
	}
	json.NewEncoder(w).Encode(resultItems)
}
