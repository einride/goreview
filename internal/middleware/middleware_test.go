package middleware

import (
	"errors"
	"testing"

	"golang.org/x/tools/go/analysis"
	"gotest.tools/v3/assert"
)

func TestApply(t *testing.T) {
	// given analyzers that always return errors
	originals := []*analysis.Analyzer{
		{
			Name: "analyzer1",
			Run: func(_ *analysis.Pass) (interface{}, error) {
				return nil, errors.New("boom")
			},
			Requires: []*analysis.Analyzer{
				{
					Name: "require1",
					Run: func(*analysis.Pass) (interface{}, error) {
						return nil, errors.New("boom")
					},
				},
				{
					Name: "require2",
					Run: func(*analysis.Pass) (interface{}, error) {
						return nil, errors.New("boom")
					},
				},
			},
		},
		{
			Name: "analyzer2",
			Run: func(*analysis.Pass) (interface{}, error) {
				return nil, errors.New("boom")
			},
		},
	}
	// when applying middleware that never returns errors
	m := func(_ RunFn) RunFn {
		return func(*analysis.Pass) (interface{}, error) {
			return nil, nil
		}
	}
	// then no errors should be returned
	for _, a := range Apply([]Middleware{m}, originals) {
		_, err := a.Run(nil)
		assert.NilError(t, err)
		for _, r := range a.Requires {
			_, err = r.Run(nil)
			assert.NilError(t, err)
		}
	}
}
