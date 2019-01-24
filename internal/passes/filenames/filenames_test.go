package filenames_test

import (
	"testing"

	"github.com/einride/goreview/internal/passes/filenames"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, filenames.Analyzer(), "a")
}
