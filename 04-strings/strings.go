package main

import "fmt"

func main() {
	var h string
	h = "Hello"
	w := "world"
	fmt.Println(h + " " + w)

	hw := "Hello, \n\t\"World\"!"
	fmt.Println(hw)

	s := `Hello,
	"World!"`
	fmt.Printf(s)

	e := "Hello world!"
	e2 := "👋🏻 🌍"
	fmt.Printf(e + " " + e2)
}
