package utils

import "strings"
import snowballstemmer "github.com/kljensen/snowball/english"

func lowerCaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopWordFilter(tokens []string) []string {
	var stopWords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {},
		"if": {}, "in": {}, "is": {}, "it": {}, "not": {},
		"the": {}, "to": {}, "was": {}, "will": {}, "with": {},
	}
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = snowballstemmer.Stem(token, false)
	}
	return r
}
