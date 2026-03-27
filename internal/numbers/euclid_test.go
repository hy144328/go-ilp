package numbers

import (
	"testing"
)

func TestGreatestCommonDivisor(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"two positive":               {[]int{12, 18}, 6},
		"two positive reverse":       {[]int{18, 12}, 6},
		"three positive":             {[]int{12, 18, 21}, 3},
		"one positive, one negative": {[]int{12, -18}, 6},
		"two negative":               {[]int{-12, -18}, 6},
		"one positive":               {[]int{1}, 1},
		"one zero":                   {[]int{0}, 0},
		"one negative":               {[]int{-1}, 1},
		"one positive, one zero":     {[]int{1, 0}, 1},
		"two zero":                   {[]int{0, 0}, 0},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			res := GreatestCommonDivisor(testIt.got[0], testIt.got[1:]...)
			if res != testIt.want {
				t.Errorf("gcd(%v) = %d != %d", testIt.got, res, testIt.want)
			}
		})
	}
}

func TestLeastCommonMultiple(t *testing.T) {
	tests := map[string]struct {
		got  []int
		want int
	}{
		"two positive":               {[]int{12, 18}, 36},
		"two positive reverse":       {[]int{18, 12}, 36},
		"three positive":             {[]int{12, 18, 21}, 252},
		"one positive, one negative": {[]int{12, -18}, 36},
		"two negative":               {[]int{-12, -18}, 36},
		"one positive":               {[]int{1}, 1},
		"one zero":                   {[]int{0}, 0},
		"one negative":               {[]int{-1}, 1},
		"one positive, one zero":     {[]int{1, 0}, 0},
		"two zero":                   {[]int{0, 0}, 0},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			res := LeastCommonMultiple(testIt.got[0], testIt.got[1:]...)
			if res != testIt.want {
				t.Errorf("lcm(%v) = %d != %d", testIt.got, res, testIt.want)
			}
		})
	}
}
