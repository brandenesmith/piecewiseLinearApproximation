package piecewiseLinearApproximation_test

import (
	"piecewiseLinearApproximation"
	"testing"
)

var maxTest = [][]float64{
	{1, 2, 2},
	{3, 5, 5},
	{3, 2, 3},
	{-1, -2, -1},
	{7000, 6999, 7000},
	{2.01, 2.11, 2.11},
	{1.0001, 1.00, 1.0001},
	{0.001, 0.0009, 0.001},
	{0, 0, 0},
}

func TestMax(t *testing.T) {
	for i := 0; i < len(maxTest); i++ {
		answer := piecewiseLinearApproximation.Max(maxTest[i][0], maxTest[i][1])
		if answer != maxTest[i][2] {
			t.Errorf("Max returned %g expecting %g", answer, maxTest[i][2])
		}
	}
}
