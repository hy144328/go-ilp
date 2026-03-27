package linopt

import (
	"errors"
	"testing"

	"github.com/hy144328/go-ilp/linalg"
)

func TestToStandard(t *testing.T) {
	tests := map[string]struct{
		got CanonicalForm[int]
		want StandardForm[int]
	}{
		"base": {
			got: CanonicalForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{5, 6},
				C: linalg.Vector[int]{1, 1},
			},
			want: StandardForm[int]{
				A: linalg.Matrix[int]{
					{1, 2, 1, 0},
					{3, 4, 0, 1},
				},
				B: linalg.Vector[int]{5, 6},
				C: linalg.Vector[int]{1, 1, 0, 0},
			},
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			form := testIt.got.ToStandard()

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

func TestValidateCanonical(t *testing.T) {
	tests := map[string]struct{
		form CanonicalForm[int]
		sol linalg.Vector[int]
		err error
	}{
		"exact solution": {
			CanonicalForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{17, 39},
				C: linalg.Vector[int]{1, 1},
			},
			linalg.Vector[int]{5, 6},
			nil,
		},
		"valid solution": {
			CanonicalForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{17, 39},
				C: linalg.Vector[int]{1, 1},
			},
			linalg.Vector[int]{1, 1},
			nil,
		},
		"not solution": {
			CanonicalForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{17, 39},
				C: linalg.Vector[int]{1, 1},
			},
			linalg.Vector[int]{10, 10},
			linalg.ErrNotSolution,
		},
		"negative solution": {
			CanonicalForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{17, 39},
				C: linalg.Vector[int]{1, 1},
			},
			linalg.Vector[int]{-1, -1},
			ErrNegativeSolution,
		},
	}

	for k, testIt := range tests {
		t.Run(k, func(t *testing.T) {
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

func TestValidateStandard(t *testing.T) {
	tests := map[string]struct{
		form StandardForm[int]
		sol linalg.Vector[int]
		err error
	}{
		"solution": {
			StandardForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{17, 39},
				C: linalg.Vector[int]{1, 1},
			},
			linalg.Vector[int]{5, 6},
			nil,
		},
		"not solution": {
			StandardForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{17, 39},
				C: linalg.Vector[int]{1, 1},
			},
			linalg.Vector[int]{10, 10},
			linalg.ErrNotSolution,
		},
		"negative solution": {
			StandardForm[int]{
				A: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				B: linalg.Vector[int]{17, 39},
				C: linalg.Vector[int]{1, 1},
			},
			linalg.Vector[int]{-1, -1},
			ErrNegativeSolution,
		},
	}

	for k, testIt := range tests {
		t.Run(k, func(t *testing.T) {
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
