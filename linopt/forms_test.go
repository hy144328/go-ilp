package linopt

import (
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
				a: linalg.Matrix[int]{
					{1, 2},
					{3, 4},
				},
				b: linalg.Vector[int]{5, 6},
				c: linalg.Vector[int]{1, 1},
			},
			want: StandardForm[int]{
				a: linalg.Matrix[int]{
					{1, 2, 1, 0},
					{3, 4, 0, 1},
				},
				b: linalg.Vector[int]{5, 6},
				c: linalg.Vector[int]{1, 1, 0, 0},
			},
		},
	}

	for testId, testIt := range tests {
		t.Run(testId, func(t *testing.T) {
			form := testIt.got.ToStandard()

			if !form.a.Equals(testIt.want.a) {
				t.Errorf("a != a.\n\ngot:\n%v\n\nwant:\n%v\n", form.a, testIt.want.a)
			}

			if !form.b.Equals(testIt.want.b) {
				t.Errorf("b != b.\n\ngot:\n%v\n\nwant:\n%v\n", form.b, testIt.want.b)
			}

			if !form.c.Equals(testIt.want.c) {
				t.Errorf("c != c.\n\ngot:\n%v\n\nwant:\n%v\n", form.c, testIt.want.c)
			}
		})
	}
}
