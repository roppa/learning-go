package datastructures

import "testing"

func TestAddNoValues(t *testing.T) {
	want := 0
	got := Dictionary()
	if got != want {
		t.Errorf("Want %v got %v", want, got)
	}
}
