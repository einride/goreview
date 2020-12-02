package middleware

import (
	"errors"
	"go/ast"
	"go/types"
	"testing"

	"golang.org/x/tools/go/analysis"
	"gotest.tools/v3/assert"
)

func TestExcludeTestBinaries(t *testing.T) {
	for _, tt := range []struct {
		path     string
		excluded bool
	}{
		{path: "go.einride.tech/review", excluded: false},
		{path: "go.einride.tech/review.test", excluded: true},
	} {
		tt := tt
		t.Run(tt.path, func(t *testing.T) {
			// given a pass with one file
			pass := &analysis.Pass{
				Pkg:   types.NewPackage(tt.path, "pkg"),
				Files: []*ast.File{nil}, // length: 1
			}
			// and a run function that fails when there are any files
			run := func(pass *analysis.Pass) (interface{}, error) {
				if len(pass.Files) > 0 {
					return nil, errors.New("boom")
				}
				return nil, nil
			}
			// when excluding files
			_, err := ExcludeTestBinaries(run)(pass)
			// then files for generated test binaries should be excluded
			excluded := err == nil
			assert.Equal(t, tt.excluded, excluded)
		})
	}
}
