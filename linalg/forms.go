package linalg

import (
	"golang.org/x/exp/constraints"
)

type LinearForm[T constraints.Integer] struct {
	A Matrix[T]
	B Vector[T]
}
