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
func NewLinearSystemOfEquations[T constraints.Signed](
	A Matrix[T],
	b Vector[T],
) (LinearSystemOfEquations[T], error) {
	if A.NoRows() != b.Size() {
		return *new(LinearSystemOfEquations[T]), fmt.Errorf("%w: %d != %d.", ErrIncompatibleSizes, A.NoRows(), b.Size())
	}

	noConstraints := A.NoRows()
	noVariables := A.NoColumns()
	tab := NewTableau[T](noConstraints, noVariables+1)

	for rowCt, rowIt := range A {
		copy(tab[rowCt], rowIt)
		tab[rowCt][noVariables] = b[rowCt]
	}

	return LinearSystemOfEquations[T]{tab}, nil
}

// NoConstraints returns the number of constraints.
func (lse LinearSystemOfEquations[T]) NoConstraints() int {
	return lse.tab.NoRows()
}

// NoVariables returns the number of variables.
func (lse LinearSystemOfEquations[T]) NoVariables() int {
	return lse.tab.NoColumns() - 1
}

// leftHandSide reconstructs the Matrix on the left-hand side.
func (lse LinearSystemOfEquations[T]) leftHandSide() Matrix[T] {
	noConstraints := lse.NoConstraints()
	noVariables := lse.NoVariables()
	res := make([][]T, noConstraints)

	for rowCt, rowIt := range lse.tab {
		res[rowCt] = rowIt[:noVariables]
	}

	return res
}

// rightHandSide reconstructs the Vector on the right-hand side.
func (lse LinearSystemOfEquations[T]) rightHandSide() Matrix[T] {
	noConstraints := lse.NoConstraints()
	noVariables := lse.NoVariables()
	res := make([][]T, noConstraints)

	for rowCt, rowIt := range lse.tab {
		res[rowCt] = rowIt[noVariables:]
	}

	return res
}

// Validate throws an error if the solution does not satisfy all constraints.
func (lse LinearSystemOfEquations[T]) Validate(sol Vector[T]) error {
	lhs := lse.leftHandSide()
	rhs := lse.rightHandSide()

	res, err := lhs.MulVec(sol)
	if err != nil {
		return err
	}

	for resCt := range res {
		if res[resCt] != rhs[resCt][0] {
			return fmt.Errorf("%w: %v dot %v != %d", ErrNotSolution, lhs[resCt], sol, rhs[resCt])
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
