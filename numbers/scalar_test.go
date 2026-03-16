package numbers

import (
	"errors"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := map[string]struct{
		first Rational[int]
		second Rational[int]
		want Rational[int]
	}{
		"two numbers": {fraction[int]{1, 2}, fraction[int]{1, 3}, fraction[int]{5, 6}},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res := v.first.Add(v.second)
			if !res.Equals(v.want) {
				t.Errorf("%v + %v = %v != %v", v.first, v.second, res, v.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	tests := map[string]struct{
		first Rational[int]
		second Rational[int]
		want Rational[int]
	}{
		"two numbers": {fraction[int]{1, 2}, fraction[int]{1, 3}, fraction[int]{1, 6}},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res := v.first.Mul(v.second)
			if !res.Equals(v.want) {
				t.Errorf("%v * %v = %v != %v", v.first, v.second, res, v.want)
			}
		})
	}
}

func TestMulT(t *testing.T) {
	tests := map[string]struct{
		got Rational[int]
		fac int
		want Rational[int]
	}{
		"two numbers": {fraction[int]{5, 6}, 15, fraction[int]{25, 2}},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res := v.got.MulT(v.fac)
			if !res.Equals(v.want) {
				t.Errorf("%v * %d = %v != %v", v.got, v.fac, res, v.want)
			}
		})
	}
}

func TestEquals(t *testing.T) {
	tests := map[string]struct{
		lhs Rational[int]
		rhs Rational[int]
		want bool
	}{
		"equal": {fraction[int]{1, 3}, fraction[int]{1, 3}, true},
		"unequal": {fraction[int]{1, 3}, fraction[int]{3, 1}, false},
		"equivalent": {fraction[int]{1, 3}, fraction[int]{2, 6}, true},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if v.lhs.Equals(v.rhs) != v.want {
				if v.want {
					t.Errorf("%v != %v", v.lhs, v.rhs)
				} else {
					t.Errorf("%v == %v", v.lhs, v.rhs)
				}
			}
		})
	}
}

func TestLessThan(t *testing.T) {
	tests := map[string]struct{
		lhs Rational[int]
		rhs Rational[int]
		want bool
	}{
		"less": {fraction[int]{1, 4}, fraction[int]{1, 3}, true},
		"equal": {fraction[int]{1, 3}, fraction[int]{1, 3}, false},
		"greater": {fraction[int]{1, 2}, fraction[int]{1, 3}, false},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if v.lhs.LessThan(v.rhs) != v.want {
				if v.want {
					t.Errorf("%v >= %v", v.lhs, v.rhs)
				} else {
					t.Errorf("%v < %v", v.lhs, v.rhs)
				}
			}
		})
	}
}

func TestGreaterThan(t *testing.T) {
	tests := map[string]struct{
		lhs Rational[int]
		rhs Rational[int]
		want bool
	}{
		"less": {fraction[int]{1, 4}, fraction[int]{1, 3}, false},
		"equal": {fraction[int]{1, 3}, fraction[int]{1, 3}, false},
		"greater": {fraction[int]{1, 2}, fraction[int]{1, 3}, true},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if v.lhs.GreaterThan(v.rhs) != v.want {
				if v.want {
					t.Errorf("%v <= %v", v.lhs, v.rhs)
				} else {
					t.Errorf("%v > %v", v.lhs, v.rhs)
				}
			}
		})
	}
}

func TestLessEqual(t *testing.T) {
	tests := map[string]struct{
		lhs Rational[int]
		rhs Rational[int]
		want bool
	}{
		"less": {fraction[int]{1, 4}, fraction[int]{1, 3}, true},
		"equal": {fraction[int]{1, 3}, fraction[int]{1, 3}, true},
		"greater": {fraction[int]{1, 2}, fraction[int]{1, 3}, false},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if v.lhs.LessEqual(v.rhs) != v.want {
				if v.want {
					t.Errorf("%v > %v", v.lhs, v.rhs)
				} else {
					t.Errorf("%v <= %v", v.lhs, v.rhs)
				}
			}
		})
	}
}

func TestGreaterEqual(t *testing.T) {
	tests := map[string]struct{
		lhs Rational[int]
		rhs Rational[int]
		want bool
	}{
		"less": {fraction[int]{1, 4}, fraction[int]{1, 3}, false},
		"equal": {fraction[int]{1, 3}, fraction[int]{1, 3}, true},
		"greater": {fraction[int]{1, 2}, fraction[int]{1, 3}, true},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if v.lhs.GreaterEqual(v.rhs) != v.want {
				if v.want {
					t.Errorf("%v >= %v", v.lhs, v.rhs)
				} else {
					t.Errorf("%v < %v", v.lhs, v.rhs)
				}
			}
		})
	}
}

func TestFloor(t *testing.T) {
	tests := map[string]struct{
		got Rational[int]
		want int
	}{
		"positive fraction": {fraction[int]{5, 3}, 1},
		"negative fraction": {fraction[int]{-5, 3}, -2},
		"integer": {fraction[int]{6, 3}, 2},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if v.got.Floor() != v.want {
				t.Errorf("floor(%v) != %d", v.got, v.want)
			}
		})
	}
}

func TestCeil(t *testing.T) {
	tests := map[string]struct{
		got Rational[int]
		want int
	}{
		"positive fraction": {fraction[int]{5, 3}, 2},
		"negative fraction": {fraction[int]{-5, 3}, -1},
		"integer": {fraction[int]{6, 3}, 2},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			if v.got.Ceil() != v.want {
				t.Errorf("ceil(%v) != %d", v.got, v.want)
			}
		})
	}
}

func TestToInteger(t *testing.T) {
	tests := map[string]struct{
		got Rational[int]
		want int
		err error
	}{
		"equal": {fraction[int]{2, 1}, 2, nil},
		"unequal": {fraction[int]{2, 3}, 0, ErrNotInteger},
		"equivalent": {fraction[int]{6, 3}, 2, nil},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res, err := v.got.ToInteger()
			if v.err != nil {
				if !errors.Is(err, v.err) {
					t.Errorf("%v is not %v.", err, v.err)
				}
			} else if err != nil {
				t.Errorf("%v.", err.Error())
			} else if res != v.want {
				t.Errorf("%v != %d", v.got, v.want)
			}
		})
	}
}
