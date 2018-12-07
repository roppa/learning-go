package main

import "fmt"

func main() {
	var i int64
	i = 4200000000000000000
	j := &i
	k := *j
	fmt.Println(i, j, k)
}
