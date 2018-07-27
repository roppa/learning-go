package main

import (
	"fmt"
	"os"
)

func main() {
	var args string
	if len(os.Args) > 1 {
		args = os.Args[1]
	}
	a := 10

	if a < 5 {
		fmt.Println("a is bigger than 5")
	} else {
		fmt.Println("a is less than or equal to 5")
	}

	for i := 0; i < 10; i++ {
		if i > 7 {
			break
		}
		fmt.Println(i)
	}

	j := 0
	for {
		j++
		if j > 7 {
			fmt.Println("Breaking out of empty for")
			break
		}
	}

	hw := "Hello world"
	for key, value := range hw {
		fmt.Println(key, value, string(value))
	}

	switch args {
	case "hello":
		fmt.Println("You said hello")
	case "password":
		fmt.Println("Not telling")
	default:
		fmt.Println("No arguments")
	}

}
