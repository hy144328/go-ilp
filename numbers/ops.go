package numbers

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

func Remainder[T constraints.Integer](x T, y T) T {
	y = Abs(y)

	return (x%y + y) % y
}
