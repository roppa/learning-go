package datastructures

import (
	"fmt"
	"strings"
)

func main() {
	source := "this is a this test how now this brown test"
	words := strings.Fields(source)
	m := make(map[string]int)
	for _, word := range words {
		m[word]++
	}
	fmt.Println(m)
}
