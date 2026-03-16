package numbers

import (
	"golang.org/x/exp/constraints"
)

// A Rational represents a rational number.
type Rational[T constraints.Integer] interface {
	Numerator() T
	Denominator() T
	Add(Rational[T]) Rational[T]
	Mul(Rational[T]) Rational[T]
	MulT(T) Rational[T]
	Equals(Rational[T]) bool
	LessThan(Rational[T]) bool
	GreaterThan(Rational[T]) bool
	LessEqual(Rational[T]) bool
	GreaterEqual(Rational[T]) bool
	Floor() T
	Ceil() T
	IsInteger() bool
	ToInteger() (T, error)
}
