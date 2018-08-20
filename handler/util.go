package handler

import (
	"net/url"

	"errors"

	"github.com/fatLama-backend-challenge/search"
)

const (
	// searchTerm is the key for searchTerm in GET query.
	searchTerm = "searchTerm"

	// lat is the key for latitude in GET query.
	lat = "lat"

	// lng is the key for longitude in GET query.
	lng = "lng"

	// emptyString is used to check if url values exist
	emptyString = ""
)

// parseQuery returns SearchParams from the given url values.
// It returns an error if values cannot be casted to SearchParams fields properly.
func parseQuery(values url.Values) (*search.Params, error) {
	if values.Get(searchTerm) == emptyString {
		return nil, errors.New("searchTerm should be set for search")
	}
	if values.Get(lat) == emptyString {
		return nil, errors.New("lat should be set for search")
	}
	if values.Get(lng) == emptyString {
		return nil, errors.New("lng should be set for search")
	}
	return search.NewSearchParams(values[searchTerm][0], values[lat][0], values[lng][0])
}
