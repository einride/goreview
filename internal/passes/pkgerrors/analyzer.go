package pkgerrors

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check for usage of deprecated github.com/pkg/errors library`

const (
	pkgErrorsImportPath = `"github.com/pkg/errors"`
	xErrorsImportPath   = `"golang.org/x/xerrors"`
	errorMessage        = "use " + pkgErrorsImportPath + " instead of " + xErrorsImportPath
	errorMessageOld     = "use " + xErrorsImportPath + " instead of " + pkgErrorsImportPath
)

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "pkgerrors",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	version := strings.Split(runtime.Version(), ".")
	if version[0] != "go1" {
		return nil, fmt.Errorf("unsupported major version: %s", version[0])
	}
	minor, err := strconv.Atoi(version[1])
	if err != nil {
		return nil, fmt.Errorf("unsupported minor version: %s", version[1])
	}

	for _, f := range pass.Files {
		for _, importSpec := range f.Imports {
			if minor >= 13 {
				if importSpec.Path.Value == xErrorsImportPath {
					pass.Reportf(importSpec.Pos(), errorMessage)
				}
			} else {
				if importSpec.Path.Value == pkgErrorsImportPath {
					pass.Reportf(importSpec.Pos(), errorMessageOld)
				}
			}
		}
	}
	return nil, nil
}
