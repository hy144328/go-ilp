package numbers

import (
	"golang.org/x/exp/constraints"
)

// Abs calculates the absolute value.
func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

// Remainder calculates the smallest non-negative integer that is congruent to x modulo y.
func Remainder[T constraints.Integer](x T, y T) T {
	y = Abs(y)

	return (x%y + y) % y
}
