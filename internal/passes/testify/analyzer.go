package testify

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check for usage of testify and recommend to use gotest`

const (
	testifyModule  = "github.com/stretchr/testify"
	gotestV3Module = "gotest.tools/v3"
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "testify",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, importSpec := range f.Imports {
			if strings.HasPrefix(importSpec.Path.Value, `"`+testifyModule) {
				pass.Reportf(
					importSpec.Pos(),
					`use test assertions from "%s" instead of "%s" (see %s)`,
					gotestV3Module,
					testifyModule,
					"https://pkg.go.dev/gotest.tools/v3/assert",
				)
			}
		}
	}
	return nil, nil
}
