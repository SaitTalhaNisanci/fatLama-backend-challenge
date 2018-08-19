package main

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"sync"

	"math/rand"
	"strconv"

	"time"

	"github.com/fatLama-backend-challenge/db"
	"github.com/stretchr/testify/assert"
)

// searchTerms is used to generate a search query for testing. Note that all of them have at least one entry in the database.
var searchTerms = []string{"camera", "playstation", "car", "camera%20canon", "automatic%20car"}

func TestServerRouter(t *testing.T) {
	wg := new(sync.WaitGroup)
	itemsDB, _ := db.NewItems(dataSourcePath)
	r := initializeRouter(itemsDB)

	reqAmount := 2000
	wg.Add(reqAmount)
	for i := 0; i < reqAmount; i++ {
		go func() {
			query := generateValidSearchQuery()
			req, _ := http.NewRequest(http.MethodGet,
				query, nil)
			res := httptest.NewRecorder()
			r.ServeHTTP(res, req)
			assert.Equal(t, res.Code, http.StatusOK)
			assert.NotEmpty(t, res.Body)
			wg.Done()
		}()
	}
	wg.Wait()
}

func generateValidSearchQuery() string {
	rand.Seed(time.Now().UTC().UnixNano())
	rndIndex := rand.Int31n(int32(len(searchTerms)))
	searchTerm := searchTerms[rndIndex]
	lat := floattostr(60 * rand.Float64())
	lng := floattostr(60 * rand.Float64())
	return "/search?searchTerm=" + searchTerm + "&lat=" + lat + "&lng=" + lng
}

func floattostr(fv float64) string {
	return strconv.FormatFloat(fv, 'f', 2, 64)
}
