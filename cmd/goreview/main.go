package main

import (
	"github.com/einride/goreview/internal/middleware"
	"github.com/einride/goreview/internal/passes/filenames"
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
		filenames.Analyzer(),
		// ...insert more analyzers here
	}
}

func allMiddleware() []middleware.Middleware {
	return []middleware.Middleware{
		middleware.ExcludeTestBinaries,
	}
}

func main() {
	multichecker.Main(middleware.Apply(allMiddleware(), allAnalyzers())...)
}
