package linopt

import (
	"errors"

	"github.com/hy144328/go-ilp/pkg/linalg"
	"golang.org/x/exp/constraints"
)

var (
	ErrUnboundedAbove = errors.New("unbounded above")
)

// RunSimplex runs the simplex algorithm given a basic feasible solution.
func RunSimplex[T constraints.Signed](lp LinearProgram[T]) error {
	for {
		varCt, ok := chooseEntering(lp)
		if !ok {
			break
		}
		colCt := varCt + 1

		conCt, ok := chooseExiting(lp, varCt)
		if !ok {
			return ErrUnboundedAbove
		}
		rowCt := conCt + 1

		if err := linalg.EliminateUp(lp.Tab, rowCt, colCt); err != nil {
			panic(err)
		}
		if err := linalg.EliminateDown(lp.Tab, rowCt, colCt); err != nil {
			panic(err)
		}
		lp.Base[conCt] = varCt
	}

	return nil
}

// chooseEntering chooses the entering variable based on Bland's rule.
// If no variable is found, this means that the linear program is finished.
func chooseEntering[T constraints.Signed](
	lp LinearProgram[T],
) (int, bool) {
	for varCt, varIt := range lp.Tab[0][1:1+lp.NoVariables()] {
		if varIt < 0 {
			return varCt, true
		}
	}

	return 0, false
}

// chooseExiting chooses the exiting variable based on the minimum-ratio criterion.
// If no variable is found, this means that the linear program is unbounded above.
func chooseExiting[T constraints.Signed](
	lp LinearProgram[T],
	varIdx int,
) (int, bool) {
	colCt := varIdx + 1

	var res int
	var ok bool

	for conCt := range lp.NoConstraints() {
		rowCt := conCt + 1

		if lp.Tab[rowCt][colCt] <= 0 {
			continue
		}

		if !ok {
			res = conCt
			ok = true
			continue
		}

		if checkMinimumRatio(lp, varIdx, res, conCt) {
			res = conCt
			ok = true
		}
	}

	return res, ok
}

// checkMinimumRatio evaluates the minimum-ratio criterion.
func checkMinimumRatio[T constraints.Signed](
	lp LinearProgram[T],
	varIdx int,
	conIdx int,
	newIdx int,
) bool {
	bIdx := lp.NoVariables() + 1
	lhs := lp.Tab[newIdx+1][bIdx] * lp.Tab[conIdx+1][varIdx+1]
	rhs := lp.Tab[conIdx+1][bIdx] * lp.Tab[newIdx+1][varIdx+1]

	return lhs < rhs
}
