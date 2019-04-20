package multilinefunctions_test

import (
	"testing"

	"github.com/einride/goreview/internal/passes/multilinefunctions"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestFunctionDeclarationParameters(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), multilinefunctions.Analyzer(), "a")
}

func TestFunctionDeclarationResults(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), multilinefunctions.Analyzer(), "b")
}

func TestFunctionCalls(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), multilinefunctions.Analyzer(), "c")
}
