package multilineliterals_test

import (
	"testing"

	"github.com/einride/goreview/internal/passes/multilineliterals"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, multilineliterals.Analyzer(), "a")
}
