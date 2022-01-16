package middleware

import (
	"testing"

	"go.einride.tech/review/internal/passes/filenames"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := Apply([]Middleware{ExcludeGeneratedFiles}, []*analysis.Analyzer{filenames.Analyzer()})[0]
	analysistest.Run(t, testdata, analyzer, "excludegeneratedtestfiles")
}
