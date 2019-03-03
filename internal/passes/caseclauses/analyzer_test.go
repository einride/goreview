package caseclauses_test

import (
	"testing"

	"github.com/einride/goreview/internal/passes/caseclauses"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, caseclauses.Analyzer(), "a")
}
