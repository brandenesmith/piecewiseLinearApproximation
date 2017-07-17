package piecewiseLinearApproximation

// Pair represents an ordered pair of
// 64 bit float values.
type Pair struct {
	X, Y float64
}

// Equation represents an equation (Expression)
// as a string and a domain interval (Pair)
type Equation struct {
	Expression string
	Interval   Pair
    Slope float64
    Plot func(float64) float64
}
