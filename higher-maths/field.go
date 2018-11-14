package maths

import (
	"fmt"
	"math"
)

// Field set of finite numbers
type Field struct {
	value int32
	prime int32
}

func (f *Field) eq(other Field) bool {
	return f.value == other.value && f.prime == other.prime
}

func (f *Field) ne(other Field) bool {
	return f.value != other.value
}

func (f *Field) add(other Field) (Field, error) {
	if f.prime != other.prime {
		return Field{}, fmt.Errorf("Cannot add two numbers in different Fields")
	}
	result, error := NewField((f.value+other.value)%f.prime, f.prime)
	return result, error
}

func (f *Field) sub(other Field) (Field, error) {
	if f.prime != other.prime {
		return Field{}, fmt.Errorf("Cannot subtract two numbers in different Fields")
	}
	result, error := NewField((f.value-other.value)%f.prime, f.prime)
	return result, error
}

func (f *Field) mul(num int32) (Field, error) {
	field, error := NewField((f.value*num)%f.prime, f.prime)
	return field, error
}

func (f *Field) pow(exponent int32) (Field, error) {
	n := exponent % (f.prime - 1)
	num := math.Mod(math.Pow(float64(f.value), float64(n)), float64(f.prime))
	result, error := NewField(int32(num), f.prime)
	return result, error
}

func (f *Field) repr() string {
	return fmt.Sprintf("Field_%d(%d)", f.prime, f.value)
}

// NewField constructs a new Field
func NewField(num int32, prime int32) (Field, error) {
	field := Field{}
	if num >= prime || num < 0 {
		return field, fmt.Errorf("Number %d not in field range 0 to %d", num, prime-1)
	}
	field.value = num
	field.prime = prime
	return field, nil
}
