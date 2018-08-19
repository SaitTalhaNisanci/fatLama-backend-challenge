package db

import (
	"database/sql"

	"errors"

	"io/ioutil"

	"github.com/fatLama-backend-challenge/model"
	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

type Items struct {
	db *sql.DB
}

// NewItems returns an Items database from the given dataSourcePath.
// It returns an error if the path doesnt exist or there was a problem with loading the database.
func NewItems(dataSourcePath string) (*Items, error) {
	items := &Items{}
	db, err := items.InitDB(dataSourcePath)
	if err != nil {
		return nil, err
	}
	items.db = db
	return items, nil
}

// InitDB initializes the database from the given dataSourcePath.
func (i *Items) InitDB(dataSourcePath string) (*sql.DB, error) {
	if _, err := ioutil.ReadFile(dataSourcePath); err != nil {
		return nil, err
	}
	db, err := sql.Open(driverName, dataSourcePath)
	return db, err
}

// LoadItems loads all the items from the database.
func (i *Items) LoadItems() ([]*model.Item, error) {
	if i.db == nil {
		return nil, errors.New("database is nil")
	}
	rows, err := i.db.Query("SELECT * FROM items")
	if err != nil {
		return nil, err
	}
	items := make([]*model.Item, 0)
	for rows.Next() {
		item := new(model.Item)
		err := rows.Scan(&item.Name, &item.Lat, &item.Lng, &item.Url, &item.ImageUrls)
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}
