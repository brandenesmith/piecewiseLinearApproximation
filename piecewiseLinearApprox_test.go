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

var minTest = [][]float64{
    {1, 2, 1},
    {3, 5, 3},
    {3, 2, 2},
    {-1, -2, -2},
    {7000, 6999, 6999},
    {2.01, 2.11, 2.01},
    {1.0001, 1.00, 1.00},
    {0.001, 0.0009, 0.0009},
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

func TestMin(t *testing.T) {
    for i := 0; i < len(minTest); i++ {
        answer := piecewiseLinearApproximation.Min(minTest[i][0], minTest[i][1])
        if answer != minTest[i][2] {
            t.Errorf("Min returned %g expecting %g", answer, minTest[i][2])
        }
    }
}
