package db

import (
	"database/sql"

	"errors"

	"io/ioutil"

	"github.com/fatLama-backend-challenge/model"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

const driverName = "sqlite3"

// InitDB initializes the database from the given dataSourcePath.
func InitDB(dataSourcePath string) error {
	var err error
	if _, err = ioutil.ReadFile(dataSourcePath); err != nil {
		return err
	}
	db, err = sql.Open(driverName, dataSourcePath)
	return err
}

// LoadItems loads all the items from the database.
func LoadItems() ([]*model.Item, error) {
	if db == nil {
		return nil, errors.New("database is nil")
	}
	rows, err := db.Query("SELECT * FROM items")
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
