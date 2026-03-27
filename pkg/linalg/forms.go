package linalg

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrNotSolution = errors.New("not a solution")
)

// A LinearForm represents the matrix-vector formulation of a lineary system of equations.
type LinearForm[T constraints.Integer] struct {
	A Matrix[T]
	B Vector[T]
}

// Validate throws an error if the solution does not satisfy the constraints of the LinearForm.
func (form LinearForm[T]) Validate(sol Vector[T]) error {
	res, err := form.A.MulVec(sol)
	if err != nil {
		return err
	}

	for resCt := range res {
		if res[resCt] != form.B[resCt] {
			return fmt.Errorf("%w: %v dot %v != %d", ErrNotSolution, form.A[resCt], sol, form.B[resCt])
		}
	}

	return nil
}
