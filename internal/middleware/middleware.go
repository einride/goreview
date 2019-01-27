package middleware

import "golang.org/x/tools/go/analysis"

// RunFn is a short-hand type alias for the (*Analyzer).Run function.
type RunFn = func(*analysis.Pass) (interface{}, error)

// Middleware modifies the behavior of a run function.
type Middleware func(RunFn) RunFn

// Apply the provided middleware(s) to the run function of each provided Analyzer.
func Apply(ms []Middleware, as []*analysis.Analyzer) []*analysis.Analyzer {
	for _, a := range as {
		for _, m := range ms {
			a.Run = m(a.Run)
			for _, r := range a.Requires {
				r.Run = m(r.Run)
			}
		}
	}
	return as
}
