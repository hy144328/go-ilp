package numbers

import (
	"golang.org/x/exp/constraints"
)

// A Rational represents a rational number.
type Rational[T constraints.Integer] interface {
	Numerator() T
	Denominator() T

	// Binary operations.
	Add(Rational[T]) Rational[T]
	AddT(T) Rational[T]
	Mul(Rational[T]) Rational[T]
	MulT(T) Rational[T]

	// Comparison operations.
	Equals(Rational[T]) bool
	EqualsT(T) bool
	LessThan(Rational[T]) bool
	LessThanT(T) bool
	GreaterThan(Rational[T]) bool
	GreaterThanT(T) bool
	LessEqual(Rational[T]) bool
	LessEqualT(T) bool
	GreaterEqual(Rational[T]) bool
	GreaterEqualT(T) bool

	// Integer conversions.
	Floor() T
	Ceil() T
	IsInteger() bool
	ToInteger() (T, error)
}
