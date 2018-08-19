package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSplitWords(t *testing.T) {
	searchTerm := "camera playstation"
	words := splitWords(searchTerm)
	assert.Equal(t, len(words), 2)
	assert.Equal(t, "camera", words[0])
	assert.Equal(t, "playstation", words[1])
}

func TestSplitWordsSingleWord(t *testing.T) {
	searchTerm := "camera"
	words := splitWords(searchTerm)
	assert.Equal(t, len(words), 1)
	assert.Equal(t, "camera", words[0])
}

func TestGenerateQuery(t *testing.T) {
	searchTerm := "camera playstation"
	query := generateQuery(searchTerm)
	assert.Contains(t, query, "%camera%")
	assert.Contains(t, query, "%playstation%")
}
