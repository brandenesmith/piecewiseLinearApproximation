# piecewiseLinearApproximation
A simple go package to compute piecewise linear approximations of time series data.
Specifically, this package takes a time series data set of points
{(t,y) | t and y are real numbers with t monotonically increasing} and creates a
constraint relation of the data. The package provides a function to produce the
constraint relation table in LaTeX.

# Using the Package
To use this package simply fork the repository to your go workspace then, from the command line,
navigate to `$GOPATH/src/piecewiseLinearApproximation` and run `go install`.

With this completed, navigate to `$GOPATH/src` and create a new directory.
In this directory create, a new file with a main function as shown below.
```Go
package main

import (
    "fmt"
    p "piecewiseLinearApproximation"
)

func main() {
    var timeSeries1 = []p.Pair {
    	p.Pair{0, 0},
    	p.Pair{1, 1},
    	p.Pair{2, 2},
    	p.Pair{3, 3},
    }

    var tollerance float64 = 2

    equations := p.PiecewiseLinearApprox(timeSeries1, tollerance)

    // Output the equations and their intervals to the standard output.
    for i := 0; i < len(equations); i++ {
        fmt.Printf("%s, %g <= t <= %g\n", equations[i].Expression,
        equations[i].Interval.X, equations[i].Interval.Y)
    }

    fmt.Printf("\n\n")

    // Output a latex table to the standard output.
    p.ToLaTeX(equations, "Solution (b)")
}
```

# Academic Integrity Notice
This repository contains source code used for academic purposes. Any party
reproducing this content for academic purposes should consult their respective
institution’s Academic Integrity Policy. Contributors to this repository strictly
prohibit any reproduction, or use, of this source code, which, violates the Academic
Integrity Policy of any institution, and are not responsible for the violations of any party.

In other words, do not cheat and, if you do, contributors to this repository are
not responsible in any way.
