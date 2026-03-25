package numbers

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrNotInteger = errors.New("not an integer")
)

type Embedding[T constraints.Integer] interface {
	IsInteger() bool
	ToInteger() (T, error)
}

type Ring[T constraints.Integer, U any] interface {
	Add(U) U
	AddT(T) U
	Mul(Rational[T]) Rational[T]
	MulT(T) Rational[T]
}

type Comparable[T constraints.Integer, U any] interface {
	Equals(U) bool
	EqualsT(T) bool
	LessThan(U) bool
	LessThanT(T) bool
	GreaterThan(U) bool
	GreaterThanT(T) bool
	LessEqual(U) bool
	LessEqualT(T) bool
	GreaterEqual(U) bool
	GreaterEqualT(T) bool
}

type Archimedean[T constraints.Integer, U any] interface {
	Floor() T
	Ceil() T
}

// A Rational represents a rational number.
type Rational[T constraints.Integer] interface {
	Numerator() T
	Denominator() T

	Embedding[T]
	Ring[T, Rational[T]]
	Comparable[T, Rational[T]]
	Archimedean[T, Rational[T]]
	fmt.Stringer
}
