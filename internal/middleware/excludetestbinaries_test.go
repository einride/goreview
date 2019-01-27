package middleware

import (
	"errors"
	"go/ast"
	"go/types"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/tools/go/analysis"
)

func TestExcludeTestBinaries(t *testing.T) {
	for _, tt := range []struct {
		path     string
		excluded bool
	}{
		{path: "github.com/einride/goreview", excluded: false},
		{path: "github.com/einride/goreview.test", excluded: true},
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
			require.Equal(t, tt.excluded, excluded)
		})
	}
}
