package main

import "fmt"

// rarely use arrays directly, better to use slice

func main() {
	var vals [3]int
	vals[0] = 1
	vals[1] = 2
	vals[2] = 3
	fmt.Println(vals)
}
