package model

import "fmt"

// Item is a model in database.
type Item struct {
	// Name is items name, it can contains multiple words.
	Name string
	// Lat is latitude of the item.
	Lat float64
	// Lng is longitude of the item.
	Lng float64
	// Url is items url. This is the main image to display for item.
	Url string
	// ImageUrls is a list of image urls in string format.
	ImageUrls string
}

// NewItem returns an Item with the given parameters.
func NewItem(name string, lat float64, lng float64, url string, imageUrls string) *Item {
	return &Item{
		Name:      name,
		Lat:       lat,
		Lng:       lng,
		Url:       url,
		ImageUrls: imageUrls,
	}
}

func (i *Item) String() string {
	return fmt.Sprintf("Name: %s, Lat: %f, Lng: %f, Url: %s, ImageUrls: %s", i.Name, i.Lat,
		i.Lng, i.Url, i.ImageUrls)
}
