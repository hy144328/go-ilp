package numbers

import (
	"testing"
)

func TestAbs(t *testing.T) {
	tests := map[string]struct {
		got  int
		want int
	}{
		"positive": {1, 1},
		"zero":     {0, 0},
		"negative": {-1, 1},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res := Abs(v.got)
			if res != v.want {
				t.Errorf("Abs(%d) = %d != %d", v.got, v.want, res)
			}
		})
	}
}
func TestRemainder(t *testing.T) {
	tests := map[string]struct {
		dividend int
		divisor  int
		want     int
	}{
		"positive, positive": {5, 3, 2},
		"positive, negative": {5, -3, 2},
		"negative, positive": {-5, 3, 1},
		"negative, negative": {-5, -3, 1},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			res := Remainder(v.dividend, v.divisor)
			if res != v.want {
				t.Errorf("Remainder(%d, %d) = %d != %d", v.dividend, v.divisor, v.want, res)
			}
		})
	}
}
