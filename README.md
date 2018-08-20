# FatLama Backend Challenge 

[![GoDoc](https://godoc.org/github.com/SaitTalhaNisanci/fatLama-backend-challenge?status.svg)](https://godoc.org/github.com/SaitTalhaNisanci/fatLama-backend-challenge)
[![Go Report Card](https://goreportcard.com/badge/github.com/SaitTalhaNisanci/fatLama-backend-challenge)](https://goreportcard.com/report/github.com/SaitTalhaNisanci/fatLama-backend-challenge)

## Installation

Make sure you have **go 1.8+** installed, if you dont have go please [install it here](https://golang.org/doc/install).

Make sure your **GOROOT** and **GOPATH** are set correctly.

Run the following to get the code:
```
go get -u github.com/SaitTalhaNisanci/fatLama-backend-challenge
```



## Technology

- Go
- [gorilla mux](https://github.com/gorilla/mux) router and dispatcher
- [testify](https://github.com/stretchr/testify) testing
- [sqlite](https://www.sqlite.org/index.html) database engine

[CodeReviewComments](https://github.com/golang/go/wiki/CodeReviewComments) is followed 
during the development.

## Search

`/search?searchTerm=camera&lat=51.948&lng=0.172943`

Search engine splits the searchTerm to its words, and for each item 
in the database it checks if there is any matched word. Even if there is
one match with an item's name it will be loaded. For example If
`searchTerm` is **canon camera**, items that have **canon** or **camera**
or both will be loaded. When calculating a score of an item, the number 
of matched words with its name and searchTerm are linearly combined with 
the distance.

```
score = textMatchCoeff * #matchedWords - distanceCoeff*distance(item, searchLocation)
```

Note that distance is subtracted as the further it is the less relevant it is.

These coefficients can be set better with a real data and some observations. 

Default page size is 20, which means at most 20 items will be returned.

If there is no content for a search query, ``http.StatusNoContent`` is returned.

If search parameters are not valid,  ``http.StatusBadRequest`` is returned.

Currently if the number of requests per second increase the latency will increase
but as far as tested in `server_test.go` the server wont fail to serve. Note that it is 
only a smoke test not a soak test.

If an item has duplicate words in its name they will be counted as one. So
an item name `camera camera camera` will be same as `camera`.


## Design Details


The entry point for the server is `server.go`. Router and database are
initialized there. The router currently has only one endpoint which is `/search`.

There are 4 packages:

### db

db contains database methods, `process` method processes incoming 
search queries. A query channel sends search queries to this method to
be processed.

### handler

handler contains handler functions. Currently there is only one handler, search handler.
Search parameter validation etc are done in this package.


### model

model contains database models. Currently there is only one model which is **Item**.
Item contains `Name, Lat, Lng, Url, ImageUrls`.

### search

search package communicates with db package to load items and sort them by a
score. Score is calculated based on linear combination of distance and matched words.


Different functionalities are separated into difference packages for better
readability and testing.

Currently each request hits the database. For scalability and
better latency in memory cache such as `redis, hazelcast can be used.


## Test

The code is tested extensively,

To run tests:
```
go test ./...
```

To run tests with race:

```
go test -race ./...
```

Travis can be added for CI to the repository.

Also note that currently the same database is used for testing and production
but they should be separated.

