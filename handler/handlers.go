// package handler contains endpoint handlers.
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

// SearchHandler handles the search and returns the top matching 20 items from the database.
// Search currently consists of a search text, longitude and latitude.
// All of these parameters should be given for a search query.
// If get parameters cannot be parsed, 'http.StatusBadRequest' status is returned.
// If no content is found 'http.StatusNoContent' status is returned.
// The result items will be sent in JSON format.
func SearchHandler(w http.ResponseWriter, r *http.Request, itemsDB *db.Items) {
	searchParams, err := parseQuery(r.URL.Query())
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
	// encode and send result items as JSON
	json.NewEncoder(w).Encode(resultItems)
}
