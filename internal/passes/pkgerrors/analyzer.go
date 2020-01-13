package pkgerrors

import "golang.org/x/tools/go/analysis"

const Doc = `check for usage of deprecated github.com/pkg/errors library`

const (
	pkgErrorsImportPath = `"github.com/pkg/errors"`
	xErrorsImportPath   = `"golang.org/x/xerrors"`
	errorMessage        = "use " + pkgErrorsImportPath + " instead of " + xErrorsImportPath
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "pkgerrors",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, importSpec := range f.Imports {
			if importSpec.Path.Value == xErrorsImportPath {
				pass.Reportf(importSpec.Pos(), errorMessage)
			}
		}
	}
	return nil, nil
}
