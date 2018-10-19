package arithmetic

import "testing"

func TestAddNoValues(t *testing.T) {
	want := 0
	got := Add()
	if got != want {
		t.Errorf("Want %v got %v", want, got)
	}
}

func TestAddOneValue(t *testing.T) {
	want := 1
	got := Add(1)
	if got != want {
		t.Errorf("Want %v got %v", want, got)
	}
}

func TestAddTwoValues(t *testing.T) {
	want := 3
	got := Add(1, 2)
	if got != want {
		t.Errorf("Want %v got %v", want, got)
	}
}

func TestAddMultipleValues(t *testing.T) {
	want := 10
	got := Add(1, 2, 3, 4)
	if want != got {
		t.Errorf("Want %v got %v", want, got)
	}
}
