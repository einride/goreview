package filenames_test

import (
	"testing"

	"go.einride.tech/review/internal/passes/filenames"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, filenames.Analyzer(), "a")
}
