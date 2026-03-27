package linalg

import (
	"testing"
)

func TestScaleRow(t *testing.T) {
	tests := map[string]struct {
		tab  Tableau[int]
		idx  int
		fac  int
		want Tableau[int]
	}{
		"base": {[][]int{{1, 2}, {3, 4}}, 0, 2, [][]int{{2, 4}, {3, 4}}},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			testIt.tab.ScaleRow(testIt.idx, testIt.fac)

			if !testIt.tab.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", testIt.tab, testIt.want)
			}
		})
	}
}

func TestEliminateRow(t *testing.T) {
	tests := map[string]struct {
		tab    Tableau[int]
		srcIdx int
		dstIdx int
		colIdx int
		want   Tableau[int]
	}{
		"base": {[][]int{{1, 2}, {3, 4}}, 0, 1, 0, [][]int{{1, 2}, {0, -2}}},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			testIt.tab.EliminateRow(testIt.srcIdx, testIt.dstIdx, testIt.colIdx)

			if !testIt.tab.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", testIt.tab, testIt.want)
			}
		})
	}
}

func TestDeflateRow(t *testing.T) {
	tests := map[string]struct {
		tab  Tableau[int]
		idx  int
		want Tableau[int]
	}{
		"base": {[][]int{{2, 4}, {3, 4}}, 0, [][]int{{1, 2}, {3, 4}}},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			testIt.tab.DeflateRow(testIt.idx)

			if !testIt.tab.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", testIt.tab, testIt.want)
			}
		})
	}
}

func TestSwapRows(t *testing.T) {
	tests := map[string]struct {
		tab    Tableau[int]
		srcIdx int
		dstIdx int
		want   Tableau[int]
	}{
		"base": {[][]int{{1, 2}, {3, 4}}, 0, 1, [][]int{{3, 4}, {1, 2}}},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			testIt.tab.SwapRows(testIt.srcIdx, testIt.dstIdx)

			if !testIt.tab.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", testIt.tab, testIt.want)
			}
		})
	}
}
