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

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			v.tab.ScaleRow(v.idx, v.fac)

			if !equalTableaus(v.tab, v.want) {
				t.Errorf("%v != %v", v.tab, v.want)
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

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			v.tab.EliminateRow(v.srcIdx, v.dstIdx, v.colIdx)

			if !equalTableaus(v.tab, v.want) {
				t.Errorf("%v != %v", v.tab, v.want)
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

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			v.tab.DeflateRow(v.idx)

			if !equalTableaus(v.tab, v.want) {
				t.Errorf("%v != %v", v.tab, v.want)
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

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			v.tab.SwapRows(v.srcIdx, v.dstIdx)

			if !equalTableaus(v.tab, v.want) {
				t.Errorf("%v != %v", v.tab, v.want)
			}
		})
	}
}

func equalTableaus(
	tab Tableau[int],
	other Tableau[int],
) bool {
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
