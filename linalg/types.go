package linalg

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrIncompatibleSizes = errors.New("incompatible sizes")
)

// A Vector is a one-dimensional tensor.
type Vector[T constraints.Integer] []T

// A Matrix is a two-dimensional tensor.
type Matrix[T constraints.Integer] [][]T

// NewVector creates a Vector of given length.
func NewVector[T constraints.Integer](size int) Vector[T] {
	return make([]T, size)
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

// AsMatrix casts a Vector to a single-column Matrix.
func (vec Vector[T]) AsMatrix() Matrix[T] {
	res := make([][]T, vec.Size())

	for vecCt, vecIt := range vec {
		res[vecCt][0] = vecIt
	}

	return res
}

// ToMatrix copies a Vector to a single-column Matrix.
func (vec Vector[T]) ToMatrix() Matrix[T] {
	res := NewMatrix[T](vec.Size(), 1)

	for vecCt := range vec {
		res[vecCt] = vec[vecCt : vecCt+1]
	}

	return res
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

// NoRows returns the number of rows of the Matrix.
func (mat Matrix[T]) NoRows() int {
	return len(mat)
}

// NoColumns returns the number of columns of the Matrix.
func (mat Matrix[T]) NoColumns() int {
	return len(mat[0])
}

// ToVector copies a column of a Matrix to a Vector.
func (mat Matrix[T]) ToVector(idx int) Vector[T] {
	res := NewVector[T](mat.NoRows())

	for rowCt, rowIt := range mat {
		res[rowCt] = rowIt[idx]
	}

	return res
}

// Mul multiplies a Matrix with another Matrix.
func (mat Matrix[T]) Mul(other Matrix[T]) (Matrix[T], error) {
	if mat.NoColumns() != other.NoRows() {
		return nil, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, mat.NoColumns(), other.NoRows())
	}

	res := NewMatrix[T](mat.NoRows(), other.NoColumns())

	for rowCt, rowIt := range mat {
		for colCt := range other.NoColumns() {
			resIt, err := other.ToVector(colCt).Dot(rowIt)
			if err != nil {
				panic(err)
			}
			res[rowCt][colCt] = resIt
		}
	}

	return res, nil
}

// MulVec multiplies a Matrix with a Vector.
func (mat Matrix[T]) MulVec(vec Vector[T]) (Vector[T], error) {
	if mat.NoColumns() != vec.Size() {
		return nil, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, mat.NoColumns(), vec.Size())
	}

	res := NewVector[T](mat.NoRows())

	for rowCt, rowIt := range mat {
		resIt, err := vec.Dot(rowIt)
		if err != nil {
			panic(err)
		}
		res[rowCt] = resIt
	}

	return res, nil
}
