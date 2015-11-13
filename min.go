package piecewiseLinearApproximation

// Returns the minimum of two 64 bit floats provided
// as arguments.
func Min(value1, value2 float64) float64 {
	if value1 > value2 {
		return value2
	}

	return value1
}
