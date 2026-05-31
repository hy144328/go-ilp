package linopt

import (
	"errors"
	"fmt"

	"github.com/hy144328/go-ilp/pkg/linalg"
	"golang.org/x/exp/constraints"
)

var (
	ErrNotFeasible = errors.New("not feasible")
)

// A LinearProgram represents a linear program in standard formulation.
type LinearProgram[T constraints.Signed] struct {
	Tab linalg.Tableau[T]
	Base []int
}

// FromStandardForm converts from StandardForm to LinearProgram.
func FromStandardForm[T constraints.Signed](
	form StandardForm[T],
	base []int,
) (LinearProgram[T], error) {
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
	res.Base = base

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
func (lp LinearProgram[T]) ToStandardForm() (StandardForm[T], []int) {
	return StandardForm[T]{
		A: lp.leftHandSide().Copy(),
		B: linalg.FromColumn(lp.rightHandSide(), 0),
		C: linalg.FromRow(lp.weights(), 0),
	}, lp.Base
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

// Conform transforms the tableau according to the basis of the initial feasible solution.
// Returns error if basic solution turns out infeasible.
func (lp LinearProgram[T]) Conform() error {
	bIdx := lp.NoVariables() + 1

	for conCt, varCt := range lp.Base {
		rowCt := conCt + 1
		colCt := varCt + 1

		linalg.PivotColumn(lp.Tab, rowCt, colCt)

		if err := linalg.EliminateDown(lp.Tab, rowCt, colCt); err != nil {
			return fmt.Errorf("%w: %w", ErrNotFeasible, err)
		}
	}

	for conCt := len(lp.Base) - 1; conCt >= 0; conCt-- {
		varCt := lp.Base[conCt]
		rowCt := conCt + 1
		colCt := varCt + 1

		if err := linalg.EliminateUp(lp.Tab, rowCt, colCt); err != nil {
			return fmt.Errorf("%w: %w", ErrNotFeasible, err)
		}

		if lp.Tab[rowCt][colCt] * lp.Tab[rowCt][bIdx] < 0 {
			return fmt.Errorf("%w: Not positive.", ErrNotFeasible)
		}

		if lp.Tab[rowCt][bIdx] < 0 {
			if err := lp.Tab.ScaleRow(rowCt, -1); err != nil {
				panic(err)
			}
		}
	}

	return nil
}
