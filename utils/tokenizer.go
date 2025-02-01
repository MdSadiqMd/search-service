package utils

import (
	"strings"
	"unicode"
)

func tokenize(text string) []string {
	return strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
}

func analyze(text string) []string {
	tokens := tokenize(text)
	tokens = lowerCaseFilter(tokens)
	tokens = stopWordsFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
