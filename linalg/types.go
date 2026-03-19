package linalg

import (
	"errors"
	"fmt"

	"github.com/hy144328/go-ilp/numbers"
	"golang.org/x/exp/constraints"
)

var (
	ErrIncompatibleSizes = errors.New("incompatible sizes")
)

// A Vector is a one-dimensional tensor.
type Vector[T constraints.Integer] []T

// A VectorRational is a one-dimensional tensor over the rational numbers.
type VectorRational[T constraints.Integer] []numbers.Rational[T]

// A Matrix is a two-dimensional tensor.
type Matrix[T constraints.Integer] [][]T

// NewVector creates a Vector of given length.
func NewVector[T constraints.Integer](size int) Vector[T] {
	return make([]T, size)
}

// NewVector creates a Vector of given length.
func NewVectorRational[T constraints.Integer](size int) VectorRational[T] {
	return make([]numbers.Rational[T], size)
}

// NewMatrix creates a Matrix with given numbers of rows and columns.
func NewMatrix[T constraints.Integer](noRows int, noColumns int) Matrix[T] {
	res := make([][]T, noRows)

	for rowCt := range noRows {
		res[rowCt] = make([]T, noColumns)
	}

	return res
}

// Size returns the length of the Vector.
func (vec Vector[T]) Size() int {
	return len(vec)
}

// Dot calculates the dot product of one Vector with another Vector.
func (vec Vector[T]) Dot(other Vector[T]) (T, error) {
	if vec.Size() != other.Size() {
		return 0, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, vec.Size(), other.Size())
	}

	var res T

	for i := range vec {
		res += vec[i] * other[i]
	}

	return res, nil
}

// Size returns the length of the VectorRational.
func (vec VectorRational[T]) Size() int {
	return len(vec)
}

// NoRows returns the number of rows of the Matrix.
func (mat Matrix[T]) NoRows() int {
	return len(mat)
}

// NoColumns returns the number of columns of the Matrix.
func (mat Matrix[T]) NoColumns() int {
	return len(mat[0])
}

// Mul multiplies a Matrix with a Vector.
func (mat Matrix[T]) Mul(vec Vector[T]) (Vector[T], error) {
	if mat.NoColumns() != vec.Size() {
		return nil, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, mat.NoColumns(), vec.Size())
	}

	res := NewVector[T](mat.NoRows())

	for rowCt, rowIt := range mat {
		resIt, err := Vector[T](rowIt).Dot(vec)
		if err != nil {
			panic(err)
		}
		res[rowCt] = resIt
	}

	return res, nil
}
