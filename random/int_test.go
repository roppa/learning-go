package main

import (
	"testing"
)

func RandomNumber(t *testing.T) {
	var first interface{} = randomNumber()
	var second interface{} = randomNumber()
	n1, ok1 := first.(int64)
	n2, ok2 := second.(int64)
	if !ok1 || !ok2 {
		t.Errorf("Got %v, %v wanted int64", n1, n2)
	}

	if n1 == n2 {
		t.Errorf("Got %v, %v wanted different numbers", n1, n2)
	}
}
