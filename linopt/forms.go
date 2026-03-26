package linopt

import (
	"github.com/hy144328/go-ilp/linalg"
	"golang.org/x/exp/constraints"
)

// A CanonicalForm represents the canonical formulation of a linear program:
// Optimize c.x, subject to ax <= b.
type CanonicalForm[T constraints.Integer] struct {
	A linalg.Matrix[T]
	B linalg.Vector[T]
	C linalg.Vector[T]
}

// A StandardForm represents the standard formulation of a linear program:
// Optimize c.x, subject to ax = b.
type StandardForm[T constraints.Integer] struct {
	A linalg.Matrix[T]
	B linalg.Vector[T]
	C linalg.Vector[T]
}

// ToStandard converts from CanonicalForm to StandardForm.
func (form CanonicalForm[T]) ToStandard() StandardForm[T] {
	noRows := form.A.NoRows()
	noColumns := form.A.NoColumns()

	a := linalg.NewMatrix[T](
		noRows,
		noRows + noColumns,
	)
	for rowCt, rowIt := range form.A {
		copy(a[rowCt], rowIt)
		a[rowCt][noColumns + rowCt] = 1
	}

	b := linalg.NewVector[T](form.B.Size())
	copy(b, form.B)

	c := linalg.NewVector[T](form.C.Size() + noColumns)
	copy(c, form.C)

	return StandardForm[T]{a, b, c}
}
