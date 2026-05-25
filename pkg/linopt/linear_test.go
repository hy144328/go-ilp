package linopt

import (
	"errors"
	"testing"

	"github.com/hy144328/go-ilp/pkg/linalg"
)

func TestFromStandardForm(t *testing.T) {
	tests := map[string]struct {
		got  StandardForm[int]
		want LinearProgram[int]
		err  error
	}{
		"base": {
			got: StandardForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{5, -6},
				C: linalg.Vector[int]{1, 1},
			},
			want: LinearProgram[int]{
				Tab: linalg.Tableau[int]{
					{1, -1, -1, 0},
					{0, 1, 2, 5},
					{0, 3, 4, -6},
				},
			},
			err: nil,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			prob, err := FromStandardForm(testIt.got, nil)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !prob.Tab.Equals(testIt.want.Tab) {
				t.Errorf("got != want.\n\ngot:\n%v\n\nwant:\n%v\n", prob.Tab, testIt.want.Tab)
			}
		})
	}
}

func TestToStandardForm(t *testing.T) {
	tests := map[string]struct {
		got  LinearProgram[int]
		want StandardForm[int]
	}{
		"base": {
			got: LinearProgram[int]{
				Tab: linalg.Tableau[int]{
					{-1, 1, 1, 0},
					{0, 1, 2, 5},
					{0, 3, 4, 6},
				},
			},
			want: StandardForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{5, 6},
				C: linalg.Vector[int]{1, 1},
			},
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			form, _ := testIt.got.ToStandardForm()

			if !form.A.Equals(testIt.want.A) {
				t.Errorf("a != a.\n\ngot:\n%v\n\nwant:\n%v\n", form.A, testIt.want.A)
			}

			if !form.B.Equals(testIt.want.B) {
				t.Errorf("b != b.\n\ngot:\n%v\n\nwant:\n%v\n", form.B, testIt.want.B)
			}

			if !form.C.Equals(testIt.want.C) {
				t.Errorf("c != c.\n\ngot:\n%v\n\nwant:\n%v\n", form.C, testIt.want.C)
			}
		})
	}
}

/*
func TestReduce(t *testing.T) {
	tests := map[string]struct {
		lp     LinearProgram[int]
		pivots []int
		err    error
	}{
		"regular": {
			LinearProgram[int]{
				linalg.Tableau[int]{
					{1, -1, -1, -1, -1},
					{0, 1, 2, 3, 4},
					{0, 5, 6, 7, 8},
				},
			},
			[]int{0, 1},
			nil,
		},
		"skip": {
			LinearProgram[int]{
				linalg.Tableau[int]{
					{1, -1, -1, -1, -1},
					{0, 1, 2, 3, 4},
					{0, 3, 6, 7, 8},
				},
			},
			[]int{0, 2},
			nil,
		},
		"irregular": {
			LinearProgram[int]{
				linalg.Tableau[int]{
					{1, -1, -1, -1, -1},
					{0, 1, 2, 3, 4},
					{0, 2, 4, 6, 8},
				},
			},
			[]int{0},
			nil,
		},
		"irreducible": {
			LinearProgram[int]{
				linalg.Tableau[int]{
					{1, -1, -1, -1, -1},
					{0, 1, 2, 3, 4},
					{0, 2, 4, 6, 7},
				},
			},
			[]int{0},
			linalg.ErrNoSolution,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			pivots, err := testIt.lp.Reduce()

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !slices.Equal(pivots, testIt.pivots) {
				t.Errorf("%v != %v.", pivots, testIt.pivots)
			} else {
				for i := 1 + len(pivots); i < testIt.lp.Tab.NoRows(); i++ {
					for j := range testIt.lp.Tab.NoColumns() {
						if testIt.lp.Tab[i][j] != 0 {
							t.Errorf("tab[%d][%d] != 0.", i, j)
						}
					}
				}
			}
		})
	}
}
*/
