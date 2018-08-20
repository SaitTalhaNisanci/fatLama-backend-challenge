// package search contains search methods such as scoring.
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

	// separator is used to split search terms.
	separator = " "

	// distanceCoeff is used when calculating the score of an item with respect to a search query.
	distanceCoeff = 1

	// textMatchCoeff is used when calculating the score of an item with respect to a search query.
	// Matching a word means matching exactly here. These coefficients can be interpreted as follows:
	//  Assume there are 2 items in database that match our search query.
	//  First item's name matches two words from the query, second item's name matches one word from the query.
	//  Second item will be displayed before the first item if it is at least 20.000 meters more closer to the query location
	//  compared to the first item.
	// Note that with the complete database and what we want we can set a better coefficient here based on what we want
	// to display to users.
	textMatchCoeff = distanceCoeff * 20000
)

// Params holds search GET request parameters.
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

// DoSearch loads items from database with respect to the search term. It then sorts the items based on their score.
// The score is calculated based on matched words and distance.
// It returns at most pageSize items. It might return less or empty if there is less content or no content.
func DoSearch(searchParams *Params, itemsDB *db.Items, pageSize int) ([]*model.Item, error) {
	loadedItems, err := itemsDB.LoadItemsBySearchTerm(searchParams.SearchTerm())
	if err != nil {
		return nil, err
	}
	return SortByScore(loadedItems, searchParams, pageSize), nil
}

// SortByScore sorts items based on their score in ascending order with respect to given searchParams.
// It returns at most pageSize items. It might return less or empty if there is less content or no content.
func SortByScore(items []*model.Item, searchParams *Params, pageSize int) []*model.Item {
	sort.Slice(items, func(i, j int) bool {
		return calculateScore(items[i], searchParams) >
			calculateScore(items[j], searchParams)
	})
	// update pageSize if we dont have enough items
	if len(items) < pageSize {
		pageSize = len(items)
	}
	return items[:pageSize]
}

// calculateScore calculates score for a given item and given searchParams.
// To calculate score number of exactly matching words are found from the given search term.
// For example if an items name is "high quality cheap camera" and if our search term is "quality camera", the matching
// words count will be 2, "quality" and "camera". Note that duplicates dont increase the matching count, which means
// if an item's name is "camera camera camera" it wont have a higher score compared to an item that has a name "camera".
// Second score comes from the geological distance. The final score is composed of linear combination of these two scores.
// Note that distance is subtracted from the score as the further an item is the less relevant it is.
func calculateScore(item *model.Item, searchParams *Params) int {
	wordMatchCount := 0
	searchWords := strings.Split(searchParams.SearchTerm(), separator)
	for _, word := range searchWords {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(word)) {
			wordMatchCount++
		}
	}
	distInMeters := distance(item.Lat, item.Lng, searchParams.Lat(), searchParams.Lng())
	return textMatchCoeff*wordMatchCount - distanceCoeff*distInMeters
}
