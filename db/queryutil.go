package db

import "strings"

const seperator = " "

func generateQuery(searchTerm string) string {
	baseQuery := "SELECT * FROM items" +
		" WHERE"
	words := splitWords(searchTerm)
	searchQuery := baseQuery
	for i := 0; i < len(words); i++ {
		searchQuery += " item_name LIKE '%" + words[i] + "%'"
		if i != len(words)-1 {
			searchQuery += " OR "
		}
	}
	return searchQuery
}

func splitWords(words string) []string {
	return strings.Split(words, seperator)
}
