package linopt

import (
	"testing"

	"github.com/hy144328/go-ilp/linalg"
)

func TestFromCanonical(t *testing.T) {
	tests := map[string]struct{
		got CanonicalForm[int]
		want LinearProgram[int]
	}{
		"base": {
			got: CanonicalForm[int]{
				a: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				b: linalg.Vector[int]{5, 6},
				c: linalg.Vector[int]{1, 1},
			},
			want: LinearProgram[int]{
				tab: linalg.Tableau[int]{
					{1, 1, 1, 0, 0, 0},
					{0, 1, 2, 1, 0, 5},
					{0, 3, 4, 0, 1, 6},
				},
			},
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			prob := FromCanonical(testIt.got)
			if !prob.tab.Equals(testIt.want.tab) {
				t.Errorf("got != want.\n\ngot:\n%v\n\nwant:\n%v\n", prob, testIt.want.tab)
			}
		})
	}
}

func TestFromStandard(t *testing.T) {
	tests := map[string]struct{
		got StandardForm[int]
		want LinearProgram[int]
	}{
		"base": {
			got: StandardForm[int]{
				a: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				b: linalg.Vector[int]{5, 6},
				c: linalg.Vector[int]{1, 1},
			},
			want: LinearProgram[int]{
				tab: linalg.Tableau[int]{
					{1, 1, 1, 0},
					{0, 1, 2, 5},
					{0, 3, 4, 6},
				},
			},
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			prob := FromStandard(testIt.got)
			if !prob.tab.Equals(testIt.want.tab) {
				t.Errorf("got != want.\n\ngot:\n%v\n\nwant:\n%v\n", prob, testIt.want.tab)
			}
		})
	}
}
