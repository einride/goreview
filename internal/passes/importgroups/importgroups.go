package importgroups

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

const Doc = `check for unwanted blank lines between import groups`

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "importgroups",
		Doc:  Doc,
		Run:  run,
	}
}

// importGroup is adapted from importGroup in golang.org/x/tools/goimports.
func importGroup(importPath string) int {
	switch {
	case strings.HasPrefix(importPath, "appengine"):
		return 2
	case strings.Contains(importPath, "."):
		return 1
	default:
		return 0
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		if len(f.Imports) == 0 {
			continue
		}
		prevGroup := importGroup(f.Imports[0].Path.Value)
		prevLine := pass.Fset.Position(f.Imports[0].Pos()).Line
		for i := 1; i < len(f.Imports); i++ {
			pos := f.Imports[i].Pos()
			group := importGroup(f.Imports[i].Path.Value)
			line := pass.Fset.Position(pos).Line
			if line > prevLine+1 && group <= prevGroup {
				pass.Reportf(pos, "remove blank line above import")
			}
			prevGroup = group
			prevLine = line
		}
	}
	return nil, nil
}
