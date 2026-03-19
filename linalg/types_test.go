package linalg

import (
	"errors"
	"testing"
)

func TestDot(t *testing.T) {
	tests := map[string]struct{
		first Vector[int]
		second Vector[int]
		want int
		err error
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

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res, err := v.first.Dot(v.second)

			if v.err != nil {
				if !errors.Is(err, v.err) {
					t.Errorf("%v is not %v.", err, v.err)
				}
			} else if err != nil {
				t.Errorf("%v", err)
			} else if res != v.want {
				t.Errorf("%v != %v", res, v.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := map[string]struct{
		first Matrix[int]
		second Matrix[int]
		want Matrix[int]
		err error
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

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res, err := v.first.Mul(v.second)

			if v.err != nil {
				if !errors.Is(err, v.err) {
					t.Errorf("%v is not %v.", err, v.err)
				}
			} else if err != nil {
				t.Errorf("%v", err)
			} else if !equalMatrices(res, v.want) {
				t.Errorf("%v != %v", res, v.want)
			}
		})
	}
}

func TestMulVec(t *testing.T) {
	tests := map[string]struct{
		mat Matrix[int]
		vec Vector[int]
		want Vector[int]
		err error
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

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res, err := v.mat.MulVec(v.vec)

			if v.err != nil {
				if !errors.Is(err, v.err) {
					t.Errorf("%v is not %v.", err, v.err)
				}
			} else if err != nil {
				t.Errorf("%v", err)
			} else if !equalVectors(res, v.want) {
				t.Errorf("%v != %v", res, v.want)
			}
		})
	}
}

func equalVectors(first, second Vector[int]) bool {
	if first.Size() != second.Size() {
		return false
	}

	for i := range first {
		if first[i] != second[i] {
			return false
		}
	}

	return true
}

func equalMatrices(first, second Matrix[int]) bool {
	if first.NoRows() != second.NoRows() {
		return false
	}

	if first.NoColumns() != second.NoColumns() {
		return false
	}

	for i := range first.NoRows() {
		for j := range first.NoColumns() {
			if first[i][j] != second[i][j] {
				return false
			}
		}
	}

	return true
}
