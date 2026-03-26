package linalg

import (
	"github.com/hy144328/go-ilp/numbers"
	"golang.org/x/exp/constraints"
)

// PivotColumn swaps rows such that the absolute value of the entry is larger or equal to any entry below.
func PivotColumn[T constraints.Integer](
	tab Tableau[T],
	rowIdx int,
	colIdx int,
) {
	idx := rowIdx
	val := numbers.Abs(tab[rowIdx][colIdx])

	for rowCt := rowIdx + 1; rowCt < tab.NoRows(); rowCt++ {
		if valIt := numbers.Abs(tab[rowCt][colIdx]); val < valIt {
			idx = rowCt
			val = valIt
		}
	}

	tab.SwapRows(rowIdx, idx)
}

// EliminateDown subtracts the given row from the rows below such that their values in the same column vanish.
// The rows are deflated to prevent overflow.
func EliminateDown[T constraints.Integer](
	tab Tableau[T],
	rowIdx int,
	colIdx int,
) error {
	for rowCt := rowIdx + 1; rowCt < tab.NoRows(); rowCt++ {
		if err := tab.EliminateRow(rowIdx, rowCt, colIdx); err != nil {
			return err
		}

		tab.DeflateRow(rowCt)
	}

	return nil
}

// EliminateUp subtracts the given row from the rows above such that their values in the same column vanish.
// The rows are deflated to prevent overflow.
func EliminateUp[T constraints.Integer](
	tab Tableau[T],
	rowIdx int,
	colIdx int,
) error {
	for rowCt := range rowIdx {
		if err := tab.EliminateRow(rowIdx, rowCt, colIdx); err != nil {
			return err
		}

		tab.DeflateRow(rowCt)
	}

	return nil
}
