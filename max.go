package piecewiseLinearApproximation

// Returns the maximum of the two 64 bit floats provided
// as arguments.
func Max(value1, value2 float64) float64 {
	if value1 < value2 {
		return value2
	}

	return value1
}
