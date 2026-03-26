package linalg

import (
	"errors"
	"slices"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := map[string]struct{
		lse LinearSystemOfEquations[int]
		sol Vector[int]
		err error
	}{
		"true": {LinearSystemOfEquations[int]{Tableau[int]{{1, 2, 17}, {3, 4, 39}}}, Vector[int]{5, 6}, nil},
		"false": {LinearSystemOfEquations[int]{Tableau[int]{{1, 2, 17}, {3, 4, 39}}}, Vector[int]{1, 1}, ErrNotSolution},
	}

	for k, testIt := range tests {
		t.Run(k, func(t *testing.T) {
			err := testIt.lse.Validate(testIt.sol)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Errorf("%v.", err)
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
				t.Errorf("%v.", err)
			} else if !slices.Equal(pivots, testIt.pivots) {
				t.Errorf("%v != %v.", pivots, testIt.pivots)
			}
		})
	}
}
