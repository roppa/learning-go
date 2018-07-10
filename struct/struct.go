package main

import "fmt"

type user struct {
	name  string
	email string
	age   int
}

func main() {

	bob := user{
		name:  "Bob",
		email: "bob@johnson.com",
		age:   45,
	}

	fmt.Println("Name", bob.name)
	fmt.Println("Email", bob.email)
	fmt.Println("Age", bob.age)

	paul := struct {
		name  string
		email string
		age   int
	}{
		name:  "Paul",
		email: "paul@thompson.com",
		age:   46,
	}

	fmt.Println("Name", paul.name)
	fmt.Println("Email", paul.email)
	fmt.Println("Age", paul.age)
}
