package linalg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/hy144328/go-ilp/numbers"
	"golang.org/x/exp/constraints"
)

var (
	ErrNotEquivalent = errors.New("not equivalent")
	ErrZeroPivot = errors.New("zero pivot")
)

type Tableau[T constraints.Integer] [][]T

// NewTableau creates a Tableau with given numbers of rows and columns.
func NewTableau[T constraints.Integer](noRows int, noColumns int) Tableau[T] {
	res := make([][]T, noRows)

	for rowCt := range noRows {
		res[rowCt] = make([]T, noColumns)
	}

	return res
}

// NoRows returns the number of rows of the Tableau.
func (tab Tableau[T]) NoRows() int {
	return len(tab)
}

// NoColumns returns the number of columns of the Tableau.
func (tab Tableau[T]) NoColumns() int {
	return len(tab[0])
}

// ScaleRow multiplies a row by a constant factor.
func (tab Tableau[T]) ScaleRow(idx int, fac T) error {
	if fac == 0 {
		return fmt.Errorf("%w: Factor must not be zero.", ErrNotEquivalent)
	}

	row := tab[idx]

	for colCt := range row {
		row[colCt] *= fac
	}

	return nil
}

// EliminateRow subtracts the multiple of one row from the multiple of another row such that the corresponding entry in the column vanishes.
func (tab Tableau[T]) EliminateRow(srcIdx int, dstIdx int, colIdx int) error {
	src := tab[srcIdx]
	dst := tab[dstIdx]

	facSrc := dst[colIdx]
	facDst := src[colIdx]

	if facSrc == 0 {
		return nil
	} else if facDst == 0 {
		return fmt.Errorf("%w: tab[%d, %d] = 0.", ErrZeroPivot, srcIdx, colIdx)
	}

	for colCt, colIt := range src {
		dst[colCt] *= facDst
		dst[colCt] -= facSrc * colIt
	}

	return nil
}

// DeflateRow divides a row by a constant factor such that all entries are coprime.
func (tab Tableau[T]) DeflateRow(idx int) {
	row := tab[idx]

	fac := numbers.GreatestCommonDivisor(row[0], row[1:]...)
	if fac == 0 {
		return
	}

	for colCt := range row {
		row[colCt] /= fac
	}
}

// SwapRows exchanges one row with another row.
func (tab Tableau[T]) SwapRows(srcIdx int, dstIdx int) {
	tab[srcIdx], tab[dstIdx] = tab[dstIdx], tab[srcIdx]
}

// Equals reports whether a Tableau is equal to another Tableau.
func (tab Tableau[T]) Equals(other Tableau[T]) bool {
	if len(tab) != len(other) {
		return false
	}
	if len(tab[0]) != len(other[0]) {
		return false
	}

	for rowCt := range tab {
		for colCt := range tab[0] {
			if tab[rowCt][colCt] != other[rowCt][colCt] {
				return false
			}
		}
	}

	return true
}

// String formats a Tableau as a string.
func (tab Tableau[T]) String() string {
	as := make([]string, tab.NoRows())

	for rowCt, rowIt := range tab {
		asIt := make([]string, tab.NoColumns())

		for colCt, colIt := range rowIt {
			asIt[colCt] = strconv.Itoa(int(colIt))
		}

		as[rowCt] = strings.Join(asIt, "\t")
	}

	return strings.Join(as, "\n")
}
