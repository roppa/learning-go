package main

import (
	"fmt"
)

func main() {
	var h int               //will get initialised to default 0 value
	var i int = 10          //can specify the type
	var j = 10              //type is inferred
	k := 10                 //or you can use shorthand, but has to be inside a function
	fmt.Println(h, i, j, k) // 0 10 10 10
}
