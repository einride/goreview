package filenames

import (
	"path/filepath"
	"strings"
	"unicode"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check that file names follow conventions`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "filenames",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		filename := filepath.Base(pass.Fset.File(f.Pos()).Name())
		if strings.IndexFunc(filename, unicode.IsUpper) != -1 {
			pass.Reportf(f.Pos(), "file names must be lowercase")
		}
	}
	return nil, nil
}
