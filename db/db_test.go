package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const databasePath = "../fatlama.sqlite3"

func TestInitDBWithInvalidPath(t *testing.T) {
	invalidPath := "./invalid.sqlite3"
	itemsDB := &Items{}
	if _, err := itemsDB.InitDB(invalidPath); err == nil {
		t.Errorf("Database should not start with invalidPath %s", invalidPath)
	}
}

func TestInitDB(t *testing.T) {
	validPath := databasePath
	itemsDB := &Items{}
	if _, err := itemsDB.InitDB(validPath); err != nil {
		t.Errorf("InitDB should not return an error for valid path %s err: %s", validPath, err)
	}
}

func TestItems_LoadItemsBySearchTerm(t *testing.T) {
	itemsDB, err := NewItems(databasePath)
	assert.NoError(t, err)
	resultItems, err := itemsDB.LoadItemsBySearchTerm("camera")
	assert.NoError(t, err)
	assert.NotEmpty(t, resultItems)
}

func TestItems_LoadItemsByInvalidSearchTerm(t *testing.T) {
	itemsDB, err := NewItems(databasePath)
	assert.NoError(t, err)
	resultItems, err := itemsDB.LoadItemsBySearchTerm("invalidSearchTerm")
	assert.NoError(t, err)
	assert.Empty(t, resultItems)
}
