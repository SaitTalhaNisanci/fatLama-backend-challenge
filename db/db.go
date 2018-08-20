package db

import (
	"database/sql"

	"errors"

	"io/ioutil"

	"github.com/fatLama-backend-challenge/model"
	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

// Items is a database for items.
type Items struct {
	db        *sql.DB
	queryChan chan queryRequest
}

// queryResult is the result of a search query.
type queryResult struct {
	items []*model.Item
	err   error
}

// result is a helper function to return items and error.
func (q *queryResult) result() ([]*model.Item, error) {
	return q.items, q.err
}

// queryRequest is used for doing queries.
// searchQuery is used to do the query and the result of the query
// is sent to queryResultChan.
type queryRequest struct {
	searchQuery     string
	queryResultChan chan *queryResult
}

// NewItems returns an Items database from the given dataSourcePath.
// It returns an error if the path doesnt exist or there was a problem with loading the database.
func NewItems(dataSourcePath string) (*Items, error) {
	items := &Items{
		queryChan: make(chan queryRequest, 1000),
	}
	db, err := items.initDB(dataSourcePath)
	if err != nil {
		return nil, err
	}
	items.db = db
	go items.process()
	return items, nil
}

// process handles queries.
func (i *Items) process() {
	for {
		select {
		case queryReq := <-i.queryChan:
			i.doQuery(queryReq)
		}
	}
}

// initDB initializes the database from the given dataSourcePath.
func (i *Items) initDB(dataSourcePath string) (*sql.DB, error) {
	if _, err := ioutil.ReadFile(dataSourcePath); err != nil {
		return nil, err
	}
	db, err := sql.Open(driverName, dataSourcePath)
	return db, err
}

// LoadItemsBySearchTerm loads items that have the given search term in its item name.
func (i *Items) LoadItemsBySearchTerm(searchTerm string) ([]*model.Item, error) {
	if i.db == nil {
		return nil, errors.New("database is nil")
	}
	searchQuery := generateQuery(searchTerm)
	queryRes := make(chan *queryResult, 0)
	queryReq := queryRequest{
		searchQuery:     searchQuery,
		queryResultChan: queryRes,
	}
	i.queryChan <- queryReq
	return (<-queryRes).result()
}

// loadItemsInternal iterates through the rows to construct item slice.
func (i *Items) loadItemsInternal(rows *sql.Rows) ([]*model.Item, error) {
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

// doQuery does a query with searchQuery and puts the results to queryResultChan.
func (i *Items) doQuery(req queryRequest) {
	rows, err := i.db.Query(req.searchQuery)
	if err != nil {
		req.queryResultChan <- &queryResult{nil, err}
		return
	}
	go func() {
		items, err := i.loadItemsInternal(rows)
		req.queryResultChan <- &queryResult{items, err}
	}()
}
