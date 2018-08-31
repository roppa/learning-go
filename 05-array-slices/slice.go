package main

import "fmt"

func sliceExample() {
	stringSlice := make([]string, 3)
	fmt.Println("emp:", stringSlice)

	stringSlice[0] = "a"
	stringSlice[1] = "b"
	stringSlice[2] = "c"
	fmt.Println("set:", stringSlice)
	fmt.Println("get:", stringSlice[2])

	fmt.Println("len:", len(stringSlice))
}
