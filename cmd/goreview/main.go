package main

import (
	"strings"

	"github.com/einride/goreview/internal/passes/importgroups"
	"github.com/einride/goreview/internal/passes/multilineliterals"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

// allAnalyzers returns all analyzers to include in the analysis.
func allAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		importgroups.Analyzer(),
		multilineliterals.Analyzer(),
		// ...insert more analyzers here
	}
}

// runFn is a short-hand type alias for the (*Analyzer).Run function.
type runFn = func(*analysis.Pass) (interface{}, error)

// middleware modifies the behavior of a runFn.
type middleware func(runFn) runFn

func main() {
	multichecker.Main(applyMiddleware([]middleware{
		excludeGeneratedTestBinaries,
		// ...insert more middleware here
	}, allAnalyzers())...)
}

// applyMiddleware applies the provided middleware(s) to the run function of each provided Analyzer.
func applyMiddleware(ms []middleware, as []*analysis.Analyzer) []*analysis.Analyzer {
	for _, a := range as {
		for _, m := range ms {
			a.Run = m(a.Run)
		}
	}
	return as
}

// excludeGeneratedTestBinaries modifies a runFn to completely skip generated test binary packages.
//
// Generated test binary packages have a `.test`-suffix.
func excludeGeneratedTestBinaries(run runFn) runFn {
	return func(pass *analysis.Pass) (i interface{}, e error) {
		if strings.HasSuffix(pass.Pkg.Path(), ".test") {
			return nil, nil
		}
		return run(pass)
	}
}
