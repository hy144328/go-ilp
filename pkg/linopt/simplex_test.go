package linopt

import (
	"testing"

	"github.com/hy144328/go-ilp/pkg/linalg"
)

func TestRunSimplex(t *testing.T) {
	form := CanonicalForm[int]{
		A: linalg.Matrix[int]{
			{3, 2, 1},
			{2, 5, 3},
		},
		B: linalg.Vector[int]{10, 15},
		C: linalg.Vector[int]{2, 3, 4},
	}
	lp, err := FromStandardForm(form.ToStandard(), []int{3, 4})
	if err != nil {
		panic(err)
	}

	if err := RunSimplex(lp); err != nil {
		panic(err)
	}

	res := lp.Tab[0][1+lp.NoVariables()] / lp.Tab[0][0]
	if res != 20 {
		t.Errorf("Got: %v. Expected: %v.", res, 20)
	}
}
