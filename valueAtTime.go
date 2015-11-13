package piecewiseLinearApproximation

// Returns the value of the piecewise linear approximation
// at a specified time.
func ValueAtTime(lower, upper, time float64, begin Pair) float64 {
	return ((lower+upper)/2)*(time-begin.X) + begin.Y
}
