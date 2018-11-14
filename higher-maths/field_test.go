package maths

import (
	"reflect"
	"testing"
)

func TestNewField(t *testing.T) {
	field, error := NewField(0, 0)
	if error == nil {
		t.Error("NewField with 0, 0 should throw error")
	}

	if !reflect.DeepEqual(field, Field{}) {
		t.Errorf("Field should be empty")
	}
}

func TestEq(t *testing.T) {
	field, _ := NewField(0, 1)
	other, _ := NewField(0, 1)
	if !field.eq(other) {
		t.Errorf("Fields should be equal")
	}
}

func TestNe(t *testing.T) {
	field, _ := NewField(0, 1)
	other, _ := NewField(1, 1)
	if field.ne(other) {
		t.Errorf("Fields should not be equal")
	}
}

func TestAdd(t *testing.T) {
	field, _ := NewField(7, 13)
	other, _ := NewField(12, 13)
	target, _ := NewField(6, 13)
	result, error := field.add(other)
	if !result.eq(target) || error != nil {
		t.Errorf("Should add fields")
	}
}

func TestIncorrectAdd(t *testing.T) {
	field, _ := NewField(7, 13)
	other, _ := NewField(12, 14)
	_, error := field.add(other)
	if error == nil {
		t.Errorf("Should not be able to add fields with different primes")
	}
}

func TestSub(t *testing.T) {
	field, _ := NewField(11, 19)
	other, _ := NewField(9, 19)
	target, _ := NewField(2, 19)
	result, error := field.sub(other)
	if !result.eq(target) || error != nil {
		t.Errorf("Should subtract fields")
	}
}

func TestMul(t *testing.T) {
	field, _ := NewField(5, 19)
	target, _ := NewField(15, 19)
	result, error := field.mul(3)
	if !result.eq(target) || error != nil {
		t.Errorf("Should multiply fields")
	}
}

func TestPow(t *testing.T) {
	field, _ := NewField(7, 19)
	target, _ := NewField(1, 19)
	result, error := field.pow(3)
	if !result.eq(target) || error != nil {
		t.Errorf("Should exponentiate number")
	}
}

func TestRepr(t *testing.T) {
	field, _ := NewField(0, 1)
	if field.repr() != "Field_1(0)" {
		t.Errorf("ToString not ")
	}
}
