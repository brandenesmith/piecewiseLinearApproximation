package piecewiseLinearApproximation_test

import (
	"piecewiseLinearApproximation"
	"testing"
	"strings"
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

var slopeTest = [][]piecewiseLinearApproximation.Pair {
	// Zero slope
	{piecewiseLinearApproximation.Pair{0, 1}, piecewiseLinearApproximation.Pair{5, 1}},
	{piecewiseLinearApproximation.Pair{3, 7}, piecewiseLinearApproximation.Pair{8, 7}},
	{piecewiseLinearApproximation.Pair{2, 22}, piecewiseLinearApproximation.Pair{22, 22}},

	// Positive slope
	{piecewiseLinearApproximation.Pair{0, 0}, piecewiseLinearApproximation.Pair{1, 1}},
	{piecewiseLinearApproximation.Pair{1, 3}, piecewiseLinearApproximation.Pair{2, 7}},
	{piecewiseLinearApproximation.Pair{7, 2}, piecewiseLinearApproximation.Pair{9, 2.5}},

	// Negative slope
	{piecewiseLinearApproximation.Pair{0, 1}, piecewiseLinearApproximation.Pair{1, 0}},
	{piecewiseLinearApproximation.Pair{1, 7}, piecewiseLinearApproximation.Pair{2, 3}},
	{piecewiseLinearApproximation.Pair{7, 2.5}, piecewiseLinearApproximation.Pair{9, 2}},
}

var slopeSolutions = []float64 {
	0,
	0,
	0,
	1,
	4,
	0.25,
	-1,
	-4,
	-0.25,
}

var tollerance1 float64 = 1

var timeSeries1 = []piecewiseLinearApproximation.Pair {
	piecewiseLinearApproximation.Pair{1, 2},
	piecewiseLinearApproximation.Pair{2, 4},
	piecewiseLinearApproximation.Pair{3, 2},
	piecewiseLinearApproximation.Pair{4, 4},
	piecewiseLinearApproximation.Pair{5, 5},
	piecewiseLinearApproximation.Pair{6, 4},
	piecewiseLinearApproximation.Pair{7, 5},
	piecewiseLinearApproximation.Pair{8, 5},
	piecewiseLinearApproximation.Pair{9, 6},
	piecewiseLinearApproximation.Pair{10, 4},
	piecewiseLinearApproximation.Pair{11, 4},
	piecewiseLinearApproximation.Pair{12, 6},
	piecewiseLinearApproximation.Pair{13, 6},
	piecewiseLinearApproximation.Pair{14, 6},
	piecewiseLinearApproximation.Pair{15, 7},
	piecewiseLinearApproximation.Pair{16, 7},
	piecewiseLinearApproximation.Pair{17, 9},
	piecewiseLinearApproximation.Pair{18, 10},
	piecewiseLinearApproximation.Pair{19, 11},
	piecewiseLinearApproximation.Pair{20, 13},
}

var equations1 = []string {
	"y = 2(t - 1) + 2",
	"y = -2(t - 2) + 4",
	"y = 1(t - 3) + 2",
	"y = -0.41666666666666663(t - 7) + 6",
	"y = 0.7208333333333332(t - 11) + 4.333333333333334",
	"y = 2.9000000000000004(t - 19) + 10.1",
}

var intervals1 = []piecewiseLinearApproximation.Pair {
	piecewiseLinearApproximation.Pair{1, 2},
	piecewiseLinearApproximation.Pair{2, 3},
	piecewiseLinearApproximation.Pair{3, 7},
	piecewiseLinearApproximation.Pair{7, 11},
	piecewiseLinearApproximation.Pair{11, 19},
	piecewiseLinearApproximation.Pair{19, 20},
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

func TestSlope(t *testing.T) {
	for i := 0; i < len(slopeTest); i++ {
		answer := piecewiseLinearApproximation.Slope(slopeTest[i][0], slopeTest[i][1])
		if answer != slopeSolutions[i] {
			t.Errorf("Slop returned %g, expecting %g", answer, slopeSolutions[i])
		}
	}
}

func TestPiecewiseLinearApprox(t *testing.T) {
	solution := piecewiseLinearApproximation.PiecewiseLinearApprox(timeSeries1, tollerance1)
	for i := 0; i < len(equations1); i++ {
		expression := solution[i].Expression
		interval := solution[i].Interval
		if strings.Compare(expression, equations1[i]) != 0 {
			t.Errorf("equations do not match")
		} else if interval.X != intervals1[i].X || interval.Y != intervals1[i].Y {
			t.Errorf("intervals do not match")
		}
	}
}
