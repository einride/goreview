package middleware

import (
	"strings"

	"golang.org/x/tools/go/analysis"
)

// ExcludeTestBinaries modifies an analyzer to completely skip generated test binary packages.
//
// Generated test binary packages have a `.test`-suffix.
func ExcludeTestBinaries(run RunFn) RunFn {
	return func(pass *analysis.Pass) (i interface{}, e error) {
		if strings.HasSuffix(pass.Pkg.Path(), ".test") {
			pass.Files = nil
		}
		return run(pass)
	}
}
