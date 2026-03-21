package main

import (
	"strings"
)

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	rawText := strings.Trim(lowerText, " ")
	split := strings.Split(rawText, " ")
	var clean []string
	for _, word := range split {
		if len(word) > 0 {
			clean = append(clean, word)
		}
	}
	return clean
}
