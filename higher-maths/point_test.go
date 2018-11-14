package maths

import (
	"testing"
)

func TestNewPoint(t *testing.T) {
	_, error := NewPoint(-1, -1, 5, 7)
	if error != nil {
		t.Error("NewPoint with -1, -1, 5, 7 should return new Point")
	}
}
