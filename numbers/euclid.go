package numbers

import (
	"golang.org/x/exp/constraints"
)

// GreatestCommonDivisor calculates the greatest common divisor of the given numbers.
// It returns 0 if all numbers are 0.
func GreatestCommonDivisor[T constraints.Integer](a T, as ...T) T {
	a = Abs(a)

	for _, b := range as {
		b = Abs(b)

		if a < b {
			a, b = b, a
		}

		for b > 0 {
			a, b = b, a%b
		}
	}

	return a
}

// LeastCommonMultiple calculates the least common multiple of the given numbers.
// It returns 0 if any number is 0.
func LeastCommonMultiple[T constraints.Integer](a T, as ...T) T {
	a = Abs(a)

	for _, b := range as {
		b = Abs(b)

		divisor := GreatestCommonDivisor(a, b)
		if divisor == 0 {
			return 0
		}

		a *= b / divisor
	}

	return a
}
