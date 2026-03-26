package linopt

import (
	"errors"
	"fmt"

	"github.com/hy144328/go-ilp/linalg"
	"golang.org/x/exp/constraints"
)

var (
	ErrIncompatibleSizes = errors.New("incompatible sizes")
)

// A CanonicalForm represents the canonical formulation of a linear program:
// Optimize c.x, subject to ax <= b.
type CanonicalForm[T constraints.Integer] struct {
	a linalg.Matrix[T]
	b linalg.Vector[T]
	c linalg.Vector[T]
}

// A StandardForm represents the standard formulation of a linear program:
// Optimize c.x, subject to ax = b.
type StandardForm[T constraints.Integer] struct {
	a linalg.Matrix[T]
	b linalg.Vector[T]
	c linalg.Vector[T]
}

// NewCanonicalForm constructs a CanonicalForm including inequality constraints and cost function.
func NewCanonicalForm[T constraints.Integer](
	a linalg.Matrix[T],
	b linalg.Vector[T],
	c linalg.Vector[T],
) (CanonicalForm[T], error) {
	res := CanonicalForm[T]{a, b, c}

	if a.NoRows() != b.Size() {
		return res, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, a.NoRows(), b.Size())
	}

	if a.NoColumns() != c.Size() {
		return res, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, a.NoColumns(), c.Size())
	}

	return res, nil
}

// NewStandardForm constructs a StandardForm including equality constraints and cost function.
func NewStandardForm[T constraints.Integer](
	a linalg.Matrix[T],
	b linalg.Vector[T],
	c linalg.Vector[T],
) (StandardForm[T], error) {
	res := StandardForm[T]{a, b, c}

	if a.NoRows() != b.Size() {
		return res, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, a.NoRows(), b.Size())
	}

	if a.NoColumns() != c.Size() {
		return res, fmt.Errorf("%w: %d, %d.", ErrIncompatibleSizes, a.NoColumns(), c.Size())
	}

	return res, nil
}

// NoConstraints() returns the number of constraints in the canonical formulation.
func (form CanonicalForm[T]) NoConstraints() int {
	return form.a.NoRows()
}

// NoVariables() returns the number of variables in the canonical formulation.
func (form CanonicalForm[T]) NoVariables() int {
	return form.a.NoColumns()
}

// ToStandard converts from CanonicalForm to StandardForm.
func (form CanonicalForm[T]) ToStandard() StandardForm[T] {
	noConstraints := form.NoConstraints()
	noVariables := form.NoVariables()

	a := linalg.NewMatrix[T](
		noConstraints,
		noVariables + noConstraints,
	)
	for rowCt, rowIt := range form.a {
		copy(a[rowCt], rowIt)
		a[rowCt][noVariables + rowCt] = 1
	}

	b := linalg.NewVector[T](noConstraints)
	copy(b, form.b)

	c := linalg.NewVector[T](noVariables + noConstraints)
	copy(c, form.c)

	res, err := NewStandardForm(a, b, c)
	if err != nil {
		panic(err)
	}

	return res
}

// NoConstraints() returns the number of constraints in the standard formulation.
func (form StandardForm[T]) NoConstraints() int {
	return form.a.NoRows()
}

// NoVariables() returns the number of variables in the standard formulation.
func (form StandardForm[T]) NoVariables() int {
	return form.a.NoColumns()
}
