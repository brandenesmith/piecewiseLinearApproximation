package piecewiseLinearApproximation

import (
	"math"
	"strconv"
)

func getPlotFunction(slope float64, beginX float64, beginY float64) func(float64) float64 {
	return func(x float64) float64 {
		return slope * (x - beginX) + beginY
	}
}

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
			slope := (upper+lower)/2
			expression = "y = " +
				strconv.FormatFloat(slope, 'g', -1, 64) + "(t - " +
				strconv.FormatFloat(begin.X, 'g', -1, 64) +
				") + " + strconv.FormatFloat(begin.Y, 'g', -1, 64)

			interval = Pair{start, end}
			equations = append(equations, Equation{expression, interval, slope, getPlotFunction(slope, begin.X, begin.Y)})

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

	slope := (upper+lower)/2
	expression = "y = " +
		strconv.FormatFloat(slope, 'g', -1, 64) + "(t - " +
		strconv.FormatFloat(begin.X, 'g', -1, 64) +
		") + " + strconv.FormatFloat(begin.Y, 'g', -1, 64)

	interval = Pair{start, end}
	equations = append(equations, Equation{expression, interval, slope, getPlotFunction(slope, begin.X, begin.Y)})

	return equations
}
