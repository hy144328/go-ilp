package linopt

import (
	"testing"

	"github.com/hy144328/go-ilp/pkg/linalg"
)

func TestSolveSimplex(t *testing.T) {
	form := CanonicalForm[int]{
		A: linalg.Matrix[int]{
			{3, 2, 1},
			{2, 5, 3},
		},
		B: linalg.Vector[int]{10, 15},
		C: linalg.Vector[int]{2, 3, 4},
	}

	sol, err := SolveSimplex(form.ToStandard())
	if err != nil {
		panic(err)
	}

	if sol.Residual != 20 {
		t.Errorf("Got: %v. Expected: %v.", sol.Residual, 20)
	}
}
