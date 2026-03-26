package linalg

import (
	"errors"
	"fmt"

	"golang.org/x/exp/constraints"
)

var (
	ErrNoSolution = errors.New("no solution")
	ErrNotSolution = errors.New("not a solution")
)

// A LinearSystemOfEquations is defined by a Matrix on the left-hand side and a Vector on the right-hand side.
// They are combined in a single tableau.
type LinearSystemOfEquations[T constraints.Signed] struct {
	tab Tableau[T]
}

// NewLinearSystemOfEquations constructs a LinearSystemOfEquations from a Matrix and a Vector.
func FromLinearForm[T constraints.Signed](
	form LinearForm[T],
) (LinearSystemOfEquations[T], error) {
	noConstraints := form.A.NoRows()
	noVariables := form.A.NoColumns()
	res := LinearSystemOfEquations[T]{}

	if size := form.B.Size(); size != noConstraints {
		return res, fmt.Errorf("%w: %d != %d.", ErrIncompatibleSizes, size, noConstraints)
	}

	tab := NewTableau[T](noConstraints, noVariables+1)
	res.tab = tab

	for rowCt, rowIt := range form.A {
		copy(tab[rowCt], rowIt)
		tab[rowCt][noVariables] = form.B[rowCt]
	}

	return res, nil
}

// NoConstraints returns the number of constraints.
func (lse LinearSystemOfEquations[T]) NoConstraints() int {
	return lse.tab.NoRows()
}

// NoVariables returns the number of variables.
func (lse LinearSystemOfEquations[T]) NoVariables() int {
	return lse.tab.NoColumns() - 1
}

// ToLinearForm converts a LinearSystemOfEquations to LinearForm.
func (lse LinearSystemOfEquations[T]) ToLinearForm() LinearForm[T] {
	return LinearForm[T]{
		A: lse.leftHandSide(),
		B: lse.rightHandSide().ToVector(0),
	}
}

// leftHandSide reconstructs the Matrix on the left-hand side.
func (lse LinearSystemOfEquations[T]) leftHandSide() Matrix[T] {
	noConstraints := lse.NoConstraints()
	noVariables := lse.NoVariables()
	return lse.tab.Slice(0, noConstraints, 0, noVariables)
}

// rightHandSide reconstructs the Vector on the right-hand side.
func (lse LinearSystemOfEquations[T]) rightHandSide() Matrix[T] {
	noConstraints := lse.NoConstraints()
	noVariables := lse.NoVariables()
	return lse.tab.Slice(0, noConstraints, noVariables, noVariables+1)
}

// Validate throws an error if the solution does not satisfy all constraints.
func (lse LinearSystemOfEquations[T]) Validate(sol Vector[T]) error {
	form := lse.ToLinearForm()

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

// Reduce minimizes the number of independent constraints.
func (lse LinearSystemOfEquations[T]) Reduce() ([]int, error) {
	noConstraints := lse.NoConstraints()
	noVariables := lse.NoVariables()
	pivots := make([]int, 0, noConstraints)

	for rowCt, colCt := 0, 0; rowCt < noConstraints && colCt < noVariables; {
		PivotColumn(lse.tab, rowCt, colCt)
		if lse.tab[rowCt][colCt] == 0 {
			colCt++
			continue
		}

		pivots = append(pivots, colCt)
		if err := EliminateDown(lse.tab, rowCt, colCt); err != nil {
			panic(err)
		}

		rowCt++
		colCt++
	}

	for rowCt := len(pivots); rowCt < noConstraints; rowCt++ {
		if lse.tab[rowCt][noVariables] != 0 {
			return pivots, fmt.Errorf("%w: Inhomogeneous null row.\n\n%v\n", ErrNoSolution, lse.tab)
		}
	}
	lse.tab = lse.tab[:len(pivots)]

	return pivots, nil
}
