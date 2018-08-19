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

func (s *Params) SearchTerm() string {
	return s.searchTerm
}

func (s *Params) Lat() float64 {
	return s.lat
}

func (s *Params) Lng() float64 {
	return s.lng
}

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

func DoSearch(searchTerm string, itemsDB *db.Items) ([]*model.Item, error) {
	return itemsDB.LoadItemsBySearchTerm(searchTerm)
}

func SortByRelevance(items []*model.Item, lat float64, lng float64, pageSize int) []*model.Item {
	sort.Slice(items, func(i, j int) bool {
		return distance(items[i].Lat, items[i].Lng, lat, lng) <
			distance(items[j].Lat, items[j].Lng, lat, lng)
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
