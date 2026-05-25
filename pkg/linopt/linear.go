package linopt

import (
	"fmt"

	"github.com/hy144328/go-ilp/pkg/linalg"
	"golang.org/x/exp/constraints"
)

// A LinearProgram represents a linear program in standard formulation.
type LinearProgram[T constraints.Signed] struct {
	Tab linalg.Tableau[T]
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
	res.Tab = tab

	tab[0][0] = -1
	copy(tab[0][1:], form.C)
	tab.ScaleRow(0, -1)

	for rowCt, rowIt := range form.A {
		copy(tab[1+rowCt][1:], rowIt)
		tab[1+rowCt][1+noVariables] = form.B[rowCt]
	}

	return res, nil
}

// NoConstraints returns the number of constraints.
func (lp LinearProgram[T]) NoConstraints() int {
	return lp.Tab.NoRows() - 1
}

// NoVariables returns the number of variables.
func (lp LinearProgram[T]) NoVariables() int {
	return lp.Tab.NoColumns() - 2
}

// ToStandardForm converts a LinearProgram to StandardForm.
func (lp LinearProgram[T]) ToStandardForm() StandardForm[T] {
	return StandardForm[T]{
		A: lp.leftHandSide().Copy(),
		B: linalg.FromColumn(lp.rightHandSide(), 0),
		C: linalg.FromRow(lp.weights(), 0),
	}
}

func (lp LinearProgram[T]) leftHandSide() linalg.Matrix[T] {
	noConstraints := lp.NoConstraints()
	noVariables := lp.NoVariables()
	return lp.Tab.Slice(1, 1+noConstraints, 1, 1+noVariables)
}

func (lp LinearProgram[T]) rightHandSide() linalg.Matrix[T] {
	noConstraints := lp.NoConstraints()
	noVariables := lp.NoVariables()
	return lp.Tab.Slice(1, 1+noConstraints, 1+noVariables, 2+noVariables)
}

func (lp LinearProgram[T]) weights() linalg.Matrix[T] {
	return lp.Tab.Slice(0, 1, 1, 1+lp.NoVariables())
}

// Reduce minimizes the number of non-trivial constraints in a LinearProgram.
// The underlying tableau is modified in place but no rows are dropped.
// The first non-zero coefficient of each non-trivial constraint is returned.
func (lp LinearProgram[T]) Reduce() ([]int, error) {
	lse := linalg.LinearSystemOfEquations[T]{
		Tab: lp.Tab,
	}
	pivots, err := lse.Reduce()

	pivots = pivots[1:]
	for pivotCt, pivotIt := range pivots {
		pivots[pivotCt] = pivotIt - 1
	}

	return pivots, err
}
