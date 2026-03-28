package linopt

import (
	"fmt"

	"github.com/hy144328/go-ilp/pkg/linalg"
	"golang.org/x/exp/constraints"
)

// A LinearProgram represents a linear program in standard formulation.
type LinearProgram[T constraints.Signed] struct {
	tab linalg.Tableau[T]
}

// FromCanonical converts from StandardForm to LinearProgram.
func FromStandardForm[T constraints.Signed](form StandardForm[T]) (LinearProgram[T], error) {
	noConstraints := form.A.NoRows()
	noVariables := form.A.NoColumns()
	res := LinearProgram[T]{}

	if size := form.B.Size(); size != noConstraints {
	        return res, fmt.Errorf("%w: %d != %d.", linalg.ErrIncompatibleSizes, size, noConstraints)
	}

	if size := form.C.Size(); size != noVariables {
	        return res, fmt.Errorf("%w: %d != %d.", linalg.ErrIncompatibleSizes, size, noVariables)
	}

	tab := linalg.NewTableau[T](noConstraints+1, noVariables+2)
	res.tab = tab

	tab[0][0] = -1
	copy(tab[0][1:], form.C)

	for rowCt, rowIt := range form.A {
		copy(tab[1+rowCt][1:], rowIt)
		tab[1+rowCt][1+noVariables] = form.B[rowCt]
	}

	return res, nil
}

// NoConstraints returns the number of constraints.
func (problem LinearProgram[T]) NoConstraints() int {
	return problem.tab.NoRows() - 1
}

// NoVariables returns the number of variables.
func (problem LinearProgram[T]) NoVariables() int {
	return problem.tab.NoColumns() - 2
}

// ToStandardForm converts a LinearProgram to StandardForm.
func (problem LinearProgram[T]) ToStandardForm() StandardForm[T] {
	return StandardForm[T]{
		A: problem.leftHandSide().Copy(),
		B: linalg.FromColumn(problem.rightHandSide(), 0),
		C: linalg.FromRow(problem.weights(), 0),
	}
}

func (problem LinearProgram[T]) leftHandSide() linalg.Matrix[T] {
	noConstraints := problem.NoConstraints()
	noVariables := problem.NoVariables()
	return problem.tab.Slice(1, 1+noConstraints, 1, 1+noVariables)
}

func (problem LinearProgram[T]) rightHandSide() linalg.Matrix[T] {
	noConstraints := problem.NoConstraints()
	noVariables := problem.NoVariables()
	return problem.tab.Slice(1, 1+noConstraints, 1+noVariables, 2+noVariables)
}

func (problem LinearProgram[T]) weights() linalg.Matrix[T] {
	return problem.tab.Slice(0, 1, 1, 1+problem.NoVariables())
}
