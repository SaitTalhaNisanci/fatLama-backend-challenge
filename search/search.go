package search

import (
	"sort"

	"github.com/fatLama-backend-challenge/model"
)

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
