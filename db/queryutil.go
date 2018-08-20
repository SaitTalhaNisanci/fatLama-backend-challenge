package db

import "strings"

// separator is used to split search term.
const separator = " "

// generateQuery generates SQL query with search term.
// It adds a 'LIKE' pattern for each word in search term.
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

// splitWords split the words by space and returns a string slice.
func splitWords(words string) []string {
	return strings.Split(words, separator)
}
