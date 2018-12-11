package main

import (
	"fmt"
)

/*
// this would cause a fatal error: all goroutines are asleep - deadlock!
func main() {
	channel := make(chan string)
	channel <- "helloooo"
	message := <-channel
	fmt.Println(message)
}
*/

func main() {
	channel := make(chan string)
	go func() {
		channel <- "helloooo"
	}()
	message := <-channel
	fmt.Println(message)
}
