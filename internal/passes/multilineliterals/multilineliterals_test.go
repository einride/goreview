package multilineliterals_test

import (
	"testing"

	"github.com/einride/goreview/internal/middleware"
	"github.com/einride/goreview/internal/passes/multilineliterals"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	// need to exclude generated test binaries, since test data contains test files
	analyzer := middleware.Apply(
		[]middleware.Middleware{middleware.ExcludeTestBinaries},
		[]*analysis.Analyzer{multilineliterals.Analyzer()})[0]
	analysistest.Run(t, testdata, analyzer, "a")
}
