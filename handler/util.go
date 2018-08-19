package handler

import (
	"net/url"
	"strconv"
)

const (
	// bitSize is the size of bits for converted floats after the decimal point.
	bitSize = 8

	// searchTerm is the key for searchTerm in GET query.
	searchTerm = "searchTerm"

	// lat is the key for latitude in GET query.
	lat = "lat"

	// lng is the key for longitude in GET query.
	lng = "lng"
)

// SearchParams holds GET request parameters.
type SearchParams struct {
	searchTerm string
	lat        float64
	lng        float64
}

func (s *SearchParams) SearchTerm() string {
	return s.searchTerm
}

func (s *SearchParams) Lat() float64 {
	return s.lat
}

func (s *SearchParams) Lng() float64 {
	return s.lng
}

func newSearchParams(searchTerm string, latStr string, lngStr string) (*SearchParams, error) {
	lat, err := strconv.ParseFloat(latStr, bitSize)
	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(lngStr, bitSize)
	if err != nil {
		return nil, err
	}
	searchParams := &SearchParams{
		searchTerm: searchTerm,
		lat:        lat,
		lng:        lng,
	}
	return searchParams, nil
}

// parseQuery returns SearchParams from the given urlValues.
// It returns an error if values cannot be casted to SearchParams fields properly.
func parseQuery(values url.Values) (*SearchParams, error) {
	return newSearchParams(values[searchTerm][0], values[lat][0], values[lng][0])
}
