package linalg

import (
	"errors"
	"testing"
)

func TestPivotColumn(t *testing.T) {
	tests := map[string]struct{
		tab Tableau[int]
		rowIdx int
		colIdx int
		want Tableau[int]
	}{
		"swap": {Tableau[int]{{1, 2}, {3, 4}}, 0, 0, Tableau[int]{{3, 4}, {1, 2}}},
		"no swap": {Tableau[int]{{3, 4}, {1, 2}}, 0, 0, Tableau[int]{{3, 4}, {1, 2}}},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			PivotColumn(testIt.tab, testIt.rowIdx, testIt.colIdx)
			if !testIt.tab.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", testIt.tab, testIt.want)
			}
		})
	}
}

func TestEliminateDown(t *testing.T) {
	tests := map[string]struct{
		tab Tableau[int]
		rowIdx int
		colIdx int
		want Tableau[int]
		err error
	}{
		"eliminate": {Tableau[int]{{1, 2, 3}, {4, 5, 6}}, 0, 0, Tableau[int]{{1, 2, 3}, {0, -1, -2}}, nil},
		"no eliminate": {Tableau[int]{{1, 2, 3}, {0, 5, 6}}, 0, 0, Tableau[int]{{1, 2, 3}, {0, 5, 6}}, nil},
		"no pivot": {Tableau[int]{{0, 2, 3}, {4, 5, 6}}, 0, 0, Tableau[int]{{0, 2, 3}, {4, 5, 6}}, ErrZeroPivot},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			err := EliminateDown(testIt.tab, testIt.rowIdx, testIt.colIdx)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !testIt.tab.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", testIt.tab, testIt.want)
			}
		})
	}
}

func TestEliminateUp(t *testing.T) {
	tests := map[string]struct{
		tab Tableau[int]
		rowIdx int
		colIdx int
		want Tableau[int]
		err error
	}{
		"eliminate": {Tableau[int]{{1, 2, 3}, {4, 5, 6}}, 1, 0, Tableau[int]{{0, 1, 2}, {4, 5, 6}}, nil},
		"no eliminate": {Tableau[int]{{0, 2, 3}, {4, 5, 6}}, 1, 0, Tableau[int]{{0, 2, 3}, {4, 5, 6}}, nil},
		"no pivot": {Tableau[int]{{1, 2, 3}, {0, 5, 6}}, 1, 0, Tableau[int]{{0, 2, 3}, {4, 5, 6}}, ErrZeroPivot},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			err := EliminateUp(testIt.tab, testIt.rowIdx, testIt.colIdx)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !testIt.tab.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", testIt.tab, testIt.want)
			}
		})
	}
}
