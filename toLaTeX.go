package piecewiseLinearApproximation

import (
    "fmt"
)

// Outputs a LaTeX table representing the constraint relation
// to the standard output. 
func ToLaTeX(equations []Equation, relationName string) {
    fmt.Printf("\\begin{figure}[h!]\n")
    fmt.Printf("\\centering\n")
    fmt.Printf("\\textbf{")
    fmt.Printf(relationName)
    fmt.Printf("}\\\\")
    fmt.Printf("\n")
    fmt.Printf("\t\\begin{tabular}{| l l l l|}\n")
    fmt.Printf("\t\t\\hline\n")
    fmt.Printf("\t\tId & $T$ & $Y$ & \\\\ \n")
    fmt.Printf("\t\t\\hline\n")
    fmt.Printf("\t\t\\hline\n")

    for i := 0; i < len(equations); i ++ {
        fmt.Printf("\t\t%d & $t$ & $y$ & $%s$, $%g \\leq t \\leq %g$\\\\ \n", i+1,
            equations[i].Expression, equations[i].Interval.X, equations[i].Interval.Y)
        fmt.Printf("\t\t\\hline\n")
    }

    fmt.Printf("\t\\end{tabular}\n")
    fmt.Printf("\\end{figure}\n\n")
}
