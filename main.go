package main

import (
	"go.einride.tech/review/internal/middleware"
	"go.einride.tech/review/internal/passes/comments"
	"go.einride.tech/review/internal/passes/filenames"
	"go.einride.tech/review/internal/passes/importgroups"
	"go.einride.tech/review/internal/passes/labels"
	"go.einride.tech/review/internal/passes/multilinefunctions"
	"go.einride.tech/review/internal/passes/multilineliterals"
	"go.einride.tech/review/internal/passes/ptypes"
	"go.einride.tech/review/internal/passes/stderror"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
)

// allAnalyzers returns all analyzers to include in the analysis.
func allAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{
		importgroups.Analyzer(),
		multilineliterals.Analyzer(),
		multilinefunctions.Analyzer(),
		filenames.Analyzer(),
		comments.Analyzer(),
		labels.Analyzer(),
		stderror.Analyzer(),
		ptypes.Analyzer(),
		// ...insert more analyzers here
	}
}

// allMiddleware returns all middleware to apply to the included analyzers.
func allMiddleware() []middleware.Middleware {
	return []middleware.Middleware{
		middleware.ExcludeTestBinaries,
		middleware.ExcludeGeneratedFiles,
		// ...insert more middleware here
	}
}

func main() {
	multichecker.Main(middleware.Apply(allMiddleware(), allAnalyzers())...)
}
