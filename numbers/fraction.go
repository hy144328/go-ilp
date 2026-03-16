package numbers

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrNotInteger = errors.New("not an integer")
	ErrZeroDivisor = errors.New("zero divisor")
)

// A fraction is an implementation of a Rational.
type fraction[T constraints.Integer] struct {
	num   T
	denom T
}

// NewFraction constructs a fraction from a numerator and a denominator.
// It raises an exception if the denominator is zero.
func NewFraction[T constraints.Integer](
	num T,
	denom T,
) (fraction[T], error) {
	var res fraction[T]

	if denom == 0 {
		return res, fmt.Errorf("%w: %d/%d.", ErrZeroDivisor, num, denom)
	} else if denom < 0 {
		num = -num
		denom = -denom
	}

	divisor := GreatestCommonDivisor(num, denom)
	res.num = num / divisor
	res.denom = denom / divisor

	return res, nil
}

// FromInteger constructs a fraction from a numerator, and sets the denominator to one.
func FromInteger[T constraints.Integer](num T) fraction[T] {
	res, err := NewFraction(num, 1)
	if err != nil {
		panic(err)
	}
	return res
}

// ZeroFraction constructs a zero fraction.
func ZeroFraction[T constraints.Integer]() fraction[T] {
	return FromInteger[T](0)
}

// Numerator returns the numerator of a fraction.
func (x fraction[T]) Numerator() T {
	return x.num
}

// Denominator returns the denominator of a fraction.
func (x fraction[T]) Denominator() T {
	return x.denom
}

// Add adds one fraction to another fraction.
func (x fraction[T]) Add(y Rational[T]) Rational[T] {
	res, err := NewFraction(
		x.num*y.Denominator()+x.denom*y.Numerator(),
		x.denom*y.Denominator(),
	)
	if err != nil {
		panic(err)
	}

	return res
}

// Mul multiplies one fraction by another fraction.
func (x fraction[T]) Mul(y Rational[T]) Rational[T] {
	res, err := NewFraction(x.num*y.Numerator(), x.denom*y.Denominator())
	if err != nil {
		panic(err)
	}

	return res
}

// MulT multiplies a fraction by an integer.
func (x fraction[T]) MulT(fac T) Rational[T] {
	res, err := NewFraction(x.num*fac, x.denom)
	if err != nil {
		panic(err)
	}

	return res
}

// Equals reports whether one fraction is equal to another fraction.
func (x fraction[T]) Equals(y Rational[T]) bool {
	lhs := x.num * y.Denominator()
	rhs := x.denom * y.Numerator()
	return lhs == rhs
}

// LessThan reports whether one fraction is less than another fraction.
func (x fraction[T]) LessThan(y Rational[T]) bool {
	lhs := x.num * y.Denominator()
	rhs := x.denom * y.Numerator()
	return lhs < rhs
}

// GreaterThan reports whether one fraction is greater than another fraction.
func (x fraction[T]) GreaterThan(y Rational[T]) bool {
	lhs := x.num * y.Denominator()
	rhs := x.denom * y.Numerator()
	return lhs > rhs
}

// LessEqual reports whether one fraction is less than or equal to another fraction.
func (x fraction[T]) LessEqual(y Rational[T]) bool {
	lhs := x.num * y.Denominator()
	rhs := x.denom * y.Numerator()
	return lhs <= rhs
}

// GreaterEqual reports whether one fraction is greater than or equal to another fraction.
func (x fraction[T]) GreaterEqual(y Rational[T]) bool {
	lhs := x.num * y.Denominator()
	rhs := x.denom * y.Numerator()
	return lhs >= rhs
}

// Floor return the greatest integer less than or equal to the fraction.
func (x fraction[T]) Floor() T {
	mod := Remainder(x.num, x.denom)
	return (x.num - mod) / x.denom
}

// Floor return the lowest integer greater than or equal to the fraction.
func (x fraction[T]) Ceil() T {
	mod := Remainder(x.num - 1, x.denom) + 1
	return (x.num - mod) / x.denom + 1
}

// IsInteger reports whether the fraction is an integer.
func (x fraction[T]) IsInteger() bool {
	return x.num % x.denom == 0
}

// ToInteger returns the value of the fraction as an integer.
// It raises an exception if the value is not an integer.
func (x fraction[T]) ToInteger() (T, error) {
	if !x.IsInteger() {
		return 0, fmt.Errorf("%w: %v.", ErrNotInteger, x)
	}

	return x.num / x.denom, nil
}
