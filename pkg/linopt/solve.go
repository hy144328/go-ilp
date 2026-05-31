package linopt

import (
	"errors"

	"github.com/hy144328/go-ilp/pkg/linalg"
	"golang.org/x/exp/constraints"
)

var (
	ErrNotInteger = errors.New("not integer")
)

// A Solution summarizes the outcome of a LinearProgram.
type Solution[T constraints.Integer] struct {
	X linalg.Vector[T]
	Base []int
	Residual T
}

// SolveSimplex runs the simplex algorithm from scratch.
func SolveSimplex[T constraints.Signed](
	form StandardForm[T],
) (Solution[T], error) {
	lp, err := preprocess(form)
	if err != nil {
		return *new(Solution[T]), err
	}

	if err := phaseOne(lp); err != nil {
		return *new(Solution[T]), err
	}

	return phaseTwo(lp)
}

func preprocess[T constraints.Signed](
	form StandardForm[T],
) (LinearProgram[T], error) {
	lp, err := FromStandardForm(form, nil)
	if err != nil {
		return lp, err
	}

	lse := linalg.LinearSystemOfEquations[T]{
		Tab: lp.Tab.Slice(1, 0, 1, 0),
	}

	pivots, err := lse.Reduce()
	if err != nil {
		return lp, err
	}

	lp.Base = pivots
	if err := lp.Conform(); err != nil {
		return lp, err
	}

	return lp, nil
}

func phaseOne[T constraints.Signed](
	lp LinearProgram[T],
) error {
	noVariables := lp.NoVariables()
	noBasic := len(lp.Base)

	fac := lp.Tab[0][0]
	weights := linalg.FromRow(lp.weights(), 0)
	val := lp.Tab[0][noVariables+1]

	slacken(lp)

	lp.Tab[0][0] = 1
	for varCt := range noVariables {
		lp.Tab[0][1+varCt] = 0
	}
	for baseCt := range noBasic {
		lp.Tab[0][1+noVariables+baseCt] = 1
		lp.Base[baseCt] = noVariables + baseCt
	}
	lp.Tab[0][1+noVariables+noBasic] = 0

	if err := lp.Conform(); err != nil {
		return err
	}

	if err := RunSimplex(lp); err != nil {
		return err
	}

	tighten(lp)

	lp.Tab[0][0] = fac
	copy(lp.Tab[0][1:], weights)
	lp.Tab[0][1+noVariables] = val

	if err := lp.Conform(); err != nil {
		return err
	}

	return nil
}

func slacken[T constraints.Signed](lp LinearProgram[T]) {
	noVariables := lp.NoVariables()
	noBasic := len(lp.Base)

	lp.Tab[0] = append(lp.Tab[0], make([]T, noBasic)...)
	lp.Tab[0][1+noVariables+noBasic] = lp.Tab[0][1+noVariables]
	lp.Tab[0][1+noVariables] = 0

	for conCt, conIt := range lp.Tab[1:] {
		conIt = append(conIt, make([]T, noBasic)...)

		conIt[1+noVariables+noBasic] = conIt[1+noVariables]
		conIt[1+noVariables] = 0
		if conIt[1+noVariables+noBasic] >= 0 {
			conIt[1+noVariables+conCt] = 1
		} else {
			conIt[1+noVariables+conCt] = -1
		}

		lp.Tab[1+conCt] = conIt
	}
}

func tighten[T constraints.Signed](lp LinearProgram[T]) {
	noVariables := lp.NoVariables()
	noBasic := len(lp.Base)

	for conCt, conIt := range lp.Tab {
		conIt[1+noVariables-noBasic] = conIt[1+noVariables]
		lp.Tab[conCt] = conIt[:2+noVariables-noBasic]
	}
}

func phaseTwo[T constraints.Signed](
	lp LinearProgram[T],
) (Solution[T], error) {
	if err := RunSimplex(lp); err != nil {
		return *new(Solution[T]), err
	}

	noVariables := lp.NoVariables()
	bIdx := noVariables + 1

	x := make([]T, noVariables)
	for conCt, varCt := range lp.Base {
		rowCt := conCt + 1
		colCt := varCt + 1

		if lp.Tab[rowCt][bIdx] % lp.Tab[rowCt][colCt] != 0 {
			return *new(Solution[T]), ErrNotInteger
		}

		x[varCt] = lp.Tab[rowCt][bIdx] / lp.Tab[rowCt][colCt]
	}

	return Solution[T]{
		X: x,
		Base: lp.Base,
		Residual: lp.Tab[0][bIdx] / lp.Tab[0][0],
	}, nil
}
