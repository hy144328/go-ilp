package linalg

import (
	"errors"
	"testing"
)

func TestValidate(t *testing.T) {
	tests := map[string]struct{
		form LinearForm[int]
		sol Vector[int]
		err error
	}{
		"solution": {
			LinearForm[int]{
				A: Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: Vector[int]{17, 39},
			},
			Vector[int]{5, 6},
			nil,
		},
		"no solution": {
			LinearForm[int]{
				A: Matrix[int]{{1, 2}, {3, 4}},
				B: Vector[int]{17, 39},
			},
			Vector[int]{1, 1},
			ErrNotSolution,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			err := testIt.form.Validate(testIt.sol)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			}
		})
	}
}
