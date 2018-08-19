package search

import (
	"sort"

	"strings"

	"strconv"

	"github.com/fatLama-backend-challenge/db"
	"github.com/fatLama-backend-challenge/model"
)

const (

	// bitSize is the size of bits for converted floats after the decimal point.
	bitSize = 8

	distanceCoeff  = 2
	textMatchCoeff = 10
)

// Params holds GET request parameters.
type Params struct {
	searchTerm string
	lat        float64
	lng        float64
}

// SearchTerm returns search term for query.
func (s *Params) SearchTerm() string {
	return s.searchTerm
}

// Lat returns latitude for query.
func (s *Params) Lat() float64 {
	return s.lat
}

// Lng return longitude for query.
func (s *Params) Lng() float64 {
	return s.lng
}

// NewSearchParams returns search params with given arguments. It returns an error
// if given lng or lat cannot be converted to float64 type.
func NewSearchParams(searchTerm string, latStr string, lngStr string) (*Params, error) {
	lat, err := strconv.ParseFloat(latStr, bitSize)
	if err != nil {
		return nil, err
	}

	lng, err := strconv.ParseFloat(lngStr, bitSize)
	if err != nil {
		return nil, err
	}
	searchParams := &Params{
		searchTerm: searchTerm,
		lat:        lat,
		lng:        lng,
	}
	return searchParams, nil
}

func DoSearch(searchParams *Params, itemsDB *db.Items, pageSize int) ([]*model.Item, error) {
	loadedItems, err := itemsDB.LoadItemsBySearchTerm(searchParams.SearchTerm())
	if err != nil {
		return nil, err
	}
	return SortByRelevance(loadedItems, searchParams, pageSize), nil
}

func SortByRelevance(items []*model.Item, searchParams *Params, pageSize int) []*model.Item {
	sort.Slice(items, func(i, j int) bool {
		return distance(items[i].Lat, items[i].Lng, searchParams.Lat(), searchParams.Lng()) <
			distance(items[j].Lat, items[j].Lng, searchParams.Lat(), searchParams.Lng())
	})
	// update pageSize if we dont have enough items
	if len(items) < pageSize {
		pageSize = len(items)
	}
	return items[:pageSize]
}

func giveScore(item *model.Item, searchWords []string) int {
	wordMatchCount := 0
	for _, word := range searchWords {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(word)) {
			wordMatchCount++
		}
	}
	return 0
}
