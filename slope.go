package piecewiseLinearApproximation

// Returns the slope of a line passing through both
// points (Pair) provided as arguments.
func Slope(point1, point2 Pair) float64 {
	slope := (point2.Y - point1.Y) / (point2.X - point1.X)
    
	return slope
}
