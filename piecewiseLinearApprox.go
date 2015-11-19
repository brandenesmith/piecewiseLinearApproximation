package piecewiseLinearApproximation

import (
	"math"
	"strconv"
)

func PiecewiseLinearApprox(timeSeries []Pair, tollerance float64) []Equation {
	var expression string
	var interval Pair

	equations := make([]Equation, 0)
	begin := timeSeries[0]
	lower := - math.MaxFloat64
	upper := math.MaxFloat64

	start := begin.X
	end := begin.X

	for i := 0; i < len(timeSeries)-1; i++ {
		lowerPrime := Max(lower, Slope(begin, Pair{timeSeries[i+1].X,
			timeSeries[i+1].Y - tollerance}))
		upperPrime := Min(upper, Slope(begin, Pair{timeSeries[i+1].X,
			timeSeries[i+1].Y + tollerance}))
		end = timeSeries[i].X

		if lowerPrime <= upperPrime {
			lower = lowerPrime
			upper = upperPrime
		} else {
			expression = "y = " +
				strconv.FormatFloat((upper+lower)/2, 'g', -1, 64) + "(t - " +
				strconv.FormatFloat(begin.X, 'g', -1, 64) +
				") + " + strconv.FormatFloat(begin.Y, 'g', -1, 64)

			interval = Pair{start, end}
			equations = append(equations, Equation{expression, interval})

			begin = Pair{timeSeries[i].X, ValueAtTime(lower, upper,
				timeSeries[i].X, begin)}
			lower = Slope(begin, Pair{timeSeries[i+1].X,
				timeSeries[i+1].Y - tollerance})
			upper = Slope(begin, Pair{timeSeries[i+1].X,
				timeSeries[i+1].Y + tollerance})
			start = end
		}
	}
	end = timeSeries[len(timeSeries)-1].X

	expression = "y = " +
		strconv.FormatFloat((upper+lower)/2, 'g', -1, 64) + "(t - " +
		strconv.FormatFloat(begin.X, 'g', -1, 64) +
		") + " + strconv.FormatFloat(begin.Y, 'g', -1, 64)

	interval = Pair{start, end}
	equations = append(equations, Equation{expression, interval})

	return equations
}
