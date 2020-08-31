package main

import (
	"github.com/einride/goreview/internal/middleware"
	"github.com/einride/goreview/internal/passes/comments"
	"github.com/einride/goreview/internal/passes/filenames"
	"github.com/einride/goreview/internal/passes/importgroups"
	"github.com/einride/goreview/internal/passes/labels"
	"github.com/einride/goreview/internal/passes/multilinefunctions"
	"github.com/einride/goreview/internal/passes/multilineliterals"
	"github.com/einride/goreview/internal/passes/pkgerrors"
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
		pkgerrors.Analyzer(),
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
