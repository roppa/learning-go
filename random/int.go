package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(randomNumber())
}

func randomNumber() int64 {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	return int64(r1.Intn(99999))
}
