package handler

import (
	"net/http"

	"encoding/json"

	"github.com/fatLama-backend-challenge/db"
	"github.com/fatLama-backend-challenge/model"
	"github.com/fatLama-backend-challenge/search"
)

// defaultPageSize is at most how many items will be returned as a response.
// If there are less than 20 items matching the query the result will be less than 20.
const defaultPageSize = 20
const noItemFound = "No item is found for the query!"

// SearchHandler handles the search and returns the top matching 20 items.
// Search currently consists of a search text, longitude and latitude.
func SearchHandler(w http.ResponseWriter, r *http.Request, itemsDB *db.Items) {
	vars := r.URL.Query()
	searchParams, err := parseQuery(vars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resultItems, err := doSearch(searchParams, itemsDB)
	// If there is no content return http.StatusNoContent response.
	if len(resultItems) == 0 {
		http.Error(w, noItemFound, http.StatusNoContent)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resultItems = search.SortByRelevance(resultItems, searchParams.Lat(), searchParams.Lng(), defaultPageSize)
	json.NewEncoder(w).Encode(resultItems)
}

func doSearch(searchParams *SearchParams, itemsDB *db.Items) ([]*model.Item, error) {
	return itemsDB.LoadItemsBySearchTerm(searchParams.SearchTerm())
}
