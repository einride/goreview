package middleware

import (
	"testing"

	"go.einride.tech/review/internal/passes/importgroups"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analyzer := Apply([]Middleware{ExcludeGeneratedFiles}, []*analysis.Analyzer{importgroups.Analyzer()})[0]
	analysistest.Run(t, testdata, analyzer, "excludegeneratedtestfiles")
}
