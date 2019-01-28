package middleware

import (
	"regexp"

	"golang.org/x/tools/go/analysis"
)

const generatedFilesCommentRegexp = `^// Code generated .* DO NOT EDIT\.$`

// ExcludeTestBinaries modifies an analyzer to completely skip generated test binary packages.
//
// Generated test binary packages have a `.test`-suffix.
func ExcludeGeneratedFiles(run RunFn) RunFn {
	r := regexp.MustCompile(generatedFilesCommentRegexp)
	return func(pass *analysis.Pass) (interface{}, error) {
		filteredFiles := pass.Files[:0]
	FileLoop:
		for _, f := range pass.Files {
			for _, cg := range f.Comments {
				for _, c := range cg.List {
					if r.MatchString(c.Text) {
						continue FileLoop // skip generated file
					}
				}
			}
			filteredFiles = append(filteredFiles, f)
		}
		pass.Files = filteredFiles
		return run(pass)
	}
}
