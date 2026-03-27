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

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			res := Abs(testIt.got)
			if res != testIt.want {
				t.Errorf("Abs(%d) = %d != %d", testIt.got, testIt.want, res)
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

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			res := Remainder(testIt.dividend, testIt.divisor)
			if res != testIt.want {
				t.Errorf("Remainder(%d, %d) = %d != %d", testIt.dividend, testIt.divisor, testIt.want, res)
			}
		})
	}
}
