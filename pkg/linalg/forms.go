package linalg

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrNotSolution = errors.New("not a solution")
)

// A LinearForm represents the matrix-vector formulation of a linear system of equations.
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

	if !res.Equals(form.B) {
		return fmt.Errorf(
			"%w: A x = y = b\n\nA:\n%v\n\nx:\n%v\n\ny:\n%v\n\nb:\n%v\n",
			ErrNotSolution,
			form.A,
			sol,
			res,
			form.B,
		)
	}

	return nil
}
