package linopt

import (
	"errors"
	"testing"

	"github.com/hy144328/go-ilp/linalg"
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
				B: linalg.Vector[int]{5, 6},
				C: linalg.Vector[int]{1, 1},
			},
			want: LinearProgram[int]{
				tab: linalg.Tableau[int]{
					{-1, 1, 1, 0},
					{0, 1, 2, 5},
					{0, 3, 4, 6},
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
				t.Errorf("got != want.\n\ngot:\n%v\n\nwant:\n%v\n", prob, testIt.want.tab)
			}
		})
	}
}
