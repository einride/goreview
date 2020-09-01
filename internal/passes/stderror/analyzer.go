package stderror

import "golang.org/x/tools/go/analysis"

const Doc = `check for usage of non-standard error libraries`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "stderror",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, importSpec := range f.Imports {
			if importSpec.Path.Value == `"golang.org/x/xerrors"` {
				pass.Reportf(
					importSpec.Pos(),
					"use package \"errors\" and \"fmt.Errorf\" instead of %s",
					importSpec.Path.Value,
				)
			}
		}
	}
	return nil, nil
}
