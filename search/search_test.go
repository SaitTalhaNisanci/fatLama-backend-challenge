package search

import (
	"testing"

	"github.com/fatLama-backend-challenge/model"
	"github.com/stretchr/testify/assert"
)

func TestCalculateScore(t *testing.T) {
	item := model.NewItem("canon camera", 12.212, 21.2133, "", "")
	searchParams, _ := NewSearchParams("camera", "12.212", "21.2133")
	score := calculateScore(item, searchParams)
	assert.InEpsilon(t, score, textMatchCoeff*1, 0.1)
}

func TestCalculateScoreNoMatch(t *testing.T) {
	item := model.NewItem("canon camera", 12.212, 21.2133, "", "")
	searchParams, _ := NewSearchParams("playstation", "12.212", "21.2133")
	score := calculateScore(item, searchParams)
	assert.Equal(t, score, textMatchCoeff*0)
}

func TestCalculateScoreTwoItems(t *testing.T) {
	item1 := model.NewItem("canon camera", 12.212, 21.2133, "", "")
	item2 := model.NewItem("canon camera", 14.212, 21.2133, "", "")
	searchParams, _ := NewSearchParams("playstation", "12.512", "21.2133")
	score1 := calculateScore(item1, searchParams)
	score2 := calculateScore(item2, searchParams)
	if score1 <= score2 {
		t.Error("item1 should have a higher score than item2 as it is closer.")
	}
}

func TestCalculateScoreTwoItemsDifferentNames(t *testing.T) {
	item1 := model.NewItem("canon camera", 12.212, 21.2133, "", "")
	item2 := model.NewItem("playstation with 3 games", 12.000, 21.2133, "", "")
	searchParams, _ := NewSearchParams("playstation", "12.512", "21.2133")
	score1 := calculateScore(item1, searchParams)
	score2 := calculateScore(item2, searchParams)
	if score1 >= score2 {
		t.Error("item2 should have a higher score than item1 as it matches one word in searchTerm.")
	}
}

func TestCalculateScoreDuplicates(t *testing.T) {
	item1 := model.NewItem("camera camera camera", 12.212, 21.2133, "", "")
	item2 := model.NewItem("camera", 12.212, 21.2133, "", "")
	searchParams, _ := NewSearchParams("playstation", "12.512", "21.2133")
	score1 := calculateScore(item1, searchParams)
	score2 := calculateScore(item2, searchParams)
	assert.Equal(t, score1, score2)
}
