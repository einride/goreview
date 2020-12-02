package comments_test

import (
	"testing"

	"go.einride.tech/review/internal/passes/comments"
	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, comments.Analyzer(), "a")
}
