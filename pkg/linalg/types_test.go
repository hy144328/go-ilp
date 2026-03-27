package linalg

import (
	"errors"
	"testing"
)

func TestDot(t *testing.T) {
	tests := map[string]struct {
		first  Vector[int]
		second Vector[int]
		want   int
		err    error
	}{
		"compatible": {
			[]int{1, 2, 3},
			[]int{4, 5, 6},
			32,
			nil,
		},
		"incompatible": {
			[]int{1, 2, 3},
			[]int{4, 5},
			0,
			ErrIncompatibleSizes,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			res, err := testIt.first.Dot(testIt.second)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if res != testIt.want {
				t.Errorf("%v != %v", res, testIt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := map[string]struct {
		first  Matrix[int]
		second Matrix[int]
		want   Matrix[int]
		err    error
	}{
		"compatible": {
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[][]int{{7}, {8}, {9}},
			[][]int{{50}, {122}},
			nil,
		},
		"incompatible": {
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[][]int{{7}, {8}},
			[][]int{},
			ErrIncompatibleSizes,
		},
		"wrong orientation": {
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[][]int{{7, 8, 9}},
			[][]int{},
			ErrIncompatibleSizes,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			res, err := testIt.first.Mul(testIt.second)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !res.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", res, testIt.want)
			}
		})
	}
}

func TestMulVec(t *testing.T) {
	tests := map[string]struct {
		mat  Matrix[int]
		vec  Vector[int]
		want Vector[int]
		err  error
	}{
		"compatible": {
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[]int{7, 8, 9},
			[]int{50, 122},
			nil,
		},
		"incompatible": {
			[][]int{{1, 2, 3}, {4, 5, 6}},
			[]int{7, 8},
			[]int{},
			ErrIncompatibleSizes,
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			res, err := testIt.mat.MulVec(testIt.vec)

			if testIt.err != nil {
				if !errors.Is(err, testIt.err) {
					t.Errorf("%v is not %v.", err, testIt.err)
				}
			} else if err != nil {
				t.Error(err)
			} else if !res.Equals(testIt.want) {
				t.Errorf("got != want\n\ngot:\n%v\n\nwant:\n%v\n", res, testIt.want)
			}
		})
	}
}
