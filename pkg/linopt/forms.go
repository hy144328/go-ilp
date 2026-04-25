package linopt

import (
	"errors"
	"fmt"

	"github.com/hy144328/go-ilp/pkg/linalg"
	"golang.org/x/exp/constraints"
)

var (
	ErrNegativeSolution = errors.New("negative solution")
)

// A CanonicalForm represents the canonical formulation of a linear program:
// Optimize c.x, subject to ax <= b and x >= 0.
type CanonicalForm[T constraints.Integer] struct {
	A linalg.Matrix[T]
	B linalg.Vector[T]
	C linalg.Vector[T]
}

// A StandardForm represents the standard formulation of a linear program:
// Optimize c.x, subject to ax = b and x >= 0.
type StandardForm[T constraints.Integer] struct {
	A linalg.Matrix[T]
	B linalg.Vector[T]
	C linalg.Vector[T]
}

// ToStandard converts from CanonicalForm to StandardForm.
func (form CanonicalForm[T]) ToStandard() StandardForm[T] {
	noRows := form.A.NoRows()
	noColumns := form.A.NoColumns()

	a := linalg.NewMatrix[T](
		noRows,
		noRows+noColumns,
	)
	for rowCt, rowIt := range form.A {
		copy(a[rowCt], rowIt)
		a[rowCt][noColumns+rowCt] = 1
	}

	b := form.B.Copy()

	c := linalg.NewVector[T](form.C.Size() + noColumns)
	copy(c, form.C)

	return StandardForm[T]{a, b, c}
}

// Validate throws an error if the solution does not satisfy the constraints of the CanonicalForm.
func (form CanonicalForm[T]) Validate(sol linalg.Vector[T]) error {
	for solCt, solIt := range sol {
		if solIt < 0 {
			return fmt.Errorf("%w: sol[%d] = %d >= 0.", ErrNegativeSolution, solCt, solIt)
		}
	}

	res, err := form.A.MulVec(sol)
	if err != nil {
		return err
	}

	for resCt := range res {
		if res[resCt] > form.B[resCt] {
			return fmt.Errorf(
				"%w: A x = y <= b\n\nA:\n%v\n\nx:\n%v\n\ny:\n%v\n\nb:\n%v\n",
				linalg.ErrNotSolution,
				form.A,
				sol,
				res,
				form.B,
			)
		}
	}

	return nil
}

// Validate throws an error if the solution does not satisfy the constraints of the StandardForm.
func (form StandardForm[T]) Validate(sol linalg.Vector[T]) error {
	for solCt, solIt := range sol {
		if solIt < 0 {
			return fmt.Errorf("%w: sol[%d] = %d >= 0.", ErrNegativeSolution, solCt, solIt)
		}
	}

	res, err := form.A.MulVec(sol)
	if err != nil {
		return err
	}

	if !res.Equals(form.B) {
		return fmt.Errorf(
			"%w: A x = y = b\n\nA:\n%v\n\nx:\n%v\n\ny:\n%v\n\nb:\n%v\n",
			linalg.ErrNotSolution,
			form.A,
			sol,
			res,
			form.B,
		)
	}

	return nil
}
