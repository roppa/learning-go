package datastructures

import (
	"strings"
)

func getDictionary(source string) map[string]int {
	words := strings.Fields(source)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	return m
}
