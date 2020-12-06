package ptypes

import "golang.org/x/tools/go/analysis"

const ptypesPackage = `"github.com/golang/protobuf/ptypes"`

const knownTypesPackage = `"google.golang.org/protobuf/types/known"`

const Doc = `check for usage of deprecated package ` + ptypesPackage

func Analyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "ptypes",
		Doc:  Doc,
		Run:  run,
	}
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, f := range pass.Files {
		for _, importSpec := range f.Imports {
			if importSpec.Path.Value == ptypesPackage {
				pass.Reportf(
					importSpec.Pos(),
					`use the the packages from %s instead of the deprecated package %s (see %s)`,
					knownTypesPackage,
					ptypesPackage,
					"https://github.com/golang/protobuf/releases#v1.4-well-known-types",
				)
			}
		}
	}
	return nil, nil
}
