package datastructures

import "testing"

func TestDictionary(t *testing.T) {
	source := "this is a this test how now this brown test"
	want := 7
	got := len(getDictionary(source))
	if got != want {
		t.Errorf("Want %v got %v", want, got)
	}
}
