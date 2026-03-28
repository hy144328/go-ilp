package linopt

import (
	"errors"
	"testing"

	"github.com/hy144328/go-ilp/pkg/linalg"
)

func TestFromStandardForm(t *testing.T) {
	tests := map[string]struct{
		got StandardForm[int]
		want LinearProgram[int]
		err error
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
				tab: linalg.Tableau[int]{
					{1, -1, -1, 0},
					{0, 1, 2, 5},
					{0, -3, -4, 6},
				},
			},
			err: nil,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			prob, err := FromStandardForm(testIt.got)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !prob.tab.Equals(testIt.want.tab) {
				t.Errorf("got != want.\n\ngot:\n%v\n\nwant:\n%v\n", prob.tab, testIt.want.tab)
			}
		})
	}
}

func TestToStandardForm(t *testing.T) {
	tests := map[string]struct{
		got LinearProgram[int]
		want StandardForm[int]
	}{
		"base": {
			got: LinearProgram[int]{
				tab: linalg.Tableau[int]{
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
			form := testIt.got.ToStandardForm()

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
