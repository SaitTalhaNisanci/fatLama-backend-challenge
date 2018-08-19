package db

import "testing"

func TestInitDBWithInvalidPath(t *testing.T) {
	invalidPath := "./invalid.sqlite3"
	itemsDB := &Items{}
	if _, err := itemsDB.InitDB(invalidPath); err == nil {
		t.Errorf("Database should not start with invalidPath %s", invalidPath)
	}
}

func TestInitDB(t *testing.T) {
	validPath := "../fatlama.sqlite3"
	itemsDB := &Items{}
	if _, err := itemsDB.InitDB(validPath); err != nil {
		t.Errorf("InitDB should not return an error for valid path %s err: %s", validPath, err)
	}
}
