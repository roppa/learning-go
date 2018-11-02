package main

import (
	"log"
	"time"
)

var waitCh = make(chan bool)

func main() {
	start()
	<-waitCh
}

func start() {
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		for range ticker.C {
			log.Println("Hello")
		}
	}()

	ticker2 := time.NewTicker(5 * time.Second)
	go func() {
		for range ticker2.C {
			waitCh <- true
		}
	}()
}
