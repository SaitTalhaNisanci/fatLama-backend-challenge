package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"github.com/fatLama-backend-challenge/db"
	"github.com/fatLama-backend-challenge/model"
	"github.com/stretchr/testify/assert"
)

const databasePath = "../fatlama.sqlite3"

func TestSearchHandler(t *testing.T) {
	req, _ := http.NewRequest("GET", "/search?searchTerm=camera&lat=54.948&lng=0.172943", nil)
	res := httptest.NewRecorder()
	itemsDB, err := db.NewItems(databasePath)
	assert.NoError(t, err)
	SearchHandler(res, req, itemsDB)
	assert.Equal(t, res.Code, http.StatusOK)
	items := make([]*model.Item, 0)
	json.Unmarshal(res.Body.Bytes(), &items)
	if len(items) > defaultPageSize {
		t.Errorf("/search returned %d results expected %d", len(items), defaultPageSize)
	}
}

func TestSearchHandlerWithInvalidLat(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/search?searchTerm=camera&lat=invalid&lng=0.172943", nil)
	res := httptest.NewRecorder()
	itemsDB, err := db.NewItems(databasePath)
	assert.NoError(t, err)
	SearchHandler(res, req, itemsDB)
	assert.Equal(t, res.Code, http.StatusBadRequest)
}

func TestSearchHandlerWithInvalidLng(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/search?searchTerm=camera&lat=12.11&lng=invalid", nil)
	res := httptest.NewRecorder()
	itemsDB, err := db.NewItems(databasePath)
	assert.NoError(t, err)
	SearchHandler(res, req, itemsDB)
	assert.Equal(t, res.Code, http.StatusBadRequest)
}

func TestSearchHandlerWithNoContent(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/search?searchTerm=invalidSearchTerm&lat=12.11&lng=11.21", nil)
	res := httptest.NewRecorder()
	itemsDB, err := db.NewItems(databasePath)
	assert.NoError(t, err)
	SearchHandler(res, req, itemsDB)
	assert.Equal(t, res.Code, http.StatusNoContent)
}
