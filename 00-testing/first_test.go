package example

import "testing"

/*
	Remember test files need to be named [your file]_test.go
	To run: go test
*/
func TestSayHello(t *testing.T) {
	target := "Hello world!"
	result := sayHello()
	if result != target {
		t.Errorf("Incorrect greeting, got %s, want %s", result, target)
	}
}
