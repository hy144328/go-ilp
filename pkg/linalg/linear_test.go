package linalg

import (
	"errors"
	"slices"
	"testing"
)

func TestFromLinearForm(t *testing.T) {
	tests := map[string]struct{
		got LinearForm[int]
		want LinearSystemOfEquations[int]
		err error
	}{
		"base": {
			got: LinearForm[int]{
				A: Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: Vector[int]{5, 6},
			},
			want: LinearSystemOfEquations[int]{
				tab: Tableau[int]{
					{1, 2, 5},
					{3, 4, 6},
				},
			},
			err: nil,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			prob, err := FromLinearForm(testIt.got)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !prob.tab.Equals(testIt.want.tab) {
				t.Errorf("got != want.\n\ngot:\n%v\n\nwant:\n%v\n", prob, testIt.want.tab)
			}
		})
	}
}

func TestToLinearForm(t *testing.T) {
	tests := map[string]struct{
		got LinearSystemOfEquations[int]
		want LinearForm[int]
	}{
		"base": {
			got: LinearSystemOfEquations[int]{
				tab: Tableau[int]{
					{1, 2, 5},
					{3, 4, 6},
				},
			},
			want: LinearForm[int]{
				A: Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: Vector[int]{5, 6},
			},
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			form := testIt.got.ToLinearForm()

			if !form.A.Equals(testIt.want.A) {
				t.Errorf("a != a.\n\ngot:\n%v\n\nwant:\n%v\n", form.A, testIt.want.A)
			}

			if !form.B.Equals(testIt.want.B) {
				t.Errorf("b != b.\n\ngot:\n%v\n\nwant:\n%v\n", form.B, testIt.want.B)
			}
		})
	}
}

func TestReduce(t *testing.T) {
	tests := map[string]struct{
		lse LinearSystemOfEquations[int]
		pivots []int
		err error
	}{
		"regular": {LinearSystemOfEquations[int]{Tableau[int]{{1, 2, 3, 4}, {5, 6, 7, 8}}}, []int{0, 1}, nil},
		"skip": {LinearSystemOfEquations[int]{Tableau[int]{{1, 2, 3, 4}, {3, 6, 7, 8}}}, []int{0, 2}, nil},
		"irregular": {LinearSystemOfEquations[int]{Tableau[int]{{1, 2, 3, 4}, {2, 4, 6, 8}}}, []int{0}, nil},
		"irreducible": {LinearSystemOfEquations[int]{Tableau[int]{{1, 2, 3, 4}, {2, 4, 6, 7}}}, []int{0}, ErrNoSolution},
	}

	for k, testIt := range tests {
		t.Run(k, func(t *testing.T) {
			pivots, err := testIt.lse.Reduce()

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !slices.Equal(pivots, testIt.pivots) {
				t.Errorf("%v != %v.", pivots, testIt.pivots)
			}
		})
	}
}
