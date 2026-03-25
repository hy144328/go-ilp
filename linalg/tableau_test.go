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

			if !v.tab.Equals(v.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", v.tab, v.want)
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

			if !v.tab.Equals(v.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", v.tab, v.want)
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

			if !v.tab.Equals(v.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", v.tab, v.want)
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

			if !v.tab.Equals(v.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", v.tab, v.want)
			}
		})
	}
}
