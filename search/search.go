package search

import (
	"sort"

	"github.com/fatLama-backend-challenge/model"
)

func SortByRelevance(items []*model.Item, lat float64, lng float64, pageSize int64) []*model.Item {
	sort.Slice(items, func(i, j int) bool {
		return distance(items[i].Lat, items[i].Lng, lat, lng) <
			distance(items[j].Lat, items[j].Lng, lat, lng)
	})
	return items[:pageSize]
}
