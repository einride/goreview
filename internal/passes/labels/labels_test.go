package labels_test

import (
	"testing"

	"go.einride.tech/review/internal/passes/labels"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, labels.Analyzer(), "a")
}
