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

// searchParams holds GET request parameters.
type searchParams struct {
	searchTerm string
	lat        float64
	lng        float64
}

func newSearchParams(searchTerm string, latStr string, lngStr string) (*searchParams, error) {
	lat, err := strconv.ParseFloat(latStr, bitSize)
	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(lngStr, bitSize)
	if err != nil {
		return nil, err
	}
	searchParams := &searchParams{
		searchTerm: searchTerm,
		lat:        lat,
		lng:        lng,
	}
	return searchParams, nil
}

// parseQuery returns searchParams from the given urlValues.
// It returns an error if values cannot be casted to searchParams fields properly.
func parseQuery(values url.Values) (*searchParams, error) {
	return newSearchParams(values[searchTerm][0], values[lat][0], values[lng][0])
}
