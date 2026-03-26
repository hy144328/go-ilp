package linopt

import (
	"github.com/hy144328/go-ilp/linalg"
	"golang.org/x/exp/constraints"
)

// A LinearProgram represents a linear program in standard formulation.
type LinearProgram[T constraints.Integer] struct {
	tab linalg.Tableau[T]
}

// FromCanonical converts from CanonicalForm to LinearProgram.
func FromCanonical[T constraints.Integer](form CanonicalForm[T]) LinearProgram[T] {
	return FromStandard(form.ToStandard())
}

// FromCanonical converts from StandardForm to LinearProgram.
func FromStandard[T constraints.Integer](form StandardForm[T]) LinearProgram[T] {
	noConstraints := form.NoConstraints()
	noVariables := form.NoVariables()
	tab := linalg.NewTableau[T](noConstraints+1, noVariables+2)

	tab[0][0] = 1
	copy(tab[0][1:], form.c)

	for rowCt, rowIt := range form.a {
		copy(tab[1+rowCt][1:], rowIt)
		tab[1+rowCt][1+noVariables] = form.b[rowCt]
	}

	return LinearProgram[T]{tab}
}

// NoConstraints returns the number of constraints.
func (problem LinearProgram[T]) NoConstraints() int {
	return problem.tab.NoRows() - 1
}

// NoVariables returns the number of variables.
func (problem LinearProgram[T]) NoVariables() int {
	return problem.tab.NoColumns() - 2
}
