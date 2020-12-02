package importgroups_test

import (
	"testing"

	"go.einride.tech/review/internal/passes/importgroups"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, importgroups.Analyzer(), "a")
}
