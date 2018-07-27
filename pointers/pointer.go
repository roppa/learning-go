package main

import "fmt"

type person struct {
	Name string
	Age  int
}

func showPerson(p *person) {
	p.Name = "Bobson"
}

func main() {
	count := 10
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")
	increment(&count)
	println("count:\tValue Of[", count, "]\tAddr Of[", &count, "]")

	bob := person{Name: "Bob", Age: 34}
	showPerson(&bob)
	fmt.Println(bob)
}

func increment(inc *int) {
	*inc++
	println("inc:\tValue Of[", inc, "]\tAddr Of[", &inc, "\tValue Points To[", *inc, "]")
}
