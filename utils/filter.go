package utils

import (
	"strings"

	snowballeng "github.com/kljensen/snowball/english"
)

func lowerCaseFilter(tokens []string) []string {
	r := make([]string, len(tokens))
	for i, token := range tokens {
		r[i] = strings.ToLower(token)
	}
	return r
}

func stopWordsFilter(tokens []string) []string {
	var stopWords = map[string]struct{}{
		"a": {}, "and": {}, "be": {}, "have": {}, "i": {}, "in": {}, "of": {}, "that": {}, "the": {}, "to": {},
	}
	r := make([]string, 0, len(tokens))
	for _, token := range tokens {
		if _, ok := stopWords[token]; !ok {
			r = append(r, token)
		}
	}
	return r
}

func stemmerFilter(token []string) []string {
	r := make([]string, len(token))
	for i, token := range token {
		r[i] = snowballeng.Stem(token, false)
	}
	return r
}
