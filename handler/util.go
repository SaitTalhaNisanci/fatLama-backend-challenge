package handler

import (
	"net/url"

	"github.com/fatLama-backend-challenge/search"
)

const (
	// searchTerm is the key for searchTerm in GET query.
	searchTerm = "searchTerm"

	// lat is the key for latitude in GET query.
	lat = "lat"

	// lng is the key for longitude in GET query.
	lng = "lng"
)

// parseQuery returns SearchParams from the given urlValues.
// It returns an error if values cannot be casted to SearchParams fields properly.
func parseQuery(values url.Values) (*search.Params, error) {
	return search.NewSearchParams(values[searchTerm][0], values[lat][0], values[lng][0])
}
