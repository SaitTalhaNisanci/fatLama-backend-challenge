// package db contains database methods.
package db

import (
	"database/sql"

	"errors"

	"io/ioutil"

	"github.com/fatLama-backend-challenge/model"

	// for sqlite3
	_ "github.com/mattn/go-sqlite3"
)

const driverName = "sqlite3"

// Items is a sqlite3 database for items.
type Items struct {
	db        *sql.DB           //database
	queryChan chan queryRequest //queryChan is used to handle search queries
}

// queryResult is the result of a search query that contains the returned elements and an error.
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
	// start process method in a new subroutine which will handle incoming queries.
	go items.process()
	return items, nil
}

// process handles queries.
func (i *Items) process() {
	// TODO:: close the channel when the server goes down.
	for {
		select {
		case queryReq := <-i.queryChan:
			i.doQuery(queryReq)
		}
	}
}

// initDB initializes the database from the given dataSourcePath.
// sqlite3 engine is used for database.
// if the given dataSourcePath doesn't exist an error is returned.
func (i *Items) initDB(dataSourcePath string) (*sql.DB, error) {
	if _, err := ioutil.ReadFile(dataSourcePath); err != nil {
		return nil, err
	}
	db, err := sql.Open(driverName, dataSourcePath)
	return db, err
}

// LoadItemsBySearchTerm loads items that have the given search term in its item name.
// If the given search term contains multiple words an item that has any of the words in its name
// will be loaded.
func (i *Items) LoadItemsBySearchTerm(searchTerm string) ([]*model.Item, error) {
	if i.db == nil {
		return nil, errors.New("database is nil")
	}
	// TODO:: use in memory cache such as redis, hazelcast to decrease loading items latency.
	// generate query for database search.
	searchQuery := generateQuery(searchTerm)

	// queryResChan will get the response for search query when it is ready.
	queryResChan := make(chan *queryResult, 0)
	queryReq := queryRequest{
		searchQuery:     searchQuery,
		queryResultChan: queryResChan,
	}
	// send query request to itemsDB queryChan to search the database.
	i.queryChan <- queryReq
	return (<-queryResChan).result()
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

	// process the rows in a new subroutine as it is independent from querying the database.
	go func() {
		items, err := i.loadItemsInternal(rows)
		req.queryResultChan <- &queryResult{items, err}
	}()
}
